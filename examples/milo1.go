/*
   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.

   File :      milo1.go

   Purpose :   Demonstrates how to solve a small mixed
               integer linear optimization problem using the MOSEK Java API.
*/
package main

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
)

func main() {
    const numcon int32 = 2
    const numvar int32 = 2

    // Since the value infinity is never used, we define
    // 'infinity' symbolic purposes only
    const infinity = 0.0

    bkc := []mosek.Boundkey{ mosek.MSK_BK_UP, mosek.MSK_BK_LO }
    blc := []float64{ -infinity,         -4.0 }
    buc := []float64{ 250.0,             infinity }

    bkx := []mosek.Boundkey{ mosek.MSK_BK_LO, mosek.MSK_BK_LO }
    blx := []float64{ 0.0,               0.0 }
    bux := []float64{ infinity,          infinity }

    c   := []float64{1.0, 0.64 }

    asub := []int32{  0,   1,     0,    1 }
    aval := []float64{ 50.0, 3.0,  31.0, -2.0 }

    ptrb := []int64{ 0, 2 }
    ptre := []int64{ 2, 4 }

    task,err := mosek.NewTask()
    if err != nil { panic(err) }
    defer task.Delete()

    // Directs the log task stream to the 'printstr' function. 
    task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })

    task.PutItgSolutionCallbackFunc(
        func(xx []float64)bool {
          fmt.Printf("New integer solution: %v\n",xx)
          return false
        })
    /* Append 'numcon' empty constraints.
    The constraints will initially have no bounds. */
    if err := task.AppendCons(numcon); err != nil { panic(err) }

    /* Append 'numvar' variables.
    The variables will initially be fixed at zero (x=0). */
    if err := task.AppendVars(numvar); err != nil { panic(err) }

    for j,cj := range c {
        /* Set the linear term c_j in the objective.*/
        if err := task.PutCJ(int32(j), cj); err != nil { panic(err) }
        /* Set the bounds on variable j.
           blx[j] <= x_j <= bux[j] */
        if err := task.PutVarBound(int32(j), bkx[j], blx[j], bux[j]); err != nil { panic(err) }
        /* Input column j of A */
        if err := task.PutACol(int32(j),    /* Variable (column) index.*/
                               asub[ptrb[j]:ptre[j]],     /* Row index of non-zeros in column j.*/
                               aval[ptrb[j]:ptre[j]]); err != nil { panic(err) } /* Non-zero Values of column j. */
    }
    /* Set the bounds on constraints.
     for i=1, ...,numcon : blc[i] <= constraint i <= buc[i] */
    for i,bki := range bkc { 
        if err := task.PutConBound(int32(i), bki, blc[i], buc[i]); err != nil { panic(err) }
    }
    /* Specify integer variables. */
    for j,_ := range(bkx) {
        if err := task.PutVarType(int32(j), mosek.MSK_VAR_TYPE_INT); err != nil { panic(err) }
    }

    /* Set max solution time */
    if err := task.PutDouParam(mosek.MSK_DPAR_MIO_MAX_TIME, 60.0); err != nil { panic(err) }


    /* A maximization problem */
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE); err != nil { panic(err) }
    /* Solve the problem */
      
    if trm,err := task.Optimize(); err != nil { 
        panic(err) 
    } else if trm != mosek.MSK_RES_OK {
        fmt.Printf(" Mosek warning: %s\n",trm)
    }

    // Print a summary containing information
    //   about the solution for debugging purposes
    if err := task.SolutionSummary(mosek.MSK_STREAM_MSG); err != nil { panic(err) }
    xx,err := task.GetXx(mosek.MSK_SOL_ITG); if err != nil { panic(err) }
    /* Get status information about the solution */
    solsta,err := task.GetSolSta(mosek.MSK_SOL_ITG); if err != nil { panic(err) }

    switch solsta {
    case mosek.MSK_SOL_STA_INTEGER_OPTIMAL:
        fmt.Printf("Optimal solution: %v\n",xx)
    case mosek.MSK_SOL_STA_PRIM_FEAS:
        fmt.Printf("Feasible solution: %v\n",xx)
    case mosek.MSK_SOL_STA_UNKNOWN:
        prosta,err := task.GetProSta(mosek.MSK_SOL_ITG); if err != nil { panic(err) }
        switch prosta {
        case mosek.MSK_PRO_STA_PRIM_INFEAS_OR_UNBOUNDED:
            fmt.Println("Problem status Infeasible or unbounded")
        case mosek.MSK_PRO_STA_PRIM_INFEAS:
            fmt.Println("Problem status Infeasible.")
        case mosek.MSK_PRO_STA_UNKNOWN:
            fmt.Println("Problem status unknown.")
        default:
            fmt.Printf("Other problem status: %s\n",prosta)
        }
    default:
      fmt.Println("Other solution status")
    }
}
