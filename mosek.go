
/* MOSEK package.

   Thin wrapper for MOSEK solver API.
*/
package mosek

// #include <stdlib.h>
// #include <stdint.h>
// #cgo LDFLAGS: -lmosek64
//
// typedef void * MSKuserhandle_t;
// typedef const int8_t * string_t;
// typedef void * MSKtask_t;
// typedef void * MSKenv_t;
// typedef void  (* MSKstreamfunc) ( uintptr_t, const string_t);
// typedef int32_t  (* MSKcallbackfunc) ( uintptr_t, uintptr_t, int32_t, double *,int32_t *, int64_t *);
//
// extern void msk_cgo_streamfunc(uintptr_t, char*);
// extern int msk_cgo_callbackfunc(uintptr_t, int32_t, double*, int32_t*, int64_t*);
//
// extern int MSK_makeenv(MSKenv_t *, const string_t);
// extern int MSK_makeemptytask(MSKenv_t, MSKtask_t *);
// extern int MSK_deleteenv(MSKenv_t *);
// extern int MSK_deletetask(MSKtask_t *);
// extern int MSK_getlasterror64(MSKtask_t,int*,int64_t,int64_t*,char*);
// extern int MSK_linkfunctotaskstream(MSKtask_t task,int32_t whichstream, uintptr_t handle, MSKstreamfunc func);
// extern int32_t MSK_putcallbackfunc(MSKtask_t task,MSKcallbackfunc func,MSKuserhandle_t handle);
// extern int32_t MSK_clonetask(MSKtask_t,MSKtask_t *);
//
// static void stream_proxy(uintptr_t handle, const string_t msg) { msk_cgo_streamfunc(handle,(char *)msg); }
// static int32_t MSK_wrap_linkfunctotaskstream(MSKtask_t task, int32_t whichstream, uintptr_t handle) { return MSK_linkfunctotaskstream(task, whichstream, handle, stream_proxy); }
//
// static int callbackfunc_proxy(uintptr_t taskp, uintptr_t handlep, int32_t code, double * dinf, int32_t * iinf, int64_t * liinf) { return msk_cgo_callbackfunc(handlep,code,dinf,iinf,liinf); }
// static int32_t MSK_wrap_putcallbackfunc(MSKtask_t task, uintptr_t handle) { return MSK_putcallbackfunc(task, callbackfunc_proxy, (MSKuserhandle_t) handle); }
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

// Mosek response and error message
type MosekError struct {
    code Rescode
    msg string
}
func (self*MosekError) Error() string {
    return fmt.Sprintf("%d: %s",self.code,self.msg)
}

type ArrayLengthError struct {
    fun string
    arg string
}
func (self*ArrayLengthError) Error() string {
    return fmt.Sprintf("%s: Argument %s is too short",self.fun,self.arg)   
}


// Mosek Env. Multiple tasks can be create in each Env.
type Env struct {
        r    Rescode
        cptr C.MSKenv_t
}

// Mosek Task represents a single optimizatrion problem and solver information.
type Task struct {
    cptr            C.MSKtask_t
    streamfunc      [4]func(string)
    callbackfunc    func(Callbackcode)bool
    infcallbackfunc func(Callbackcode,[]float64,[]int32,[]int64)bool
}

func (t * Task) ptr() C.MSKtask_t { return t.cptr }
func (e * Env)  ptr() C.MSKenv_t  { return e.cptr }

func (self * Env) getlasterror(res C.int32_t) (Rescode,string) {
    return Rescode(res),""
}

func (self * Task) getlasterror(res C.int32_t) (Rescode,string) {
    var lastcode C.int32_t = res
    var lastmsglen C.long
    if 0 != C.MSK_getlasterror64(self.ptr(),&lastcode,0, &lastmsglen, nil) {
        return Rescode(lastcode),""
    } else {
        lastmsgbytes := make([]C.char,lastmsglen+1)
        if 0 != C.MSK_getlasterror64(self.ptr(),&lastcode,lastmsglen,&lastmsglen,(*C.char)(&lastmsgbytes[0])) {
            return Rescode(lastcode),""
        } else {
            return Rescode(lastcode),C.GoString(&lastmsgbytes[0])
        }
    }
}

//export msk_cgo_streamfunc
func msk_cgo_streamfunc(handle C.uintptr_t, msg * C.char) {
    if fun,ok := cgo.Handle(handle).Value().(func(string)); ok {
        fun(C.GoString(msg))
    }
}
//export msk_cgo_callbackfunc
func msk_cgo_callbackfunc(
	handle  C.uintptr_t,
	code    C.int32_t,
	dinf  * C.double,
	iinf  * C.int32_t,
	liinf * C.int64_t) (r C.int) {

        if task,ok := cgo.Handle(handle).Value().(*Task); ok {
            r = 0
            if task.infcallbackfunc != nil {
                _dinf  := (*[int(MSK_DINF_END)]float64)(unsafe.Pointer(dinf))[0:MSK_DINF_END]
                _iinf  := (*[int(MSK_IINF_END)]int32)  (unsafe.Pointer(iinf))[0:MSK_IINF_END]
                _liinf := (*[int(MSK_LIINF_END)]int64) (unsafe.Pointer(liinf))[0:MSK_LIINF_END]

                if task.infcallbackfunc(Callbackcode(code),_dinf,_iinf,_liinf) {
                    r = 1
                } else {
                    r = 0
                }
            }
            if task.callbackfunc != nil {
                if task.callbackfunc(Callbackcode(code)) {
                    r = 1
                } else {
                    r = 0
                }
            }
        }
	return
}

// Create a new Env
func MakeEnv() (env Env, err error) {
        if res := C.MSK_makeenv(&env.cptr,nil); res != 0 {
            err = &MosekError{code : Rescode(res) }
        }
        return
}

// Create a new Task in the Env
func (env *Env) MakeTask() (task *Task, err error) {
    var t Task
    if res := C.MSK_makeemptytask(env.ptr(), &t.cptr); res != 0 {
        err = &MosekError{ code:Rescode(res) }
    } else {
        task = &t 
    }
    return
}

// Create a new Task in the global Env
func NewTask() (task *Task, err error) {
    var t Task
    if res := C.MSK_makeemptytask(nil, &t.cptr); res != 0 {
        err = &MosekError{ code:Rescode(res) }
    } else { 
        task = &t
    }
    return
}

func (self *Task) Clone() (task *Task,err error) {
    var t Task
    if r := C.MSK_clonetask(self.ptr(),&t.cptr); r != 0 {
        lastcode,lastmsg := self.getlasterror(r)
        err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
    } else {
        task = &t 
    }
    return
}

// Delete Env. The Environment MUST NOT be used subsequently.
func (self *Env) Delete() (err error) {
    if r := C.MSK_deleteenv(&self.cptr); r != 0 {
        err = &MosekError{code : Rescode(r) }
    }
    return 
}

// Delete Task. The Task MUST NOT be used subsequently.
func (self *Task) Delete() (err error) {
    if r := C.MSK_deletetask(&self.cptr); r != 0 {
        err = &MosekError{code : Rescode(r) }
    }
    return
}


// Attach a stream callback function for receiving log messages.
func (self *Task) PutStreamFunc(whichstream Streamtype, fun func(string)) (err error) {
	if fun == nil {
            if self.streamfunc[whichstream] != nil {
                self.streamfunc[whichstream] = nil
                if r := C.MSK_linkfunctotaskstream(self.ptr(), C.int32_t(whichstream), 0, nil); r != 0 {
                    lastcode,lastmsg := self.getlasterror(r)
                    err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
                }
            }
	} else {
            self.streamfunc[whichstream] = fun
            handle := C.uintptr_t(cgo.NewHandle(fun))
            if r := C.MSK_wrap_linkfunctotaskstream(self.ptr(),C.int32_t(whichstream),handle); r != 0 {
                lastcode,lastmsg := self.getlasterror(r)
                err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
            }
	}
        return
}


func (self *Task) updateCallbackFunc() (err error) {
    if self.callbackfunc == nil && self.infcallbackfunc == nil {
        if r := C.MSK_putcallbackfunc(self.ptr(), nil, nil); r != 0 {
            lastcode,lastmsg := self.getlasterror(r)
            err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
            return
        }
    } else {
        handle := C.uintptr_t(cgo.NewHandle(self))
        if r := C.MSK_wrap_putcallbackfunc(self.ptr(), handle); r != 0 {
            lastcode,lastmsg := self.getlasterror(r)
            err = &MosekError{code:Rescode(lastcode),msg:lastmsg}
            return
        }
    }
    return
}


// Attach a callback function for receiving progress status during optimizations.
func (self *Task) PutCallbackFunc(fun func(Callbackcode) bool) error {
	self.callbackfunc = fun
        return self.updateCallbackFunc()
}

// Attach a callback function for receiving information during optimizations.
func (self *Task) PutInfoCallbackFunc(fun func(Callbackcode,[]float64,[]int32,[]int64) bool) (err error)  {
	self.infcallbackfunc = fun
        return self.updateCallbackFunc()
}

func minint(a []int) (r int) {
        if len(a) == 0 { panic("Minimum of empty array") }
        r = a[0]
        for i := 1; i < len(a); i++ {
                if a[i] < r { r = a[i] }
        }
        return
}

func sum[T int32|int64|float64](data []T) T {
    var r T
    for _,v := range data { r += v }
    return r
}


