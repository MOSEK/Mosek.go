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
        code.append(f'var {tmpvar} : {atype2nimdecl(thenitem.type)}')
        code.append(f'if {cond.value}:')
        code.extend(['  ' + l for l in thenitem.code])
        code.append(f'  {tmpvar} = {thenitem.value}')
        code.append(f'else:')
        code.extend(['  ' + l for l in elseitem.code])
        code.append(f'  {tmpvar} = {elseitem.value}')

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
                    code.append(f'var {tmpvar} : {atype2nimdecl(a.type.t)}')
                    callargs.append(f'addr({tmpvar})')
                else:
                    code.append('var {tmpvar} : {atype2nimdecl(a.type)}')
                    callargs.append(f'{tmpvar}')

                if isinstance(a,ag.ReturnArg):
                    value = tmpvar
            else:
                code.extend(a.code)
                callargs.append(a.value)
        callargs = ','.join(callargs)
        code.append(f'handle_res(MSK_{func["name"]}({callargs}))')
        return ag.ComputeNodePartial(value,code)

class FuncGen(ag.BaseFuncGenerator):
    def __init__(self,tis,cls,func,tmpvargen):
        self.__tmpvargen = tmpvargen
        super().__init__(tis,
                         cls,
                         func,
                         tmpvargen = tmpvargen,
                         addcollectors = ['nativearg'])
        self.__cls = cls
        self.__func = func
        self.funname = func['name']
        self.funapiname = func.get('api-name',self.funname)
    def warn(self,msg,*args): logging.warning(msg,*args)
    def error(self,msg,*args): logging.error(msg,*args)
    def info(self,msg,*args): logging.info(msg,*args)
    def codegen(self,code):
        return ComputeGen(code,self.jtis,self.__cls,self.__func,self.__tmpvargen)()
    def arg_classarg(self,a,tp,atn,ctn,n,argbrief,argdesc):
        self['callarg'].append(f'self.ptr()')
        self['selfarg'] = f'self *{objecttypemap[atn]}'
        #self['funarg'].append(f'{n} {objecttypemap[atn]}')
    def arg_ptr(self,a,basetp,atn,ctn,n,minlength,indexof,isdefaultoarg,argbrief,argdesc):
        if atn == 'enum':
            argtp = 'int32' 
        else:
            argtp = atype2nimdecl(atn)
            if argtp == 'void':
                raise ag.DontGenerate(self.__func['name'],"Void pointer arguments not supported")
        
        tmp = self.tmpvargen()
        
        assert minlength
        self['prelude'].extend([f'var {tmp} *{argtp}'])
        self['prelude'].extend(minlength.code)

        if 'i' in a['mode']:
            self['prelude'].extend([f'if len({n}) < {minlength.value} {{',
                                    f'''  err = &ArrayLengthError{fun:"{self.funapiname}",arg:"{n}"}''',
                                    '  return',
                                    '}',
                                    f'if {n} != nil {{ {tmp} = (*MSKint32t)(&{n}[0]) }}'])
            self['funarg'].append(f'{n} []{argtp}')

            if indexof:
                tmpvar = self.tmpvargen()
                ixcheck = ' || '.join([f'{tmpvar} >= len({ix})' for ix in indexof])
                self['prelude'].extend([f'for _,{tmpvar} := range {n} {{',
                                        f'  if {tmpvar} < 0 || {ixcheck} {{',
                                        f'    err = &ArrayLengthError("{self.funapiname}","{n}")',
                                        '    return',
                                        '  }',
                                        '}'])
        else:
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
            argtp = 'int32'
        else:
            argtp = atype2nimdecl(atn)

        assert a['mode'] != 'io'
        self['retarg'].append((n,argtp))
_
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
        self['prelude'].append(f'{tmpvar1} := make([[]byte,{minlength.value})')
        self['callarg'].append(f'C.CString(&tmpvar1[0])')
        self['postlude'].extend([f'var {n} string',
                                 f"if p := strings.IndexByte({tmpvar1},'\\0'); p < 0 {{",
                                 f'  {n} = string({tmpvar1})',
                                 '} else {',
                                 f'  {n} = string({tmpvar1}[:p])',
                                 '}'])
        self['retarg'].append((n,'string'))
    def arg_computed_value(self,a,tp,atn,ctn,n,lengthof,code,argbrief,argdesc):
        argtp = atype2nimdecl(atn)
        if lengthof:
            minlen = f'len({lengthof[0]["name"]})'
            for lof in lengthof[1:]:
                minlen = f'min({minlen},{lof["name"]}.len)'

            if argtp == ag.agtype('int64'):
                self['prelude'].append(f'var {n} : {argtp} = {minlen}')
            else:
                self['prelude'].append(f'var {n} : {argtp} = cast[int32]({minlen})')
            self['callarg'].append(n)
        elif code:
            argtp = atype2nimdecl(atn)
            self['prelude'].extend(code.code)
            self['prelude'].append(f'var {n} : {argtp} = {code.value}')

            self['callarg'].append(n)
        else:
            assert False
        self['nativearg'].append(f'{n} : {argtp}')
    def arg_obj(self,a,tp,atn,ctn,n,argbrief,argdesc):
        raise ag.DontGenerate(self.funname,"Not supported: Non-class-arg object arguments")
        assert False
    def arg_value(self,a,tp,atn,ctn,n,argbrief,argdesc):
        if atn == 'enum':
            if self.jtis["constclasses"][ctn]['is-enumerable']:
                argtp = ctn
            else:
                argtp = 'int32'
        else:
            argtp = atype2nimdecl(atn)

        self['funarg'].append(f'{n} : {argtp}')
        self['callarg'].append(n)
        self['nativearg'].append(f'{n} : {argtp}')
    def genfunc(self,d,func,funname,apiname,brief,desc):
        callargs = ','.join(self['callarg'])
        funargs  = ','.join(self['funarg'])
        nfunargs = ','.join(self['nativearg'])

        assert len(self['retval']) in [0,1]
        if self['retval']: retstr = ' : '+self['rettype'][0]
        else: retstr = ''

        d['nativehdrs'].append(f'proc MSK_{funname}({nfunargs}) : rescode {{. cdecl, importc .}}')
        d['funimpl'].append(f'proc {self.funapiname}*({funargs}){retstr} =')
        d['funimpl'].extend(['  '+l for l in self['prelude']])
        if func['args'] and func['args'][0]['name'] == 'task' and func['args'][0]['classarg']:
            d['funimpl'].append(f'  handle_res(task,MSK_{self.funname}({callargs}))')
        else:
            d['funimpl'].append(f'  handle_res(MSK_{self.funname}({callargs}))')
        d['funimpl'].extend(['  '+l for l in self['postlude']])
        if self['retval']:
            d['funimpl'].append(f'  return {self["retval"][0]}')

class APIGen(ag.BaseAPIGenerator):
    def __init__(self,
                 tis,
                 tmpvargen):
        self.__tmpvargen = tmpvargen
        super().__init__(tis, tmpvargen = tmpvargen, targetfilter='dotnet')

    def warn(self,msg,*args): logging.warning(msg,*args)
    def error(self,msg,*args): logging.error(msg,*args)
    def info(self,msg,*args): logging.info(msg,*args)
    def genenum(self,d,cc,enumname,prefix,members,enumerable=True):
        if cc['value-type'] == 'int':
            L = [  (i['name'], int(i['value'])) for i in cc['members'] ]
            L.sort (key = lambda i: i[1])
            ccname = cc['name']
            pfx = cc['prefix'].upper()

            d['const'].append('const (')
            for k,v in L:
                d['const'].append('    MSK_{pfx}{k} int32 = {v}')
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

        d = {'const'      : [],
             'nativehdrs' : [],
             'funimpl'    : [],
             'inftypes'   : []}

        APIGen(jtis,ag.TempVarNameGenerator(prefix="AUX",postfix="",init=0)).run(d)

        f.write('# Generated for MOSEK %s' % mosek_version)
        p = 0
        for o in re.finditer(r'^##<([A-Z0-9]+)>.*$',template,re.MULTILINE):
            f.write(template[p:o.start()])
            p = o.end()
            if o.group(1) == 'ENUMS':
                f.write('type\n')
                for l in d['enum']:
                    f.write('  '+l)
                    f.write('\n')
                f.write('\n')
                f.write('const\n')
                for l in d['const']:
                    f.write('  '+l)
                    f.write('\n')
            elif o.group(1) == 'FFIDECL':
                for l in d['nativehdrs']:
                    f.write(l)
                    f.write('\n')
                f.write('\n')
            elif o.group(1) == 'PROCS':
                for l in d['funimpl']:
                    f.write(l)
                    f.write('\n')
                f.write('\n')
            elif o.group(1) == 'INFTYPES':
                for l in d['inftypes']:
                    f.write('  ')
                    f.write(l)
                    f.write('\n')
        f.write(template[p:])

