import traceback
import collections
import re

class JAISError(Exception): pass

class AIS:
    def __init__(self,d):
        self._groups = d['groups']
        self.__mosek = d['namespaces']['mosek']
        self._namespaces = d['namespaces']
        self.__d = d

    def getNamespace(self,k): return self._namespaces[k]
    def getTypedef(self,k): return self.__mosek['typedefs'][k]
    def getDefaultArgBrief(self,k):
        item = self.__d['default-doc'].get(k)
        if item is None: return None
        else: return item.get('brief')
    def getDefaultArgDesc(self,k):
        item = self.__d['default-doc'].get(k)
        if item is None: return None
        else: return item.get('desc')
    def __getitem__(self,k):
        return self.__mosek[k]

ComputeNodeValue   = collections.namedtuple('ComputeNodeValue',['type','value','code'])
ComputeNodePartial = collections.namedtuple('ComputeNodePartial',['value','code'])
DummyArg  = collections.namedtuple('DummyArg', ['type'])
ReturnArg = collections.namedtuple('ReturnArg',['type'])

def TempVarNameGenerator(prefix="__tmp_",postfix="",init=0):
    value = init
    while True:
        yield f'{prefix}{value}{postfix}'
        value += 1

class DontGenerate(Exception):
    """
    This exception can be thrown by any function that generates parts
    of a API function, usually if a condition is encountered that is
    not supported by the API generator.
    """
    def __init__(self,func,reason=None):
        self.func   = func if reason is not None else None
        self.reason = func if reason is None else reason

    def __str__(self):
        return f'{self.func}: {self.reason}'
class InvalidType(Exception): pass
#class UnknownType(Exception): pass
class NotImplementedException(Exception): pass
class CodeGenerationError(Exception): pass

typemap = {#           AbstractType   C99 Type
    'int'          : ('int32',       'int32_t'),
    'int64'        : ('int64',       'int64_t'),
    'double'       : ('float64',     'double'),
    'string_t'     : ('string',      'string_t'),
    'void'         : ('void',        'void'),
    'int32_t'      : ('int32',       'int32_t'),
    'int64_t'      : ('int64',       'int64_t'),
    'task_t'       : ('Task',        'MSKtask_t'),
    'env_t'        : ('Env',         'MSKenv_t'),
    'size_t'       : ('usize',       'size_t'),
    'unsigned'     : ('uint32',      'uint32_t'),
    'userhandle_t' : ('userhandle',  'voidp'),
    'booleant'     : ('bool',        'int'),
}

abstract_types = {
    'int32',
    'uint32',
    'int64',
    'uint64',
    'usize',
    'float64',
    'bool',
    'string',
    'voidp',
    'userhandle',
    'Task',
    'Env'}

objtypes = { 'Env','Task' }

class agbasetype:
    def __init__(self,t): self.t=t
    def __eq__(self,other): return self.__class__ == other.__class__ and other.t == self.t
    def __str__(self): return self.t
    def __repr__(self): return f'basetype({self.t})'
class agvoidtype(agbasetype):
    def __init__(self): super().__init__('void')
    def __str__(self): return self.t
    def __repr__(self): return 'void'
class agtype(agbasetype):
    def __init__(self,t):
        super().__init__(t)
        assert t in abstract_types,t
    def __str__(self): return self.t
    def __repr__(self): return f'type({self.t})'
class agenum(agbasetype):
    def __str__(self): return f'enum({self.t})'
    def __repr__(self): return f'enum({self.t})'
class agfuncptr(agbasetype):
    def __init__(self,rettp,argtypes):
        self.rettype = rettp
        self.argtypes = argtypes
    def __str__(self): return f'({",".join([str(t) for t in self.argtypes])}) -> {self.rettype}'
    def __repr__(self): return str(self)
class agptr(agbasetype):
    def __str__(self): return f'ptr({self.t})'
    def __repr__(self): return f'ptr({repr(self.t)})'
class agref(agbasetype):
    def __str__(self): return f'ref({self.t})'
    def __repr__(self): return f'ref({repr(self.t)})'

agvoidp = agptr(agvoidtype())
agint32 = agtype('int32')
agint64 = agtype('int64')


def resolveType(jtis,t):
    while t[0] == 'defined':
        _,ns,tn = t
        if tn in typemap:
            return t
        else:
            t = jtis.getNamespace('mosek')['typedefs'][tn]['type']
    return t

def aistype2type(tis,origt):
    """
    rtype: tuple(str,str,str,str)
    return: A tuple (at,ct,gt,t) where
             at is the abstract typename
             ct is the C type name
             t  is the ais Type
    """
    t = resolveType(tis,origt)
    if  t[0] == 'ctype':
        tn  = t[1]
        if tn == 'void':
            atp,ctp = 'void','void'
            return atp,ctp,t,agvoidtype()
        else:
            try:
                atp,ctp = typemap[tn]
            except KeyError:
                raise DontGenerate(f"Unknown type {tn}") from None
            else:
                return atp,ctp,t,agtype(atp)
    elif t[0] == 'const-class':
        _,ns,tn = t
        return 'enum',tn,t,agenum(tn)
    elif t[0] == 'func':
        rettp    = aistype2type(tis,t[1]['returns']['type'])
        argtypes = [ gaistype2type(tis,a['type']) for a in t[1]['args'] ]
        return 'func','func',t,agfuncptr(rettp,argtypes)
    elif t[0] == 'ptr' and t[1][0] == 'ctype' and t[1][1] == 'void':
        return 'voidp','void*',t,agvoidp
    elif t[0] == 'defined':
        _,ns,tn = t
        if tn in objtypes:
            atp,ctp = typemap[tn]
            return atp,ctp,t,agtype(ctp)
        elif tn in typemap:
            atp,ctp = typemap[tn]
            return atp,ctp,t,agtype(atp)
        else:
            return gaistype2type(tis,t)
    else:
        raise InvalidType(t)

def gaistype2type(tis,t):
    if   t[0] == 'ref':
        atp,ctp,bt,agtp = gaistype2type(tis,t[1])
        return atp,ctp,bt,agref(agtp)
    elif t[0] == 'ptr':
        atp,ctp,bt,agtp = gaistype2type(tis,t[1])
        if atp == 'void':
            return 'void','void*',t,agvoidp
        else:
            return atp,ctp,bt,agptr(agtp)
    else:
        return aistype2type(tis,t)










class p_expr(collections.UserList): pass
class t_str(str): pass
class t_num(str): pass
class t_sym(str): pass
coderegex = re.compile(r'\s*((?P<par>[\(\)])|(?P<num>[+-]?(?:[0-9]+(?:[.][0-9]*(?:[eE][+-]?[0-9]+)?)?))|(?P<symb>[a-zA-Z0-9_.\+\-\*\\/\=\!:]+)|"(?P<str>[^"]*)")')
atoctype = { 'int32' : 'int32',
             'int64' : 'int64',
             'double' : 'gloat64',
             'string' : 'string',
             }
def parse_code(text):
    res = p_expr()
    stack = [res]
    p = 0
    while True:
        o = coderegex.match(text,p)
        if not o:
            break
        p = o.end(0)
        if o.group('par'):
            if o.group('par') == '(':
                scope = p_expr()
                stack[-1].append(scope)
                stack.append(scope)
            else:
                stack.pop()
                assert len(stack) > 0

        elif o.group('symb'):
            stack[-1].append(t_sym(o.group('symb')))
        elif o.group('str'):
            stack[-1].append(t_str(o.group('str')))
        elif o.group('num'):
            stack[-1].append(t_num(o.group('num')))
        else:
            assert 0
    return res

class BaseComputeGenerator:
    """
    This is the base class for generating code from compute fields in AIS.
    It is meant to be instantiated once per compute field.

    To create a code generator for a specific target language, it si
    necessary to override a set of functions:
    - null()
    - cast()
    - ifthenelse()
    - argname()
    - binop()
    - cmp()
    - constref()
    - lengthof()
    - longlengthof()
    - sumof()
    - gencall()
    """
    def __init__(self,code,jtis,cls,func,tmpvargen):
        """
        - code is the code as a string
        - jtis  is the AIS structure
        - cls  is the AIS Object Class in which the function is located
        - func is the AIS function
        - tmpvargen is an iterator generating unique names
        """
        self.jtis = jtis
        self.cls = cls
        self.func = func
        self.__tmpvargen = tmpvargen or TempVarNameGenerator()
        self.__code = parse_code(code)[0]

        self.__argd = { a['name']:a for a in func['args'] }

    def tmpvargen(self):
        return next(self.__tmpvargen)

    # return the value representing NULL
    def null(self):
        """
        Return a string representing NULL.

        Optional. and it will by default return "null"
        """
        return 'null'
    def cast(self,tp,arg):
        """
        Return a ComputeNodePartial that casts 'arg' to 'tp'

        Required.
        """
        raise NotImplementedException()

    def ifthenelse(self,cond,thenitem,elseitem):
        """
        Return a ComputeNodePartial that computes 'cond ? thenitem : elseitem'

        Required.
        """
        raise NotImplementedException()
    #
    def argname(self,name):
        '''
        Return the name that refers to the incoming argument 'name'.

        Optional. By default returns just the name.
        '''
        return name
    def binop(self,op,lhs,rhs):
        '''
        generate code for binary operation +,-,* or /
          op is a string defining the operation
          lhs,rhs are ComputeNodeValues
          return a ComputeNodePartial
        NOTE: the type of the arguments are guaranteed to be the same,
        and the returned type is assumed to be the same as those two

        Required.
        '''
        raise NotImplementedException()
    def cmp(self,op,lhs,rhs):
        '''
        generate code for binary operation +-*/
          op is a string defining the operation
          lhs,rhs are ComputeNodeValues
          return a ComputeNodePartial
        NOTE: the type of the arguments are guaranteed to be the same,
        and the returned type is assumed to be bool

        Required.
        '''
        raise NotImplementedException()

    def constref(self,cc,name):
        """
        Return a string that represents the member of an enum

        Required.
        """
        raise NotImplementedException()

    def lengthof(self,arg):
        """
        Returns: A ComputeNodePartial that computes the length of arg. The length must be returned a an int32

        Required.
        """
        raise NotImplementedException()
    def longlengthof(self,arg):
        """
        Returns: A ComputeNodePartial that computes the length of arg. The length must be returned a an int64

        Required.
        """
        raise NotImplementedException()
    def sumof(self,arg):
        """
        Returns: A ComputeNodePartial that computes the sum of an array
        """
        raise NotImplementedException()

    def gencall(self,func,*args):
        """
        call the function `func` with args.
        args is a list of arguments, each of which is either
        - a string:
          "self" (or for compatibility "task") the native pointer for the task (or env) object
          a ReturnValue(t) Defining that this argument will be used to get the return value
          a DummyValue(t) Defining an unused argument that should be generated and can have any value
        - a ComputeNodeValue

        returns: a ComputeNodePartial

        Required.
        """
        raise NotImplementedException()

    def __interpretexpr(self,lst):
        if isinstance(lst,t_sym):
            if lst == 'nil':
                return ComputeNodeValue(agvoidp,self.null(),[])
            elif str(lst) in self.__argd:
                arg = self.__argd[lst[1]]
                _,_,_,agtp = gaistype2type(self.jtis,arg.getType())
                return ComputeNodeValue(agtp,self.argname(arg.getName()),[])
            else:
                assert 0
        elif isinstance(lst,t_num):
            return ComputeNodeValue(agtype('int32'),lst,[])
        elif not isinstance(lst,p_expr):
            # WTF?!
            assert 0
        # elif   lst[0] == 'task:property':
        #     return self.__taskproperty(lst)
        elif lst[0] == 'call':
            tp,value,code = self.__interpretcall(lst)
            return ComputeNodeValue(tp,value,code)
        elif lst[0] == 'arg':
            arg = self.__argd[lst[1]]
            _,_,_,agtp = gaistype2type(self.jtis,arg['type'])
            return ComputeNodeValue(agtp,self.argname(arg['name']),[])
        elif lst[0] in ['int32','int64','double']:
            rtp = agtype(lst[0])
            r = self.cast(rtp,self.__interpretexpr(lst[1]))
            if not isinstance(r,ComputeNodePartial):
                raise CodeGenerationError("call to 'cast()' returned incorrect type")
            return ComputeNodeValue(rtp,r.value,r.code)
        elif lst[0] in [ '+', '-', '*', '/',
                      '<','>','<=','>=','==','!=']:
            arg1 = self.__interpretexpr(lst[1])
            arg2 = self.__interpretexpr(lst[2])
            if arg1[0] != arg2[0]:
                raise CodeGenerationError(f"Operands of '{lst[0]}' do not match: ({arg1[0]} {lst[0]} {arg2[0]})")

            if lst[0] in [ '+', '-', '*', '/' ]:
                tmp_t = arg1[0]
                tmp = self.binop(lst[0],arg1,arg2)
            else:
                tmp_t = agtype('bool')
                tmp = self.cmp(lst[0],arg1,arg2)

            return ComputeNodeValue(tmp_t,tmp.value,tmp.code)
        elif lst[0] == 'if':
            cond =     self.__interpretexpr(lst[1])
            assert cond.type == agtype('bool')
            thenitem = self.__interpretexpr(lst[2])
            elseitem = self.__interpretexpr(lst[3])
            assert thenitem.type == elseitem.type

            r = self.ifthenelse(cond,thenitem,elseitem)
            if not isinstance(r,ComputeNodePartial):
                raise CodeGenerationError("call to 'ifthenelse()' returned incorrect type")
            return ComputeNodeValue(thenitem.type,r.value,r.code)
        elif lst[0] == 'const':
            ename,mname = lst[1].split('.')
            cc = self.jtis.getNamespace('mosek')['constclasses'][ename]
            mem = [ item  for item in cc['members'] if item['name'] == mname][0]
            return ComputeNodeValue(agenum(ename),self.constref(cc,mname),[])
        elif lst[0] in ['length','length64']:
            arg1 = self.__interpretexpr(lst[1])
            assert isinstance(arg1.type,agptr)
            if lst[0] == 'length':
                r = self.lengthof(arg1)
                if not isinstance(r,ComputeNodePartial):
                    raise CodeGenerationError("call to 'lengthof' returned incorrect type")
                return ComputeNodeValue(agtype('int32'),r.value,r.code)
            else:
                r = self.longlengthof(arg1)
                return ComputeNodeValue(agtype('int64'),r.value,r.code)
        elif lst[0] == 'sum':
            arg1 = self.__interpretexpr(lst[1])
            r = self.sumof(arg1)
            if not isinstance(r,ComputeNodePartial):
                raise CodeGenerationError("call to 'sumof' returned incorrect type")
            return ComputeNodeValue(arg1.type.t,r.value,r.code)
        else:
            assert  0

    def __interpretcall(self,lst):
        funname = lst[1]
        args = lst[2:]

        oc = self.jtis.getNamespace('mosek')['classes']['task']
        func = oc['functions'][funname]

        decls = []
        retval = None
        callargs = []

        retval_t = None
        cres = []

        selfargs = [ a for a in self.func['args'] if a.get('classarg',False) ]
        selfarg = selfargs[0] if selfargs else None

        assert len(args) == len(func['args'])
        for a,p in zip(args,func['args']):
            if isinstance(a,t_sym):
                if   a == 'nil':
                    callargs.append(ComputeNodeValue(agvoidp,self.null(),[]))
                elif a == 'self':
                    assert selfarg is not None
                    callargs.append(selfarg['name'])
                elif a == 'task':
                    assert selfarg is not None
                    assert 'task' == selfarg['name']
                    callargs.append('task')
                elif a == 'dummy-value':
                    _,_,_,agtp = gaistype2type(self.jtis,p['type'])
                    callargs.append(DummyArg(agtp))
                elif a == 'return-value':
                    _,_,_,agtp = gaistype2type(self.jtis,p['type'])
                    if not isinstance(agtp,agref):
                        raise CodeGenerationError(f"Return value must be a reference in call to '{funname}'")
                    callargs.append(ReturnArg(agtp))
                    assert retval_t is None
                    retval_t = agtp.t
                else:
                    assert False
            else:
                callargs.append(self.__interpretexpr(a))

        if retval_t is None: raise CodeGenerationError(f"Missing 'return-value' in call to '{funname}''")
        value,code = self.gencall(func,*callargs)
        return ComputeNodeValue(retval_t,value,code)

    def __call__(self):
        return self.__interpretexpr(self.__code)

class BaseFuncGenerator:
    """
    This class is the base class for function generating. Each
    instance is meant to generate one instance of one function.

    To create a function generator, override the functions for
    handling different argument types and for generating the function
    code.
    """
    def __init__(self,
                 jtis,
                 cls,
                 func,
                 tmpvargen = None,
                 collectors = ['funarg','rettype','retval','decl','prelude','callarg','postlude'],
                 addcollectors = [],
                 useclassarg = True,
                 useapinames=True):
        """
        - jtis is the AIS structure
        - cls is the AIS Object Class in which the function being generated belongs
        - func  is the AIS function being generated
        - tmpvargen is a unique name generator
        - collectors is a list of 'collector' names each collector is a list and
          can be accessed as `self["name"]`. By default a basic list of entries
          are defined.
        - addcollectors Additional collectors
        - useclassarg True/False indicating if class arg ('this' or
          'self') should be treaded specially or just as another
          argument. Basically whether we call `arg_classarg()` or an
          otherwise appropriate `arg_...()` for the class argument.
        - useapinames Define if we use API names or only function names.
        """
        self.jtis = jtis
        self.__cls = cls
        self.__func = func
        self.__useapinames      = useapinames
        self.__useclassarg   = useclassarg
        self.__tmpvargen = tmpvargen or TempVarNameGenerator()

        self.__collectors = { k:[] for k in collectors+addcollectors }
        self.__funname    = self.__func['name']
        self.__apifunname = self.__func['api-name']
        self.__static     = all([not a.get('classarg') for a in self.__func['args']])

    def isstatic(self):
        return self.__static
    def get_tmpvargen(self):
        return self.__tmpvargen
    def tmpvargen(self,suffix=None):
        """
        Generate a unique name that can be used for variables and internal functions.
        """
        if suffix is None:
            return next(self.__tmpvargen)
        else:
            return next(self.__tmpvargen)+suffix

    def __getitem__(self,key):
        return self.__collectors[key]

    # Override to get messages
    def warn(self,msg,args): pass
    def error(self,msg,args): pass
    def info(self,msg,args): pass
    # Implement these to generate per-argument code
    #def ret(self,a,tp,atn,ctn,argbrief,argdesc):
    #    """
    #    Override to handle return argument
    #    """
    def arg_classarg(self,a,tp,atn,ctn,n,argbrief,argdesc):
        """
        Override to handle class arguments ('this' or 'self')
        """
        raise DontGenerate("arg_classarg() not implemented")
    def arg_ptr_enum(self,a,basetp,atn,ctn,cc,n,minlength,isdefaultoarg,argbrief,argdesc):
        """
        Override to handle arguments of type "ptr ENUM-TYPE".

        Optional. Default will just call `arg_ptr()`
        """
        self.arg_ptr(a,basetp,atn,ctn,n,minlength,None,isdefaultoarg,argbrief,argdesc)
    def arg_ptr_obj(self,a,basetp,atn,ctn,n,minlength,isdefaultoarg,argbrief,argdesc):
        """
        Override to handle arguments of type "ptr OBJECT".

        Optional. Default will just call `arg_ptr()`
        """
        self.arg_ptr(a,basetp,atn,ctn,n,minlength,None,isdefaultoarg,argbrief,argdesc)
    def arg_ptr_str(self,a,basetp,atn,ctn,n,minlength,argbrief,argdesc):
        """
        Override to handle arguments of type "ptr STRING".
        """
        self.arg_ptr(a,basetp,atn,ctn,n,minlength,None,False,argbrief,argdesc)
    def arg_ptr(self,a,basetp,atn,ctn,n,minlength,indexof,isdefaultoarg,argbrief,argdesc):
        """
        Override to handle arguments of type "ptr (ENUM|PRIMITIVE|OBJECT|STRING)".
        """
        raise DontGenerate("arg_ptr() not implemented")
    def arg_surpof(self,a,basetp,atn,ctn,n,surpof,argbrief,argdesc):
        """
        Override to handle "surp" arguments, which will always have
        type "ref (int32|int64)". The value of a surp argument can be
        computed as the length of some output array.
        """
        raise DontGenerate("arg_surp() not implemented")
    def arg_ref_enum(self,a,basetp,atn,ctn,cc,n,isdefaultoarg,argbrief,argdesc):
        """
        Override to handle arguments of type "(ref|out) ENUM-TYPE".

        Optional. Default will just call `arg_ref()`
        """
        self.arg_ref(a,basetp,atn,ctn,n,None,isdefaultoarg,argbrief,argdesc)
    def arg_ref(self,a,basetp,atn,ctn,n,indexof,isdefaultoarg,argbrief,argdesc):
        """
        Override to handle arguments of type "(ref|out) PRIMITIVE".
        """
        raise DontGenerate("arg_ref() not implemented")
    def arg_refobj(self,a,basetp,atn,ctn,n,isdefaultoarg,argbrief,argdesc):
        """
        Override to handle arguments of type "out OBJECT".
        """
        raise DontGenerate("arg_refobj() not implemented")
    def arg_refptr(self,a,basetp,atn,ctn,n,lenarg,minlength,isdefaultoarg,argbrief,argdesc):
        """
        Override to handle arguments of type "out ptr PRIMITIVE".
        """
        raise DontGenerate("arg_refobj() not implemented")
    def arg_instring(self,a,tp,atn,ctn,n,argbrief,argdesc):
        """
        Override to handle string input arguments.
        """
        raise DontGenerate("arg_instring() not implemented")
    def arg_outstring(self,a,tp,atn,ctn,n,minlength,isdefaultoarg,argbrief,argdesc):
        """
        Override to handle string output arguments.
        """
        raise DontGenerate("arg_outstring() not implemented")
    def arg_computed_value(self,a,tp,atn,ctn,n,lengthof,code,argbrief,argdesc):
        """
        Override to handle input values that are computed. These are
        always of primitive type "(int32|int64|float64)" or "ref
        (int32|int64|float64)".
        """
        raise DontGenerate("arg_outstring() not implemented")
    def arg_obj(self,a,tp,atn,ctn,n,argbrief,argdesc):
        """
        Override to handle input arguments of type object (Task or Env).
        """
        raise DontGenerate("arg_obj() not implemented")
    def arg_ref_func(self,a,tp,rettype,argtypes,n,isdefaultoarg,argbrief,argdesc):
        raise DontGenerate("arg_ref_func() not implemented")
    def arg_func(self,a,tp,rettype,argtypes,n,argbrief,argdesc):
        raise DontGenerate("arg_func() not implemented")
    def arg_value(self,a,tp,atn,ctn,n,argbrief,argdesc):
        raise DontGenerate("arg_value() not implemented")
    def arg_value_enum(self,a,tp,atn,ctn,cc,n,argbrief,argdesc):
        self.arg_value(a,tp,atn,ctn,n,argbrief,argdesc)
    def codegen(self,code):
        """
        Called to generate code for a snippet of AIS generic code.

        Must return a triple (tp,value,codelines)
          tp is an `agtype`
          value is a string
          codelines is a list of code lines with zero indentation
        """
        assert False
    # override this to generate function code
    def genfunc(self,out,func,funname,apiname,brief,desc):
        """
        Generate the actual function.

        This is called when all arguments have been processed. The
        result can be send to `out`, which is the user-defined object
        passed to `run()`.
        """
        assert False

    # call this to generate code.
    def run(self,out):
        """
        Execute the generator. The argument `out` is a user defined
        argument that can be used to gather the result.
        """
        try:
            self.__genargs()
            funbrief = self.__func['brief']
            fundesc  = self.__func.get('desc',funbrief)
            self.genfunc(out,self.__func,self.__funname,self.__apifunname,funbrief,fundesc)
        except DontGenerate as e:
            func = e.func if e.func is not None else self.__func['name']
            self.warn("Dont Generate %s: %s",func,e.reason)

    def __genargs(self):
        returntype  = self.__func['returns']['type']#.getReturnType ()
        #defaultoarg = self.__func.getDefaultOArg ()
        #classarg    = self.__func.getReplaces ()

        argsd = { a['name'] : a for a in self.__func['args'] }

        #print(self.__func['name'])
        for arg in self.__func['args']:
            try:
                t = resolveType(self.jtis,arg['type'])

                n = arg['name']

                argbrief = arg.get('brief')
                if not argbrief:
                    #argbrief = self.jtis['default-doc'].get(n,{}).get('brief')
                    argbrief = self.jtis.getDefaultArgBrief(n)
                argdesc  = arg.get('desc',argbrief)
                if not argdesc:
                    #argdesc = self.jtis['default-doc'].get(n,{}).get('desc',argbrief)
                    argdesc = self.jtis.getDefaultArgDesc(n) or argbrief

                compute  = arg.get('compute',[])
                minlength = None
                lengthof  = []
                lenarg    = None
                indexof   = []
                surpof    = []
                code      = None
                for c in compute:
                    ct = c['type']
                    if   ct == 'CODE':
                        assert code is None
                        code     = c.get('code')
                    elif ct == 'LENGTHOF':
                        lengthof = c['arg-refs']
                    elif ct == 'LENGTH':
                        lenarg = c['arg-refs'][0]
                    elif ct == 'INDEXOF':
                        indexof   = c['arg-refs']
                    elif ct == 'SURPOF':
                        surpof    = c['arg-refs']
                    elif ct == 'MINLENGTH':
                        minlength = c['code']

                if self.__useclassarg and arg.get('classarg'):
                    assert not code
                    assert not surpof
                    assert not lengthof
                    assert not lenarg
                    assert not indexof
                    atn,ctn,bt,gtn = aistype2type(self.jtis,arg['type'])
                    self.arg_classarg(arg,bt,atn,ctn,n,argbrief,argdesc)
                elif t[0] == 'ptr':
                    assert not code
                    assert not surpof
                    assert not surpof
                    isdefaultoarg = arg.get('defaultout',False)

                    if minlength is not None:
                        minlength = self.codegen(minlength)
                        assert isinstance(minlength,ComputeNodeValue),minlength
                    if lenarg is not None:
                        lenarg = argsd[lenarg]

                    tp = t[1]
                    if tp[0] in [ 'ptr','ref' ]:
                        raise DontGenerate(self.__funname,f"Argument '{n}' has type ptr<ref<...>> or ptr<ptr<...>>")

                    atn,ctn,bt,gtn = aistype2type(self.jtis,tp)
                    if bt[0] == 'const-class':
                        cc = self.jtis.getNamespace('mosek')['constclasses'][bt[2]]
                        assert not indexof
                        self.arg_ptr_enum(arg,bt,atn,ctn,cc,n,minlength,isdefaultoarg,argbrief,argdesc)
                    elif bt[0] == 'defined' and bt[2] in ['task_t','env_t']:
                        assert not indexof
                        self.arg_ptr_obj(arg,bt,atn,ctn,n,minlength,isdefaultoarg,argbrief,argdesc)
                    elif bt[0] == 'defined' and bt[2] in ['string_t']:
                        assert not indexof
                        # if arg['mode'] == 'i':
                        #     raise DontGenerate(f'Argument {n} : ptr(str) is input')
                        self.arg_ptr_str(arg,bt,atn,ctn,n,minlength,argbrief,argdesc)
                    elif gtn == agvoidp:
                        self.arg_value(arg,t,'voidp','void*',n,argbrief,argdesc)
                    else:
                        self.arg_ptr(arg,bt,atn,ctn,n,minlength,indexof,isdefaultoarg,argbrief,argdesc)
                elif t[0] == 'ref':
                    assert not code
                    assert not lengthof
                    assert not indexof
                    tt = t[1]
                    if surpof:
                        surpof = [ argsd[so] for so in surpof ]
                        atn,ctn,bt,gtn = aistype2type(self.jtis,tt)
                        self.arg_surpof(arg,bt,atn,ctn,n,surpof,argbrief,argdesc)
                    elif not tt[0] == 'ptr':
                        atn,ctn,bt,gtn = aistype2type(self.jtis,tt)
                        isdefaultoarg = arg.get('defaultout',False)
                        if bt[0] == 'func':
                            self.arg_ref_func(arg,bt,gtn.rettype,gtn.argtypes,n,isdefaultoarg,argbrief,argdesc)
                        elif atn in objtypes:
                            self.arg_refobj(arg,bt,atn,ctn,n,isdefaultoarg,argbrief,argdesc)
                        elif bt[0] == 'const-class':
                            cc = self.jtis.getNamespace('mosek')['constclasses'][bt[2]]
                            assert not indexof
                            self.arg_ref_enum(arg,bt,atn,ctn,cc,n,isdefaultoarg,argbrief,argdesc)
                        else:
                            self.arg_ref(arg,bt,atn,ctn,n,indexof,isdefaultoarg,argbrief,argdesc)
                    else: # ref ptr
                        atn,ctn,bt,gtn = aistype2type(self.jtis,tt[1])
                        assert minlength is not None or lenarg is not None
                        if minlength is not None:
                            minlength = self.codegen(minlength)
                        if lenarg is not None:
                            lenarg = argsd[lenarg]
                        self.arg_refptr(arg,bt,atn,ctn,n,lenarg,minlength,isdefaultoarg,argbrief,argdesc)
                elif t[0] == 'func':
                    atn,ctn,bt,gtn = aistype2type(self.jtis,t)
                    self.arg_func(arg,bt,gtn.rettype,gtn.argtypes,n,argbrief,argdesc)
                elif t[0] in ['ctype','defined','const-class']:
                    atn,ctn,bt,gtn = aistype2type(self.jtis,arg['type'])
                    if atn == 'string':
                        assert not code
                        assert not lengthof
                        assert not indexof
                        if arg['mode'] == 'i':
                            assert not lenarg
                            #assert not minlength, self.__func['name']
                            self.arg_instring(arg,bt,atn,ctn,n,argbrief,argdesc)
                        elif arg['mode'] == 'o':
                            if minlength is not None:
                                minlength = self.codegen(minlength)
                            isdefaultoarg = arg.get('defaultout',False)
                            if not minlength: raise CodeGenerationError("Missing minlength")
                            self.arg_outstring(arg,bt,atn,ctn,n,minlength,isdefaultoarg,argbrief,argdesc)
                        else:
                            raise DontGenerate(self.__funname,"I+O string")
                    elif code is not None or len(lengthof) > 0:
                        if lengthof:
                            lengthof = [ argsd[l] for l in lengthof ]
                        if code is not None:
                            code = self.codegen(code)
                        self.arg_computed_value(arg,bt,atn,ctn,n,lengthof,code,argbrief,argdesc)
                    elif atn in objtypes:
                        self.arg_obj(arg,bt,atn,ctn,n,argbrief,argdesc)
                    elif bt[0] == 'const-class':
                        cc = self.jtis.getNamespace('mosek')['constclasses'][bt[2]]
                        if indexof:
                            raise CodeGenerationError("indexof defined for enum type argument")
                        self.arg_value_enum(arg,bt,atn,ctn,cc,n,argbrief,argdesc)
                    else:
                        self.arg_value(arg,bt,atn,ctn,n,argbrief,argdesc)
                else:
                    raise DontGenerate(self.__funname,f"Unknown type '{t}' for argument {n}")
            except InvalidType as e:
                traceback.print_exc()
                raise DontGenerate(self.__funname,f"Invalid type of argument {n}: {e}")
            # except UnknownType as e:
            #     raise DontGenerate(self.__funname,f"Unknown type of argument {n}: {e}")
            except CodeGenerationError as e:
                raise DontGenerate(self.__funname,f"Invalid code field argument {n}: {e}")



class BaseAPIGenerator:
    def __init__(self,
                 jtis,
                 targetfilter = None, # None a string
                 useapinames=True,
                 tmpvargen = None):
        self.jtis = AIS(jtis)
        self.__useapinames = useapinames
        self.__targetfilter = targetfilter

        self.__tmpvargen = tmpvargen or TempVarNameGenerator()

    def get_tmpvargen(self):
        return self.__tmpvargen
    def warn(self,msg,args): pass
    def error(self,msg,args): pass
    def info(self,msg,args): pass

    def genenum(self,out,e,enumname,prefix,members,enumerable=True):
        """
        Called for each enum class to generate the enums code.

        Must be implemented.
        """
        pass
    def genfunc(self,out,func,cls):
        """
        Called for each function to generate the functions code

        Must be implemented.
        """
        pass

    def run(self,out):
        """
        The user calls this to execute the API generator.
        """
        allenums = [ (ccname,cc) for ccname,cc in self.jtis['constclasses'].items()
                     if cc['api-class'] != 'internal' and cc['value-type'] == 'int' ]
        for ccname,cc in allenums:
            members = cc['members']
            members.sort(key = lambda i:i['value'])

            self.genenum(out,cc,cc['name'],cc['prefix'].lower(),members,cc['is-enumerable'])

        for oc in self.jtis['classes'].values():
            classname = oc['name']
            if self.__targetfilter:
                allfuns = [ f for f in oc['functions'].values() if self.__targetfilter in f['targets'] ]
            else:
                allfuns = list(oc['functions'].values())
            allfuns.sort(key=lambda item: item['name'])
            for fun in allfuns:
                self.genfunc(out,fun,oc)

if __name__ == '__main__':
    # test
    import argparse,json,sys
    import logging
    A = argparse.ArgumentParser()
    A.add_argument('--out')
    A.add_argument('JSONais')



    class ComputeGen(BaseComputeGenerator):
        def __init__(self,code,tis,cls,func,tmpvargen):
            super().__init__(code,tis,cls,func,tmpvargen)

        def constref(self,cc,name):
            return cc['prefix']+name.upper()
        def ifthenelse(self,cond,thenitem,elseitem):
            return thenitem[0],f'({cond} ? {thenitem} : {elseitem})',cond[2]+thenitem[2],elseitem[2] # Decidedly NOT correct code
        def lengthof(self,arg):
            tp,value,code = arg
            return ComputeNodePartial(f'{value}.length',[])
        def longlengthof(self,arg):
            tp,value,code = arg
            return ComputeNodePartial(f'{value}.length',[])
        def sumof(self,arg):
            return ComputeNodePartial(f'sum({arg.value})',arg.code)
        def binop(self,op,lhs,rhs):
            return ComputeNodePartial(f'({lhs.value}{op}{rhs.value})',lhs.code+rhs.code)
        def cmp(self,op,lhs,rhs):
            return ComputeNodePartial(f'({lhs.value}{op}{rhs.value})',lhs.code+rhs.code)
        def cast(self,tp,arg):
            return ComputeNodePartial(f'{tp}({arg.value}))',arg.code)
        def gencall(self,func,*args):
            code = []
            aa = []
            for zz in args:
                try:
                    (_,a,c) = zz
                    code.extend(c)
                    aa.append(a)
                except:
                    aa.append(zz)
            aa = ','.join([ str(a) for a in aa])
            return ComputeNodePartial(f'{func["name"]}({aa})',code)

    class FuncGen(BaseFuncGenerator):
        def __init__(self,tis,cls,func,tmpvargen):
            self.__tmpvargen = tmpvargen
            super().__init__(tis,
                             cls,
                             func,
                             tmpvargen = tmpvargen)
            self.__cls = cls
            self.__func = func
        def warn(self,msg,*args): logging.warn(msg,*args)
        def error(self,msg,*args): logging.error(msg,*args)
        def info(self,msg,*args): logging.info(msg,*args)
        def codegen(self,code):
            return ComputeGen(code,self.jtis,self.__cls,self.__func,self.__tmpvargen)()

        def arg_classarg(self,a,tp,atn,ctn,n,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_ptr(self,a,basetp,atn,ctn,n,minlength,indexof,isdefaultoarg,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_surpof(self,a,basetp,atn,ctn,n,surpof,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_ref(self,a,basetp,atn,ctn,n,indexof,isdefaultoarg,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_refobj(self,a,basetp,atn,ctn,n,isdefaultoarg,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_refptr(self,a,basetp,atn,ctn,n,lenarg,minlength,isdefaultoarg,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_instring(self,a,tp,atn,ctn,n,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_outstring(self,a,tp,atn,ctn,n,minlength,isdefaultoarg,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_computed_value(self,a,tp,atn,ctn,n,lengthof,code,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_obj(self,a,tp,atn,ctn,n,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        def arg_value(self,a,tp,atn,ctn,n,argbrief,argdesc):
            self['callarg'].append((n,atn,argbrief))
        # override this to generate function code
        def genfunc(self,f,func,funname,apiname,brief,desc):
            args = ','.join([ f'{n} : {t}' for n,t,_ in self['callarg']])
            f.write(f'fun {funname}({args})\n')
            for n,_,brief in self['callarg']:
                f.write(f'  {n}\n')
                if brief is not None:
                    for l in [ l.strip() for l in brief.split('\n') ]:
                        f.write(f'    {l}\n')
                else:
                    f.write('    <uncodumented>\n')

    class APIGen(BaseAPIGenerator):
        def __init__(self,
                     jtis,
                     tmpvargen):
            self.__tmpvargen = tmpvargen
            super().__init__(jtis, tmpvargen = tmpvargen)

        def warn(self,msg,*args): logging.warn(msg,*args)
        def error(self,msg,*args): logging.error(msg,*args)
        def info(self,msg,*args): logging.info(msg,*args)
        def genenum(self,f,e,enumname,prefix,members,enumerable=True):
            f.write(f"enum {enumname}:\n")
            for mem in members:
                f.write(f'  {mem["name"]} = {mem["value"]}\n')
        def genfunc(self,f,func,cls):
            fn = FuncGen(self.jtis,cls,func,self.__tmpvargen)
            fn.run(f)



    a = A.parse_args()
    with open(a.JSONais,'rb') as f:
        jtis = json.load(f)

    apigen = APIGen(jtis,TempVarNameGenerator())
    if a.out is not None:
        with open(a.out,'wt') as f:
            apigen.run(f)
    else:
        apigen.run(sys.stdout)
