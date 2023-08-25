//
//   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//   File :      solvelinear.java
//
//   Purpose :   To demonstrate the usage of MSK_solvewithbasis
//               when solving the linear system:
//
//               1.0  x1             = b1
//               -1.0  x0  +  1.0  x1 = b2
//
//               with two different right hand sides
//
//               b = (1.0, -2.0)
//
//               and
//
//               b = (7.0, 0.0)
//
package main

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
)

func setup(task *mosek.Task,
           aval   []float64,
           asub   []int32, 
           ptr    []int,
           numvar int32) []int32 {
    // Since the value infinity is never used, we define
    // 'infinity' symbolic purposes only
    const infinity = 0.0

    skx := make([]mosek.Stakey,numvar)
    skc := make([]mosek.Stakey,numvar)

    for i := 0; i < int(numvar); i++ {
        skx[i] = mosek.MSK_SK_BAS
        skc[i] = mosek.MSK_SK_FIX
    }

    if err := task.AppendVars(numvar); err != nil { panic(err) }
    if err := task.AppendCons(numvar); err != nil { panic(err) }

    for i := 0; i < int(numvar); i++ {
        if err := task.PutACol(int32(i), asub[ptr[i]:ptr[i+1]], aval[ptr[i]:ptr[i+1]]); err != nil { panic(err) }
    }

    for i := 0 ; i < int(numvar); i++ {
        if err := task.PutConBound( int32(i), mosek.MSK_BK_FX, 0.0, 0.0); err != nil { panic(err) }
    }
    for i := int32(0) ; i < numvar; i++ {
        if err := task.PutVarBound(i, mosek.MSK_BK_FR, -infinity, infinity); err != nil { panic(err) }
    }

    // Define a basic solution by specifying
    // status keys for variables & constraints.
    if err := task.DeleteSolution(mosek.MSK_SOL_BAS); err != nil { panic(err) }

    if err := task.PutSkcSlice(mosek.MSK_SOL_BAS, 0, numvar, skc); err != nil { panic(err) }
    if err := task.PutSkxSlice(mosek.MSK_SOL_BAS, 0, numvar, skx); err != nil { panic(err) }

    if basis,err := task.InitBasisSolve(); err != nil { 
        panic(err) 
    } else {
        return basis
    }
}

func main() {
    const numcon = 2
    const numvar = 2

    aval := []float64{ -1.0,
                        1.0, 1.0 }
    asub := []int32{ 1,
                     0, 1 }

    ptr  := []int{0, 1, 3}

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    defer task.Delete()
    // Directs the log task stream to the user specified
    // method task_msg_obj.streamCB
    task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })

    // Put A matrix and factor A.
    // Call this function only once for a given task.

    basis := setup( task, aval, asub, ptr, numvar)

    /* now solve rhs */
    b := make([]float64,numvar)
    bsub := make([]int32,numvar)
    b[0] = 1
    b[1] = -2
    bsub[0] = 0
    bsub[1] = 1
    if nz,err := task.SolveWithBasis(false, 2, bsub, b); err != nil { 
        panic(err) 
    } else {
        fmt.Println("\nSolution to Bx = b:\n")

        // Print solution and show correspondents
        // to original variables in the problem 
        for i := int32(0); i < nz; i++ {
            if (basis[bsub[i]] < numcon) {
                fmt.Println("This should never happen")
            } else {
                fmt.Printf("x%d = %f\n",basis[bsub[i]] - numcon,b[bsub[i]])
            }
        }
    }
    b[0] = 7
    bsub[0] = 0
    if nz,err := task.SolveWithBasis(false, 1, bsub, b); err != nil { 
        panic(err) 
    } else {
        fmt.Println ("\nSolution to Bx = b:\n")
        // Print solution and show correspondents
        // to original variables in the problem
        for i := int32(0); i < nz; i++ {
            if basis[bsub[i]] < numcon {
                fmt.Println("This should never happen")
            } else {
                fmt.Printf("x%d = %f\n",basis[bsub[i]] - numcon,b[bsub[i]] )
            }
        }
    }
}
