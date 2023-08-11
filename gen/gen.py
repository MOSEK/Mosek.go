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

atypemap = {
    'int32'        : 'int32',
    'int64'        : 'int64',
    'string'       : 'string',
    'float64'      : 'float64',
    'double'       : 'float64',
    'void'         : 'void',
    'usize'        : 'uint',
    'uint32'       : 'uint32',
    'uint64'       : 'uint64',
    'bool'         : 'bool',
    }

ffitypemap = {
    'int'          : 'int32',
    'uint'         : 'uint32',
    'int64'        : 'int64',
    'uint64'       : 'uint64',
    'double'       : 'float64',
    'string_t'     : 'cstring',

    'void'         : 'void',
    'int32_t'      : 'int32',
    'int64_t'      : 'int64',
    }

def atype2nimdecl(t):
    if   isinstance(t,ag.agref):
        return 'ref ' + atype2num(t.t)
    elif isinstance(t,ag.agptr):
        return 'ptr ' + atype2num(t.t)
    elif isinstance(t,ag.agenum):
        return t.t
    elif isinstance(t,ag.agtype):
        return atypemap[t.t]
    elif isinstance(t,str):
        return atypemap[t]
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
        code = []
        tmpvar = self.tmpvargen()
        code.extend(cond.code)
        code.append(f'var {tmpvar} {atype2nimdecl(thenitem.type)}')
        code.append(f'if {cond.value} {{')
        code.extend(['  ' + l for l in thenitem.code])
        code.append(f'  {tmpvar} = {thenitem.value}')
        code.append(f'}} else {{')
        code.extend(['  ' + l for l in elseitem.code])
        code.append(f'  {tmpvar} = {elseitem.value}')
        code.append( '}')

        return ag.ComputeNodePartial(tmpvar,code)
    def constref(self,cc,name):
        return cc['prefix'].lower()+name
    def lengthof(self,arg):
        return ag.ComputeNodePartial(f'cast[int32]({arg.value}.len)',arg.code)
    def longlengthof(self,arg):
        return ag.ComputeNodePartial(f'{arg.value}.len',arg.code)
    def sumof(self,arg):
        return ag.ComputeNodePartial(f'{arg.value}.foldl(a+b)',arg.code)
    def gencall(self,func,*args):
        callargs = []
        value = None
        code = []
        selfargs = [ a for a in self.__func['args'] if a.get('classarg') ]
        for a in args:
            if   a in ['self','task']:
                #n = self.__func.getReplaces()
                callargs.append(f'{selfargs[0]["name"]}.nativep')
            elif isinstance(a,ag.DummyArg) or isinstance(a,ag.ReturnArg):
                tmpvar = self.tmpvargen()
                if isinstance(a.type, ag.agref):
                    code.append(f'var {tmpvar} {atype2nimdecl(a.type.t)}')
                    callargs.append(f'addr({tmpvar})')
                else:
                    code.append('var {tmpvar} {atype2nimdecl(a.type)}')
                    callargs.append(f'{tmpvar}')

                if isinstance(a,ag.ReturnArg):
                    value = tmpvar
            else:
                code.extend(a.code)
                callargs.append(a.value)
        callargs = ','.join(callargs)
        tmpres = self.tmpvargen()
        code.extend([f'if {tmpres} := C.MSK_{func["name"]}({callargs}); {tmpres} != 0 {{',
                     f'  lastcode,lastmsg = self.getlasterror({tmpres})',
                     '   err = MosekError{ code:lastcode,msg:lastmsg}',
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
        self.__clsarg = a
        self.__clstp = atn
        #self['funarg'].append(f'{n} {objecttypemap[atn]}')
    def arg_ptr(self,a,basetp,atn,ctn,n,minlength,indexof,isdefaultoarg,argbrief,argdesc):
        if atn == 'enum':
            argtp = ctn.capitalize() 
        else:
            try:
                argtp = atype2nimdecl(atn)
            except KeyError:
                raise ag.DontGenerate(self.__func['name'],"Invalid type") 
            if argtp == 'void':
                raise ag.DontGenerate(self.__func['name'],"Void pointer arguments not supported")
        
        tmp = self.tmpvargen()

        self['prelude'].extend([f'var {tmp} *{argtp}'])

        if minlength:
            self['prelude'].extend(minlength.code)

        if 'i' in a['mode']:
            if minlength:
                self['prelude'].extend([f'if len({n}) < {minlength.value} {{',
                                        f'''  err = &ArrayLengthError{{fun:"{self.funapiname}",arg:"{n}"}}''',
                                        '  return',
                                        '}'])
            self['prelude'].append(f'if {n} != nil {{ {tmp} = (*C.MSKint32t)(&{n}[0]) }}')
            self['funarg'].append(f'{n} []{argtp}')

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
            self['prelude'].extend([f'{n} := make([]{argtp},{minlength.value})',
                                    f'if len({n}) > 0 {{ {tmp} = (*{argtp})(&n[0]) }}'])
            self['retarg'].append((n,f'[]{argtp}'))
        self['callarg'].append(tmp)
            
#    def arg_surpof(self,a,basetp,atn,ctn,n,surpof,argbrief,argdesc):
#        minlen = f'{surpof[0]["name"]}.len'
#        for sof in surpof[1:]:
#            minlen = f'min({minlen},{sof["name"]}.len)'
#        argtp = atype2nimdecl(atn)
#        self['prelude'].append(f'var {n} = cast[{argtp}]({minlen})')
#        self['callarg'].append(f'unsafeAddr({n})')
#        self['nativearg'].append(f'{n} : ptr {argtp}')
    def arg_ref(self,a,basetp,atn,ctn,n,indexof,isdefaultoarg,argbrief,argdesc):
        if atn == 'enum':
            argtp = ctn.capitalize()
        else:
            argtp = atype2nimdecl(atn)

        assert a['mode'] != 'io'
        self['retarg'].append((n,argtp))
        self['callarg'].append(f'&{n}')
    def arg_ref_func(self,a,tp,rettype,argtypes,n,isdefaultoarg,argbrief,argdesc):
        raise DontGenerate("arg_ref_func() not implemented")
    def arg_refobj(self,a,basetp,atn,ctn,n,isdefaultoarg,argbrief,argdesc):
        raise ag.DontGenerate(self.funname,"Not supported: Ref to object")
    def arg_refptr(self,a,basetp,atn,ctn,n,lenarg,minlength,isdefaultoarg,argbrief,argdesc):
        raise ag.DontGenerate(self.funname,"Not supported: Ref to pointer")
    def arg_instring(self,a,tp,atn,ctn,n,argbrief,argdesc):
        self['funarg'].append(f'{n} string')
        tmp = self.tmpvargen()
        self['prelude'].append(f'{tmp} := C.CString({n})')
        self['callarg'].append(tmp)
    def arg_outstring(self,a,tp,atn,ctn,n,minlength,isdefaultoarg,argbrief,argdesc):
        tmpvar1 = self.tmpvargen()
        self['prelude'].extend(minlength.code)
        self['prelude'].append(f'{tmpvar1} := make([]byte,{minlength.value})')
        self['callarg'].append(f'C.CString(&tmpvar1[0])')
        self['postlude'].extend([f'var {n} string',
                                 f"if p := strings.IndexByte({tmpvar1},byte(0)); p < 0 {{",
                                 f'  {n} = string({tmpvar1})',
                                 '} else {',
                                 f'  {n} = string({tmpvar1}[:p])',
                                 '}'])
        self['retarg'].append((n,'string'))
    def arg_computed_value(self,a,tp,atn,ctn,n,lengthof,code,argbrief,argdesc):
        argtp = atype2nimdecl(atn)
        if lengthof:
            tmp = self.tmpvargen()
            self['prelude'].append(f'{tmp} := len({lengthof[0]["name"]})')
            for lof in lengthof[1:]:
                self['prelude'].append(f'if {tmp} < {lof["name"]} {{ {tmp} = lof["name"] }}')

            if argtp == ag.agtype('int64'):
                self['prelude'].append(f'var {n} int64 = {tmp}')
            else:
                self['prelude'].append(f'var {n} {argtp} = int32({tmp})')
            self['callarg'].append(n)
        elif code:
            argtp = atype2nimdecl(atn)
            self['prelude'].extend(code.code)
            self['prelude'].append(f'var {n} {argtp} = {code.value}')

            self['callarg'].append(n)
        else:
            assert False
    def arg_obj(self,a,tp,atn,ctn,n,argbrief,argdesc):
        raise ag.DontGenerate(self.funname,"Not supported: Non-class-arg object arguments")
    def arg_value(self,a,tp,atn,ctn,n,argbrief,argdesc):
        if atn == 'enum':
            argtp = ctn.capitalize()
        else:
            argtp = atype2nimdecl(atn)

        self['funarg'].append(f'{n} {argtp}')
        self['callarg'].append(n)
    def genfunc(self,d,func,funname,apiname,brief,desc):
        if func['status'] == 'internal':
            raise ag.DontGenerate(self.__func['name'],"Invalid type") 

        callargs = ','.join(self['callarg'])
        funargs  = ','.join(self['funarg'])
        nfunargs = ','.join(self['nativearg'])


        rets = [ f'{n} {t}' for n,t in self['retarg'] ]
        rets.append('err error')

        if len(rets) == 0: retstr = ''
        else: retstr = ' (' + ','.join(rets) + ')'
        
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
                                 '    err = &MosekError{code:lastcode,msg:lastmsg}',
                                 '    return'])
        else:
            d['funimpl'].extend([f'    err = &MosekError{{ code:{tmp} }}',
                                 '    return'])
        d['funimpl'].append('  }')

        d['funimpl'].extend(['  '+l for l in self['postlude']])
        d['funimpl'].append('  return')
        d['funimpl'].append('}')

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
            L = [  (i['name'], int(i['value'])) for i in cc['members'] ]
            L.sort (key = lambda i: i[1])
            ccname = cc['name']
            pfx = cc['prefix'].upper()

            cname = ccname.capitalize()
            d['const'].append(f'type {cname} int32')
            d['const'].append('const (')
            for k,v in L:
                d['const'].append(f'    MSK_{pfx}{k.upper()} {cname} = {v}')
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
             'funimpl'    : ['// Methods']}

        APIGen(jtis,ag.TempVarNameGenerator(prefix="_tmt",postfix="",init=0)).run(d)

        p = 0
        for o in re.finditer(r'^//<([A-Za-z0-9]+)>.*$',template,re.MULTILINE):
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
            else:
                print(f"Warning! Unknown label '{o.group(1)}'")

        f.write(template[p:])

