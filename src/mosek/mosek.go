package mosek

// #include "mosek.h"
// #cgo CFLAGS: -I/home/ulfw/mosek/7/tools/platform/linux64x86/h
// #cgo LDFLAGS: -L/home/ulfw/mosek/7/tools/platform/linux64x86/bin -lmosek64
import "C" 

import (
        "unsafe"
)

type Env struct {
        env unsafe.Pointer
}

type Task struct {
        task unsafe.Pointer
}


func MakeEnv() (env Env, res int32) {
        var envptr C.MSKenv_t
        res = int32(C.MSK_makeenv(&envptr,nil))
        if res == 0 {
                env.env = unsafe.Pointer(envptr)
        }        
        return
}

func (env *Env) MakeTask() (task Task, res int32) {
        var taskptr C.MSKtask_t
        res = int32(C.MSK_maketask(C.MSKenv_t(env.env), 0,0, &taskptr))
        if res == 0 {
                task.task = unsafe.Pointer(taskptr)
        }
        return
}

func (e *Env) DeleteEnv() {
        envptr := C.MSKenv_t(e.env)
        C.MSK_deleteenv(&envptr)
        e.env = nil
}

func (t *Task) DeleteTask() {
        taskptr := C.MSKtask_t(t.task)
        C.MSK_deletetask(&taskptr)
        t.task = nil
}

func (t * Task) ReadData(filename string) (res int32) {
        c_filename := C.CString(filename)
        defer C.free(unsafe.Pointer(c_filename))
        res = int32(C.MSK_readdataautoformat(C.MSKtask_t(t.task),c_filename))
        return
}

func (t * Task) WriteData(filename string) (res int32) {
        c_filename := C.CString(filename)
        defer C.free(unsafe.Pointer(c_filename))
        res = int32(C.MSK_writedata(C.MSKtask_t(t.task),c_filename))
        return
}

func (t * Task) GetNumCon() (numcon int, res int32) {
        var c_numcon C.MSKint32t
        res = int32(C.MSK_getnumcon(C.MSKtask_t(t.task), &c_numcon))
        numcon = int(c_numcon)
        return         
}

func (t * Task) GetNumVar() (numvar int, res int32) {
        var c_numvar C.MSKint32t
        res = int32(C.MSK_getnumvar(C.MSKtask_t(t.task), &c_numvar))
        numvar = int(c_numvar)
        return         
}

func (t * Task) PutARow(i int, subj []int32, valj []float64) (res int32) {
        nzi := len(subj)
        if len(valj) < nzi { nzi = len(valj) }
                
        res = int32(C.MSK_putarow(
                C.MSKtask_t(t.task), 
                C.MSKint32t(i), 
                C.MSKint32t(nzi), 
                (*C.MSKint32t)(&subj[0]), 
                (*C.MSKrealt)(&valj[0])))
        return
}

func (t * Task) GetARowNumNZ(i int) (nzi int, res int32) {
        var c_nzi C.MSKint32t
        res = int32(C.MSK_getarownumnz(C.MSKtask_t(t.task),C.MSKint32t(i), &c_nzi))        
        nzi = int(c_nzi)
        return 
}

func (t * Task) GetARow(i int, subj []int32, valj []float64) (rsubj []int32,rvalj []float64, res int32) {
        var c_nzi C.MSKint32t
        res = int32(C.MSK_getarownumnz(C.MSKtask_t(t.task),C.MSKint32t(i), &c_nzi))        
        if res != 0 { return }
        
        var nzi = int(c_nzi)
        
        if subj == nil { 
                rsubj = make([]int32,nzi,nzi)
        } else if len(subj) < nzi { 
                panic("Argument subj is too short in call to GetARow")
        } else {
                rsubj = subj
        }

        if valj == nil { 
                rvalj = make([]float64,nzi,nzi) 
        } else if len(valj) < nzi { 
                panic("Argument valj is too short in call to GetARow")
        } else {
                rvalj = valj
        }

        res = int32(C.MSK_getarow(C.MSKtask_t(t.task),C.MSKint32t(i), nil, (*C.MSKint32t)(&rsubj[0]),(*C.MSKrealt)(&rvalj[0])))
        
        return
}

