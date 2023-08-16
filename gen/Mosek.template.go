package mosek

// /*<comment>*/

// #include <stdlib.h>
// #include <stdint.h>
// #cgo LDFLAGS: -lmosek64
//
// typedef void * MSKuserhandle_t;
// typedef const char * string_t;
// typedef void * MSKtask_t;
// typedef void * MSKenv_t;
// typedef void  (* MSKstreamfunc) ( void *, const char *);
// typedef int32_t  (* MSKcallbackfunc) ( MSKtask_t, MSKuserhandle_t, int32_t, const double *, const int32_t *, const int64_t *);
// 
// extern void streamfunc_log(void *, char *);
// extern void streamfunc_wrn(void *, char *);
// extern void streamfunc_msg(void *, char *);
// extern void streamfunc_err(void *, char *);
// extern int callbackfunc(void *, void *, int, double*,int32_t*, int64_t*);
// extern int MSK_makeenv(MSKenv_t *, const char *);
// extern int MSK_makeemptytask(MSKenv_t, MSKtask_t *);
// extern int MSK_deleteenv(MSKenv_t *);
// extern int MSK_deletetask(MSKtask_t *);
// extern int MSK_getlasterror64(MSKtask_t,int*,int64_t,int64_t*,char*);
// extern int MSK_linkfunctotaskstream(MSKtask_t task,int32_t whichstream, MSKuserhandle_t handle, MSKstreamfunc func);
// int32_t MSK_putcallbackfunc(MSKtask_t task,MSKcallbackfunc func,MSKuserhandle_t handle);
//<extern>
import "C"

import (
    "unsafe"
    "fmt"
    "golang.org/x/exp/constraints"
)
//<consts>


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


type Env struct {
        r    Rescode
        cptr C.MSKenv_t
}

type Task struct {
        r               Rescode
        cptr            C.MSKtask_t
	streamfunc      [4]func(string)
	callbackfunc    func(int32)int
	infcallbackfunc func(int32,[]float64,[]int32,[]int64)int
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


//export streamfunc_log
func streamfunc_log(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_LOG] != nil { task.streamfunc[MSK_STREAM_LOG](C.GoString(msg)) }
}

//export streamfunc_msg
func streamfunc_msg(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_MSG] != nil { task.streamfunc[MSK_STREAM_MSG](C.GoString(msg)) }
}

//export streamfunc_wrn
func streamfunc_wrn(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_WRN] != nil { task.streamfunc[MSK_STREAM_WRN](C.GoString(msg)) }
}

//export streamfunc_err
func streamfunc_err(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[MSK_STREAM_ERR] != nil { task.streamfunc[MSK_STREAM_ERR](C.GoString(msg)) }
}

//export callbackfunc
func callbackfunc(
	nativetask unsafe.Pointer,
	handle  unsafe.Pointer,
	code    C.int,
	dinf  * C.double,
	iinf  * C.int32_t,
	liinf * C.int64_t) (C.int) {

	task := (*Task)(handle)

	var r int = 0

	if task.infcallbackfunc != nil {
		_dinf  := (*[int(MSK_DINF_END)]float64)(unsafe.Pointer(dinf))[0:MSK_DINF_END]
		_iinf  := (*[int(MSK_IINF_END)]int32)  (unsafe.Pointer(iinf))[0:MSK_IINF_END]
		_liinf := (*[int(MSK_LIINF_END)]int64) (unsafe.Pointer(liinf))[0:MSK_LIINF_END]

		r = task.infcallbackfunc(int32(code),_dinf,_iinf,_liinf)
	} else if task.callbackfunc != nil {
		r = task.callbackfunc(int32(code))
	}
	return C.int(r)
}


func MakeEnv() (env Env, err error) {
        if res := C.MSK_makeenv(&env.cptr,nil); res != 0 {
            err = &MosekError{code : Rescode(res) }
        }
        return
}

func (env *Env) MakeTask() (task Task, err error) {
    if res := C.MSK_makeemptytask(env.ptr(), &task.cptr); res != 0 {
        err = &MosekError{ code:Rescode(res) }
    } else { 
	task.streamfunc[0]   = nil
	task.streamfunc[1]   = nil
	task.streamfunc[2]   = nil
	task.streamfunc[3]   = nil
	task.callbackfunc    = nil
	task.infcallbackfunc = nil
    }
    return
}

func NewTask() (task Task, err error) {
    if res := C.MSK_makeemptytask(nil, &task.cptr); res != 0 {
        err = &MosekError{ code:Rescode(res) }
    } else { 
	task.streamfunc[0]   = nil
	task.streamfunc[1]   = nil
	task.streamfunc[2]   = nil
	task.streamfunc[3]   = nil
	task.callbackfunc    = nil
	task.infcallbackfunc = nil
    }
    return
}

func (self *Env) Delete() (err error) {
    if r := C.MSK_deleteenv(&self.cptr); r != 0 {
        err = &MosekError{code : Rescode(r) }
    }
    return 
}

func (self *Task) Delete() (err error) {
    if r := C.MSK_deletetask(&self.cptr); r != 0 {
        err = &MosekError{code : Rescode(r) }
    }
    return
}

func (self *Task) PutStreamFunc(whichstream Streamtype, fun func(string)) {
	self.streamfunc[whichstream] = fun

	if fun == nil {
		C.MSK_linkfunctotaskstream(
			self.ptr(),
			C.int32_t(whichstream),
			nil,
			nil)
	} else {
		var strmfun (*[0]byte)
		switch whichstream {
		case MSK_STREAM_MSG: strmfun = (*[0]byte)(C.streamfunc_msg)
		case MSK_STREAM_LOG: strmfun = (*[0]byte)(C.streamfunc_log)
		case MSK_STREAM_ERR: strmfun = (*[0]byte)(C.streamfunc_err)
		case MSK_STREAM_WRN: strmfun = (*[0]byte)(C.streamfunc_wrn)
		}

		C.MSK_linkfunctotaskstream(
			self.ptr(),
			C.int32_t(whichstream),
			C.MSKuserhandle_t(unsafe.Pointer(self)),
			strmfun) // ?!?
	}
}

func (self *Task) PutCallbackFunc(fun func(int32) int) {
	self.callbackfunc = fun
	if fun == nil {
		C.MSK_putcallbackfunc(self.ptr(), nil, nil)
	} else {
		C.MSK_putcallbackfunc(self.ptr(), (*[0]byte)(C.callbackfunc), C.MSKuserhandle_t(unsafe.Pointer(self)))
	}
}

func (self *Task) PutInfoCallbackFunc(fun func(int32,[]float64,[]int32,[]int64) int) {
	self.infcallbackfunc = fun
	if fun == nil {
		C.MSK_putcallbackfunc(self.ptr(), nil, nil)
	} else {
		C.MSK_putcallbackfunc(self.ptr(), (*[0]byte)(C.callbackfunc), C.MSKuserhandle_t(unsafe.Pointer(self)))
	}
}

func (e * Env)  ClearError() { e.r = MSK_RES_OK }
func (self * Task) ClearError() { self.r = MSK_RES_OK }

//func (e * Env)  GetRes() int32 { return e.r }
//func (t * Task) GetRes() int32 { return t.r }

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


//<methods>
