/*
   Copyright: Copyright (c) MOSEK ApS, Denmark. All rights reserved.

   File:      concurrent1.java

   Purpose: Demonstrates a simple implementation of a concurrent optimizer.

            The concurrent optimizer starts a few parallel optimizations
            of the same problem using different algorithms, and reports
            a solution when the first optimizer is ready.

            This example also demonstrates how to define a simple callback handler
            that stops the optimizer when requested.
*/
package main

import (
	"fmt"
	"os"
        "errors"
	"strconv"

	mosek "github.com/mosek/mosek.go"
)

/* Takes a list of tasks and optimizes then in parallel. Error and/or termination code from each optimization is
   returned in ``err`` and ``trm``.

   When one task completes with rescode == ok, others are terminated.

   Returns the index of the first task that returned
   with rescode == ok. Whether or not this task contains the
   most valuable answer, is for the caller to decide. If none
   completed without error return -1.
*/
func optimize(tasks []*mosek.Task) (first int, errs []error, trms []mosek.Rescode) {
    first = -1
    n := len(tasks)
    stop := false
    
    errs = make([]error,n)
    trms = make([]mosek.Rescode,n)
    stopc := make([]chan int,n)
    for i := 0; i < n; i++ { 
        trms[i] = mosek.MSK_RES_ERR_UNKNOWN 
        stopc[i] = make(chan int)
    }
        
    // Start parallel optimizations, one per task
    for ix,tx := range tasks {
        t := tx
        i := ix
        cn := make(chan int)
        stopc[i] = cn
        go func() {
            t.PutCallbackFunc(func(code mosek.Callbackcode) bool { return stop })
            defer close(cn)
            trm,err := t.Optimize()
            errs[i] = err
            if err == nil { 
                trms[i] = trm
                if trm == mosek.MSK_RES_OK {
                    stop = true
                    first = i
                }
            }
        }()
    }

    for _,c := range stopc { <- c }

    // For debugging, print res and trm codes for all optimizers
    for i,trm := range trms {
        if errs[i] != nil {
            fmt.Printf("Optimizer  %d : Error %s\n",i,errs[i]);
        } else {
            fmt.Printf("Optimizer  %d : Termination code = %s\n",i,trm);
        }
    }

    return
}


/* 
    Given a continuous task, set up jobs to optimize it 
    with a list of different solvers.

    Returns an index, corresponding to the optimization
    task that is returned as winTask. This is the task
    with the best possible status of those that finished.
    If none task is considered successful returns -1.
 */
func OptimizeConcurrent(task *mosek.Task, optimizers []mosek.Optimizertype) (winIndex int, winTask *mosek.Task, trm mosek.Rescode, err error) { 
    n := len(optimizers)
    tasks := make([]*mosek.Task,n);

    // Clone tasks and choose various optimizers
    for i := 0; i < n; i++ {
        if t,err := task.Clone(); err != nil {
            panic(err)
        } else if err := t.PutIntParam(mosek.MSK_IPAR_OPTIMIZER, int32(optimizers[i])); err != nil {
            panic(err)
        } else {
            tasks[i] = t 
        }
    }

    // Solve tasks in parallel
    first,_,trms := optimize(tasks)

    if (first >= 0) {
        winIndex = first
        winTask = tasks[first]
        trm = trms[first]
    } else {
        err = errors.New("No task successfully terminated")
    }
    return 
}

/*  Given a mixed-integer task, set up jobs to optimize it 
    with different values of seed. That will lead to
    different execution paths of the optimizer.

    Returns an index, corresponding to the optimization
    task that is returned as winTask. This is the task
    with the best value of the objective function.
    If none task is considered successful returns -1.

    Typically, the input task would contain a time limit. The two
    major scenarios are:
    1. Some clone ends before time limit - then it has optimum.
    2. All clones reach time limit - pick the one with best objective.
 */
func OptimizeConcurrentMIO(task *mosek.Task, seeds []int32) (winIndex int, winTask *mosek.Task, trm mosek.Rescode) {
    n := len(seeds)

    tasks := make([]*mosek.Task,n)

    // Clone tasks and choose various seeds for the optimizer
    for i,seed := range(seeds) {
        if t,err := task.Clone(); err != nil {
            panic(err)
        } else if t.PutIntParam(mosek.MSK_IPAR_MIO_SEED, seed); err != nil {
            panic(err)
        } else {
            tasks[i] = t
        }
    }

    // Solve tasks in parallel
    firstOK,errs,trms :=  optimize(tasks)

    bestObj := 1e10;
    if sense,err := task.GetObjSense(); err != nil { panic(err) 
    } else if sense == mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE {
        bestObj = -1e10
    }
    if firstOK >= 0 {
        // Pick the task that ended with res = ok
        // and contains an integer solution with best objective value

        for i,t := range(tasks) {
            if pobj,err := t.GetPrimalObj(mosek.MSK_SOL_ITG); err != nil { 
                // ignore
              // panic(err)
            } else {
                fmt.Printf("%d    %f\n",i,pobj)
            }
        }

    
        for i,t := range(tasks) {
            if errs[i] != nil {
            } else if sense,err := task.GetObjSense(); err != nil { 
                panic(err) 
            } else if solstai,err := task.GetSolSta(mosek.MSK_SOL_ITG); err != nil {
                panic(err)
            } else if pobji,err := t.GetPrimalObj(mosek.MSK_SOL_ITG); err != nil {
                // ignore 
            } else if solstai == mosek.MSK_SOL_STA_PRIM_FEAS || solstai == mosek.MSK_SOL_STA_INTEGER_OPTIMAL {
                if ((sense == mosek.MSK_OBJECTIVE_SENSE_MINIMIZE && pobji < bestObj) || 
                    (sense == mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE && pobji >= bestObj)) {
                    bestObj = pobji
                    winIndex = i
                }
            }

            if winIndex != -1 {
                winTask = tasks[winIndex]
                trm   = trms[winIndex] 
                return 
            }
        }
    }
  
    return
}

  /** 
     This is an example of how one can use the methods
         OptimizeConcurrent
         OptimizeConcurrentMIO

     argv[0] : name of file with input problem
     argv[1]: (optional) time limit
   */
func main() {
    task,err := mosek.NewTask(); if err != nil { panic(err) }

    if len(os.Args) > 1 {
        if err := task.ReadData(os.Args[1]); err != nil {
            panic(err)
        }
    } else {
        panic("Missng argument")
    }


    // Optional time limit
    if len(os.Args) > 2 {
        timeLimit,err := strconv.ParseFloat(os.Args[2],64); if err != nil { panic(err) }
        if err := task.PutDouParam(mosek.MSK_DPAR_OPTIMIZER_MAX_TIME, timeLimit); err != nil { panic(err) }
    }

    var trm  mosek.Rescode
    var t   *mosek.Task
    var idx  int

    if numint,err := task.GetNumIntVar(); err != nil {
        panic(err)
    } else if (numint == 0) {
        // If the problem is continuous
        // optimize it with three continuous optimizers.
        //(Simplex will fail for non-linear problems)

        optimizers := []mosek.Optimizertype{ 
            mosek.MSK_OPTIMIZER_CONIC,
            mosek.MSK_OPTIMIZER_DUAL_SIMPLEX,
            mosek.MSK_OPTIMIZER_PRIMAL_SIMPLEX,
        }

        idx,t,trm,err = OptimizeConcurrent(task, optimizers)
    } else {
        // Mixed-integer problem.
        // Try various seeds.
        seeds := []int32{ 42, 13, 71749373 }

        idx,t,trm = OptimizeConcurrentMIO(task, seeds)
    }          

    // Check results and print the best answer
    if idx < 0 {
        fmt.Println("All optimizers failed.")
    } else {
        fmt.Printf("Result from optimizer with index %d: trm: %s\n",idx,trm)
        if err := t.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) }); err != nil {
            panic(err)
        } else if err := t.OptimizerSummary(mosek.MSK_STREAM_LOG); err != nil {
            panic(err)
        } else if err := t.SolutionSummary(mosek.MSK_STREAM_LOG); err != nil {
            panic(err)
        }
    }
}

