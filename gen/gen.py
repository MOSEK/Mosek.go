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

