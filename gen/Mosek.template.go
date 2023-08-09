package mosek

// #include <stdlib.h>
// #include <mosek.h>
// #cgo LDFLAGS: -lmosek64
//
// extern void streamfunc_log(void *, char *);
// extern void streamfunc_wrn(void *, char *);
// extern void streamfunc_msg(void *, char *);
// extern void streamfunc_err(void *, char *);
// extern int callbackfunc(void *, void *, int, double *, int *, long long *);
import (
    "C"
    "unsafe"
)

//{consts}


type Env struct {
        r    int32
        cptr unsafe.Pointer
}

type Task struct {
        r               int32
        cptr            unsafe.Pointer
	streamfunc      [4]func(string)
	callbackfunc    func(int32)int
	infcallbackfunc func(int32,[]float64,[]int32,[]int64)int
}

func (t * Task) ptr() C.MSKtask_t { return C.MSKtask_t(t.cptr) }
func (e * Env)  ptr() C.MSKenv_t  { return C.MSKenv_t(e.cptr) }

//export streamfunc_log
func streamfunc_log(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[STREAM_LOG] != nil { task.streamfunc[STREAM_LOG](C.GoString(msg)) }
}

//export streamfunc_msg
func streamfunc_msg(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[STREAM_MSG] != nil { task.streamfunc[STREAM_MSG](C.GoString(msg)) }
}

//export streamfunc_wrn
func streamfunc_wrn(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[STREAM_WRN] != nil { task.streamfunc[STREAM_WRN](C.GoString(msg)) }
}

//export streamfunc_err
func streamfunc_err(handle unsafe.Pointer, msg * C.char) {
	task := (*Task)(handle)
	if task.streamfunc[STREAM_ERR] != nil { task.streamfunc[STREAM_ERR](C.GoString(msg)) }
}

//export callbackfunc
func callbackfunc(
	nativetask unsafe.Pointer,
	handle  unsafe.Pointer,
	code    C.int,
	dinf  * C.MSKrealt,
	iinf  * C.MSKint32t,
	liinf * C.MSKint64t) (C.int) {

	task := (*Task)(handle)

	var r int = 0

	if task.infcallbackfunc != nil {
		_dinf  := (*[int(DINF_END)]float64)(unsafe.Pointer(dinf))[0:DINF_END]
		_iinf  := (*[int(IINF_END)]int32)  (unsafe.Pointer(iinf))[0:IINF_END]
		_liinf := (*[int(LIINF_END)]int64) (unsafe.Pointer(liinf))[0:LIINF_END]

		r = task.infcallbackfunc(int32(code),_dinf,_iinf,_liinf)
	} else if task.callbackfunc != nil {
		r = task.callbackfunc(int32(code))
	}
	return C.int(r)
}


func MakeEnv() (env Env, res int32) {
        var envptr C.MSKenv_t
        res = int32(C.MSK_makeenv(&envptr,nil))
        if res == 0 {
                env.cptr = unsafe.Pointer(envptr)
        }
        return
}

func (env *Env) MakeTask() (task Task, res int32) {
        var taskptr C.MSKtask_t
        res = int32(C.MSK_makeemptytask(env.ptr(), &taskptr))
        if res != 0 { return }
	task.cptr            = unsafe.Pointer(taskptr)
	task.streamfunc[0]   = nil
	task.streamfunc[1]   = nil
	task.streamfunc[2]   = nil
	task.streamfunc[3]   = nil
	task.callbackfunc    = nil
	task.infcallbackfunc = nil

        return
}

func (e *Env) Delete() {
        envptr := e.ptr()
        C.MSK_deleteenv(&envptr)
        e.cptr = nil
}

func (t *Task) Delete() {
        taskptr := t.ptr()
        C.MSK_deletetask(&taskptr)
        t.cptr = nil
}

func (t *Task) PutStreamFunc(whichstream int32, fun func(string)) {
	t.streamfunc[whichstream] = fun

	if fun == nil {
		C.MSK_linkfunctotaskstream(
			t.ptr(),
			C.MSKstreamtypee(whichstream),
			nil,
			nil)
	} else {
		var strmfun (*[0]byte)
		switch whichstream {
		case STREAM_MSG: strmfun = (*[0]byte)(C.streamfunc_msg)
		case STREAM_LOG: strmfun = (*[0]byte)(C.streamfunc_log)
		case STREAM_ERR: strmfun = (*[0]byte)(C.streamfunc_err)
		case STREAM_WRN: strmfun = (*[0]byte)(C.streamfunc_wrn)
		}

		C.MSK_linkfunctotaskstream(
			t.ptr(),
			C.MSKstreamtypee(whichstream),
			C.MSKuserhandle_t(unsafe.Pointer(t)),
			strmfun) // ?!?
	}
}

func (t *Task) PutCallbackFunc(fun func(int32) int) {
	t.callbackfunc = fun
	if fun == nil {
		C.MSK_putcallbackfunc(t.ptr(), nil, nil)
	} else {
		C.MSK_putcallbackfunc(t.ptr(), (*[0]byte)(C.callbackfunc), C.MSKuserhandle_t(unsafe.Pointer(t)))
	}
}

func (t *Task) PutInfoCallbackFunc(fun func(int32,[]float64,[]int32,[]int64) int) {
	t.infcallbackfunc = fun
	if fun == nil {
		C.MSK_putcallbackfunc(t.ptr(), nil, nil)
	} else {
		C.MSK_putcallbackfunc(t.ptr(), (*[0]byte)(C.callbackfunc), C.MSKuserhandle_t(unsafe.Pointer(t)))
	}
}

func (e * Env)  ClearError() { e.r = RES_OK }
func (t * Task) ClearError() { t.r = RES_OK }

func (e * Env)  GetRes() int32 { return e.r }
func (t * Task) GetRes() int32 { return t.r }

func minint(a []int) (r int) {
        if len(a) == 0 { panic("Minimum of empty array") }
        r = a[0]
        for i := 1; i < len(a); i++ {
                if a[i] < r { r = a[i] }
        }
        return
}

