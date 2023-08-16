import apigenjson as ag
import json,re
import argparse


REQUIRE_JAIS = (1,0,0)

objecttypemap = {
    'Task'       : 'Task',
    'Env'        : 'Env',
    }

typemap = {
    'int32'        : 'int32',
    'int'          : 'int32',
    'uint'         : 'uint32',
    'int64'        : 'int64',
    'uint64'       : 'uint64',
    'double'       : 'float64',
    'string_t'     : 'string',
    'float64'      : 'float64',

    'void'         : 'void',
    'int32_t'      : 'int32',
    'int64_t'      : 'int64',
    }
t2cmap = {
    'int32'  : 'C.int32_t',   
    'int64'  : 'C.int64_t',
    'float64': 'C.double',
    'bool': 'C.int',
    }

atypemap = {
    'int32'        : 'C.int32_t',
    'int64'        : 'C.int64_t',
    'string'       : '* C.char',
    'float64'      : 'C.double',
    'double'       : 'C.double',
    'bool'         : 'C.int',
    }
gtypemap = {
    'int32'        : 'int32',
    'int64'        : 'int64',
    'string'       : 'string',
    'float64'      : 'float64',
    'double'       : 'float64',
    'bool'         : 'bool',
    }

def atype2cgodecl(t):
    if   isinstance(t,ag.agref):
        return '*' + atype2cgodecl(t.t)
    elif isinstance(t,ag.agptr):
        return '*' + atype2cgodecl(t.t)
    elif isinstance(t,ag.agenum):
        return 'C.int32_t'
    elif isinstance(t,ag.agtype):
        return atypemap[t.t]
    elif isinstance(t,str):
        return atypemap[t]
    else:
        assert False,t
def atype2godecl(t):
    if   isinstance(t,ag.agref):
        return '*' + atypecgodecl(t.t)
    elif isinstance(t,ag.agptr):
        return '[]' + atypecgodecl(t.t)
    elif isinstance(t,ag.agenum):
        return t.t 
    elif isinstance(t,ag.agtype):
        return gtypemap[t.t]
    elif isinstance(t,str):
        return gtypemap[t]
    else:
        assert False,t


class ComputeGen(ag.BaseComputeGenerator):
    def __init__(self,code,tis,cls,func,tmpvargen):
        super().__init__(code,tis,cls,func,tmpvargen)
        self.__func = func

    def binop(self,op,lhs,rhs):
        code = []
        code.extend(lhs.code)
        code.extend(rhs.code)
        return ag.ComputeNodePartial(f'({lhs.value} {op} {rhs.value})',code)
    def cmp(self,op,lhs,rhs):
        code = []
        code.extend(lhs.code)
        code.extend(rhs.code)
        return ag.ComputeNodePartial(f'({lhs.value} {op} {rhs.value})',code)
    def cast(self,tp,arg):
        code = []
        code.extend(arg.code)
        return ComputeNodePartial(f'cast[{tp.t}]({arg.value})',code)
    def ifthenelse(self,cond,thenitem,elseitem):
        ctp = atype2cgodecl(thenitem.type)
        code = []
        tmpvar = self.tmpvargen()
        code.extend(cond.code)
        code.append(f'var {tmpvar} {atype2cgodecl(thenitem.type)}')
        code.append(f'if {cond.value} {{')
        code.extend(['  ' + l for l in thenitem.code])
        code.append(f'  {tmpvar} = ({ctp})({thenitem.value})')
        code.append(f'}} else {{')
        code.extend(['  ' + l for l in elseitem.code])
        code.append(f'  {tmpvar} = ({ctp})({elseitem.value})')
        code.append( '}')

        return ag.ComputeNodePartial(tmpvar,code)
    def constref(self,cc,name):
        return f"MSK_{cc['prefix']}{name}".upper()
    def lengthof(self,arg):
        return ag.ComputeNodePartial(f'len({arg.value})',arg.code)
    def longlengthof(self,arg):
        return ag.ComputeNodePartial(f'len({arg.value})',arg.code)
    def sumof(self,arg):
        return ag.ComputeNodePartial(f'sum({arg.value})',arg.code)
    def gencall(self,func,*args):
        callargs = []
        value = None
        code = []
        selfargs = [ a for a in self.__func['args'] if a.get('classarg') ]
        for a in args:
            if   a in ['self','task']:
                #n = self.__func.getReplaces()
                callargs.append(f'self.ptr()')
            elif isinstance(a,ag.DummyArg) or isinstance(a,ag.ReturnArg):
                tmpvar = self.tmpvargen()
                if isinstance(a.type, ag.agref):
                    code.append(f'var {tmpvar} {atype2cgodecl(a.type.t)}')
                    callargs.append(f'(&{tmpvar})')
                else:
                    code.append('var {tmpvar} {atype2cgodecl(a.type)}')
                    ctp = atype2cgodecl(a.type.t) 
                    callargs.append(f'({ctp})({tmpvar})')

                if isinstance(a,ag.ReturnArg):
                    value = tmpvar
            else:
                ctp = atype2cgodecl(a.type) 
                code.extend(a.code)
                callargs.append(f'({ctp})({a.value})')
        callargs = ','.join(callargs)
        tmpres = self.tmpvargen()
        code.extend([f'if {tmpres} := C.MSK_{func["name"]}({callargs}); {tmpres} != 0 {{',
                     f'  lastcode,lastmsg := self.getlasterror({tmpres})',
                     '   err = &MosekError{ code:Rescode(lastcode),msg:lastmsg}',
                     '  return',
                     '}'])
        return ag.ComputeNodePartial(value,code)

class FuncGen(ag.BaseFuncGenerator):
    def __init__(self,tis,cls,func,tmpvargen):
        self.__tmpvargen = tmpvargen
        super().__init__(tis,
                         cls,
                         func,
                         tmpvargen = tmpvargen,
                         addcollectors = ['nativearg','retarg'])
        self.__cls = cls
        self.__clsarg = None
        self.__clstp = None
        self.__func = func
        self.__selfarg = None
        self.funname = func['name']
        self.funapiname = re.sub('-.',lambda o: o.group(0)[1].upper(),func.get('api-caml-name').capitalize())

    def warn(self,msg,*args): logging.warning(msg,*args)
    def error(self,msg,*args): logging.error(msg,*args)
    def info(self,msg,*args): logging.info(msg,*args)
    def codegen(self,code):
        return ComputeGen(code,self.jtis,self.__cls,self.__func,self.__tmpvargen)()
    def arg_classarg(self,a,tp,atn,ctn,n,argbrief,argdesc):
        self['callarg'].append(f'self.ptr()')
        self.__selfarg = f'self *{objecttypemap[atn]}'
        self['nativearg'].append(f'{ctn}')
        self.__clsarg = a
        self.__clstp = atn
        #self['funarg'].append(f'{n} {objecttypemap[atn]}')
    def arg_ptr(self,a,basetp,atn,ctn,n,minlength,indexof,isdefaultoarg,argbrief,argdesc):
        if atn == 'enum':
            self['nativearg'].append('int32_t *')
            argtp = ctn.capitalize() 
            cctype = '*C.int32_t' 
        elif atn == 'string':
            cctype = f'**C.char' 
            self['nativearg'].append(f'const char **')
            argtp = 'string'
        else:
            cctype = f'*C.{ctn}' 
            self['nativearg'].append(f'{ctn} *')
            try:
                argtp = atype2godecl(atn)
            except KeyError:
                raise ag.DontGenerate(self.__func['name'],"Invalid type") 
            if argtp == 'void':
                raise ag.DontGenerate(self.__func['name'],"Void pointer arguments not supported")
        
        tmp = self.tmpvargen()

        if minlength:
            self['prelude'].extend(minlength.code)

        if 'i' in a['mode']:
            if minlength:
                self['prelude'].extend([f'if int64(len({n})) < int64({minlength.value}) {{',
                                        f'''  err = &ArrayLengthError{{fun:"{self.funapiname}",arg:"{n}"}}''',
                                        '  return',
                                        '}'])
            if atn != 'string':
                self['prelude'].extend([f'var {tmp} *{argtp}'])
                self['prelude'].append(f'if len({n}) > 0 {{ {tmp} = (*{argtp})(&{n}[0]) }}')
            else:
                cctype = '(**C.char)' 
                tmpl = self.tmpvargen()
                i = self.tmpvargen()
                v = self.tmpvargen()
                self['prelude'].extend([f'{tmpl} := make([]*C.char,len({n}))',
                                        f'for {i},{v} := range {n} {{',
                                        f'  {tmpl}[{i}] = C.CString({v})',
                                        '}',
                                        f'var {tmp} **C.char;',
                                        f'if len({tmpl}) > 0 {{ {tmp} = &{tmpl}[0] }}',
                                        ])

            self['funarg'].append((f'{n} []{argtp}',argbrief))

            if indexof:
                tmpvar = self.tmpvargen()
                ixcheck = ' || '.join([f'{tmpvar} >= len({ix})' for ix in indexof])
                self['prelude'].extend([f'for _,{tmpvar} := range {n} {{',
                                        f'  if {tmpvar} < 0 || {ixcheck} {{',
                                        f'    err = &ArrayLengthError{{"{self.funapiname}","{n}"}}',
                                        '    return',
                                        '  }',
                                        '}'])
        else:
            assert minlength
            self['prelude'].extend([f'var {tmp} *{argtp}'])
            self['prelude'].extend([f'{n} = make([]{argtp},{minlength.value})',
                                    f'if len({n}) > 0 {{ {tmp} = (*{argtp})(&{n}[0]) }}'])
            self['retarg'].append((n,f'[]{argtp}',argbrief))

        self['callarg'].append(f'({cctype})({tmp})')

    def arg_ref(self,a,basetp,atn,ctn,n,indexof,isdefaultoarg,argbrief,argdesc):
        assert a['mode'] != 'io'
        if atn == 'enum':
            self['nativearg'].append(f'int32_t *')
            argtp = ctn.capitalize()
            tmp = self.tmpvargen()
            self['prelude'].append(f'var {tmp} int32')
            self['callarg'].append(f'(*C.int32_t)(&{tmp})')
            self['postlude'].append(f'{n} = {argtp}({tmp})')
        elif atn == 'bool': 
            self['nativearg'].append(f'int *')
            argtp = atype2godecl(atn)
            tmp = self.tmpvargen()
            self['prelude'].append(f'var {tmp} C.int')
            self['callarg'].append(f'(&{tmp})')
            self['postlude'].append(f'{n} = {tmp} != 0')
        else:
            self['nativearg'].append(f'{ctn} *')
            argtp = atype2godecl(atn)
            self['callarg'].append(f'(*C.{ctn})(&{n})')
        self['retarg'].append((n,argtp,argbrief))
    def arg_ref_func(self,a,tp,rettype,argtypes,n,isdefaultoarg,argbrief,argdesc):
        raise DontGenerate("arg_ref_func() not implemented")
    def arg_refobj(self,a,basetp,atn,ctn,n,isdefaultoarg,argbrief,argdesc):
        raise ag.DontGenerate(self.funname,"Not supported: Ref to object")
    def arg_refptr(self,a,basetp,atn,ctn,n,lenarg,minlength,isdefaultoarg,argbrief,argdesc):
        raise ag.DontGenerate(self.funname,"Not supported: Ref to pointer")
    def arg_instring(self,a,tp,atn,ctn,n,argbrief,argdesc):
        self['nativearg'].append(f'const char *')
        self['funarg'].append((f'{n} string',argbrief))
        tmp = self.tmpvargen()
        self['prelude'].append(f'{tmp} := C.CString({n})')
        self['callarg'].append(tmp)
    def arg_outstring(self,a,tp,atn,ctn,n,minlength,isdefaultoarg,argbrief,argdesc):
        self['nativearg'].append(f'unsigned char *')
        tmpvar1 = self.tmpvargen()
        self['prelude'].extend(minlength.code)
        self['prelude'].append(f'{tmpvar1} := make([]byte,{minlength.value})')
        self['callarg'].append(f'(*C.uchar)(&{tmpvar1}[0])')
        self['postlude'].extend([f"if p := bytes.IndexByte({tmpvar1},byte(0)); p < 0 {{",
                                 f'  {n} = string({tmpvar1})',
                                 '} else {',
                                 f'  {n} = string({tmpvar1}[:p])',
                                 '}'])
        self['retarg'].append((n,'string',argbrief))
    def arg_computed_value(self,a,tp,atn,ctn,n,lengthof,code,argbrief,argdesc):
        self['nativearg'].append(f'{ctn}')
        argtp = atype2godecl(atn)
        if lengthof:
            tmp = self.tmpvargen()
            self['prelude'].append(f'{tmp} := len({lengthof[0]["name"]})')
            for lof in lengthof[1:]:
                self['prelude'].append(f'if {tmp} < len({lof["name"]}) {{ {tmp} = len({lof["name"]}) }}')

            if argtp == ag.agtype('int64'):
                self['prelude'].append(f'var {n} int64 = int64({tmp})')
            else:
                self['prelude'].append(f'var {n} {argtp} = {argtp}({tmp})')
            self['callarg'].append(f'C.{ctn}({n})')
        elif code:
            argtp = atype2godecl(atn)
            self['prelude'].extend(code.code)
            self['prelude'].append(f'var {n} {argtp} = {argtp}({code.value})')

            self['callarg'].append(f'C.{ctn}({n})')
        else:
            assert False
    def arg_obj(self,a,tp,atn,ctn,n,argbrief,argdesc):
        raise ag.DontGenerate(self.funname,"Not supported: Non-class-arg object arguments")
    def arg_value(self,a,tp,atn,ctn,n,argbrief,argdesc):
        if atn == 'enum':
            self['nativearg'].append('int32_t')
            argtp = ctn.capitalize()
            self['callarg'].append(f'C.int32_t({n})')
        elif atn == 'bool':
            self['nativearg'].append(f'int')
            argtp = atype2godecl(atn)
            tmp = self.tmpvargen()
            self['prelude'].append(f'var {tmp} C.int; if {n} {{ {tmp} = 1; }}')
            self['callarg'].append(f'{tmp}')
        else:
            self['nativearg'].append(f'{ctn}')
            argtp = atype2godecl(atn)
            self['callarg'].append(f'C.{ctn}({n})')

        self['funarg'].append((f'{n} {argtp}',argbrief))
    def genfunc(self,d,func,funname,apiname,brief,desc):
        if func['status'] == 'internal':
            raise ag.DontGenerate(self.__func['name'],"Invalid type") 

        callargs = ','.join(self['callarg'])
        funargs  = ','.join([ a for (a,b) in self['funarg']])
        nfunargs = ','.join(self['nativearg'])

        rets = [ f'{n} {t}' for n,t,d in self['retarg'] ]
        rets.append('err error')

        if len(rets) == 0: retstr = ''
        else: retstr = ' (' + ','.join(rets) + ')'
        
        
        d['funimpl'].append('')           
        d['funimpl'].extend([ f'// {l}' for l in brief.split('\n')])
        if self['funarg'] or self['retarg']:
            d['funimpl'].append('//')           
            for n,doc in self['funarg'] + [ (f'{n} {t}',doc) for n,t,doc in self['retarg']]:
                d['funimpl'].append(f'// - {n}')
                
                L = doc.split('\n')
                d['funimpl'].extend([f'//   {l}' for l in L])

        if self.__clsarg is not None:
            d['funimpl'].append(f'func (self *{self.__clstp}) {self.funapiname}({funargs}){retstr} {{')
        else:
            d['funimpl'].append(f'func {self.funapiname}({funargs}){retstr} {{')


        d['funimpl'].extend(['  '+l for l in self['prelude']])

        tmp = self.tmpvargen()
        argstr = ','.join(self["callarg"])
        d['funimpl'].append( f'  if {tmp} := C.MSK_{self.funname}({argstr}); {tmp} != 0 {{')
        if self.__clsarg is not None:
            d['funimpl'].extend([f'    lastcode,lastmsg := self.getlasterror({tmp})',
                                 '    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}',
                                 '    return'])
        else:
            d['funimpl'].extend([f'    err = &MosekError{{ code:Rescode({tmp}) }}',
                                 '    return'])
        d['funimpl'].append('  }')

        d['funimpl'].extend(['  '+l for l in self['postlude']])
        d['funimpl'].append('  return')
        d['funimpl'].append('}')


        nargs = ','.join(self['nativearg'])
        d['extern'].append(f'extern int MSK_{self.funname}({nargs});')

class APIGen(ag.BaseAPIGenerator):
    def __init__(self,
                 tis,
                 tmpvargen):
        self.__tmpvargen = tmpvargen
        super().__init__(tis, tmpvargen = tmpvargen, targetfilter='go')

    def warn(self,msg,*args): logging.warning(msg,*args)
    def error(self,msg,*args): logging.error(msg,*args)
    def info(self,msg,*args): logging.info(msg,*args)
    def genenum(self,d,cc,enumname,prefix,members,enumerable=True):
        if cc['value-type'] == 'int':
            L = [  (i['name'], int(i['value']),i.get('brief')) for i in cc['members'] ]
            L.sort (key = lambda i: i[1])
            ccname = cc['name']
            pfx = cc['prefix'].upper()

            cname = ccname.capitalize()
            d['const'].append('')
            d['const'].append(f'type {cname} int32')
            if 'brief' in cc:
                d['const'].extend([ '// '+l for l in cc['brief'].split('\n')])
            d['const'].append('const (')
            for k,v,doc in L:
                if doc is None: doc = ''
                else: doc = ' // ' + ' '.join(doc.split('\n'))
                d['const'].append(f'    MSK_{pfx}{k.upper()} {cname} = {v}{doc}')
            if pfx:
                d['const'].append(f'    MSK_{pfx}END {cname} = {v}')
            d['const'].append(')')

    def genfunc(self,dfunc,func,cls):
        fn = FuncGen(self.jtis,cls,func,self.__tmpvargen)
        fn.run(d)


if __name__ == '__main__':
    import  argparse
    import logging,os,pathlib

    A = argparse.ArgumentParser()
    A.add_argument('--target','-o',required=True)
    A.add_argument('--template',default=pathlib.Path(__file__).parent.joinpath('Mosek.template.go'))
    A.add_argument('JAIS')
    a = A.parse_args()

    with open(a.JAIS,'rb') as jf: jtis = json.load(jf)
    with open(a.template,'rt') as tf: template = tf.read()

    mosek_version = jtis['mosek-version']
    assert REQUIRE_JAIS <= tuple(jtis['jsonais-version'])

    with open(a.target,'wt') as f:
        print(f"Write to {a.target}")

        d = {'const'      : ['// Constants'],
             'funimpl'    : ['// Methods'],
             'extern'     : []}

        APIGen(jtis,ag.TempVarNameGenerator(prefix="_tmp",postfix="",init=0)).run(d)

        p = 0
        for o in re.finditer(r'^//<([A-Za-z0-9]+)>.*$\n',template,re.MULTILINE):
            f.write(template[p:o.start()])
            p = o.end()
            if o.group(1) == 'consts':
                for l in d['const']:
                    f.write(l)
                    f.write('\n')
            elif o.group(1) == 'methods':
                for l in d['funimpl']:
                    f.write(l)
                    f.write('\n')
                f.write('\n')
            elif o.group(1) == 'comment':
                f.write('// Generated for MOSEK %s' % mosek_version)
            elif o.group(1) == 'extern':
                for l in d['extern']:
                    f.write('// ')
                    f.write(l)
                    f.write('\n')
            else:
                print(f"Warning! Unknown label '{o.group(1)}'")

        f.write(template[p:])

