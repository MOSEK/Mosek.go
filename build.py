import sys
if sys.version_info.major != 3:
    raise ValueError("Requires Python 3.x")

import os,os.path,platform
import subprocess
import http.client
import urllib.parse,urllib.request
import zipfile,tarfile
import shutil


base = os.path.abspath(os.path.dirname(__file__))

class NotSupported(Exception): pass

def urlget(url, tgtdir, tgtname=None):
    if tgtname is None:      
        urlfrag = urllib.parse.urlparse(url)
        p = urlfrag.path.rfind('/')
        if p > 0:
            tgtname = urlfrag.path[p+1:]
        if not tgtname:
            tgtname = 'default'
    
    target = os.path.join(tgtdir,tgtname) 

    with urllib.request.urlopen(url) as inf:
        with open(target,'wb') as outf:
            shutil.copyfileobj(inf,outf)

if   platform.system() == 'Linux':
    arch = subprocess.check_output(['uname','-p']).decode('ascii').strip()
    if arch == 'x86_64':
        pfname = 'linux64x86'
    elif arch == 'x86':
        pfname = 'linux32x86'
    else:
        raise NotSupported('Architecture "%s" not supported' % arch)
    distroname = 'mosektools'+pfname+'.tar.bz2'       
elif platform.system() == 'Darwin':
    pfname = 'osx64x86'
    distroname = 'mosektools'+pfname+'.tar.bz2'
else:
    raise NotSupported('Platform "%s" not supported' % system.platform())
     
if not os.path.isdir(os.path.join(base,'deps')):
    os.mkdir(os.path.join(base,'deps'))

if not os.path.isfile(os.path.join(base,'deps',distroname)):
    print("Fetch MOSEK distro...")
    urlget('http://download.mosek.com/stable/7/mosektoolslinux64x86.tar.bz2',
           os.path.join(base,'deps'))

if not os.path.isdir(os.path.join(base,'deps','mosek','7','tools','platform',pfname)):
    print("Unzip MOSEK distro '%s'" % distroname)
    with tarfile.open(os.path.join(base,'deps',distroname),'r:bz2') as f:
        membs = [ mem for mem in f.getmembers() if mem.name.startswith('mosek/7/tools/platform') ]
        f.extractall(path=os.path.join(base,'deps'),members=membs)
env = {}
env.update(os.environ)
env['CGO_CFLAGS']      = '-I{0}/mosek/7/tools/platform/{1}/h'.format(base,pfname)
env['CGO_LDFLAGS']     = '-L{0}/mosek/7/tools/platform/{1}/bin'.format(base,pfname)
env['LD_LIBRARY_PATH'] = '{0}/mosek/7/tools/platform/{1}/bin'.format(base,pfname)
env['GOPATH']          = base

if len(sys.argv) > 1:
    print("Execute: go",' '.join(sys.argv[1:]))
    subprocess.call(['go']+sys.argv[1:],env=env,cwd=base)
else:
    print("Arguments: build|install mosek")
    print("Arguments: build|install examples/...")
