//
//   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//   File :      ceo1.go
//
//   Purpose :   Demonstrates how to solve a small conic exponential
//               optimization problem using the MOSEK API.
//
package main 

import (
    mosek "github.com/mosek/mosek.go"
    "fmt"
)

func main() {
    const numcon = 1
    const numvar = 3
    const infinity = 0

    // Since the value infinity is never used, we define
    // 'infinity' symbolic purposes only

    bkc := mosek.MSK_BK_FX
    blc :=  1.0
    buc :=  1.0

    bkx := []mosek.Boundkey{ mosek.MSK_BK_FR,
                             mosek.MSK_BK_FR,
                             mosek.MSK_BK_FR }
    blx := []float64{ -infinity,
                      -infinity,
                      -infinity }
    bux := []float64{ +infinity,
                      +infinity,
                      +infinity }

    c   := []float64{ 1.0, 1.0, 0.0 }

    a      := []float64{ 1.0, 1.0, 1.0 }
    asub   := []int32{0, 1, 2}

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    // Directs the log task stream to the user specified
    // method task_msg_obj.stream
    if err := task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) }); err != nil { panic(err) }

    // Append 'numcon' empty constraints.
    // The constraints will initially have no bounds. 
    if err := task.AppendCons(numcon); err != nil { panic(err) }

    // Append 'numvar' variables.
    // The variables will initially be fixed at zero (x=0).
    if err := task.AppendVars(numvar); err != nil { panic(err) }

    /* Define the linear part of the problem */
    if err := task.PutCSlice(0, numvar, c); err != nil { panic(err) }
    if err := task.PutARow(0, asub, a); err != nil { panic(err) }
    if err := task.PutConBound(0, bkc, blc, buc); err != nil { panic(err) }
    if err := task.PutVarBoundSlice(0, numvar, bkx, blx, bux); err != nil { panic(err) }

    // Add a conic constraint
    // Create a 3x3 identity matrix F
    if err := task.AppendAfes(3); err != nil { panic(err) }
    if err := task.PutAfeFEntryList([]int64{0, 1, 2},          /* Rows */
                                    []int32{0, 1, 2},          /* Columns */
                                    []float64{1.0, 1.0, 1.0}); err != nil { panic(err) }


    // Exponential cone (x(0),x(1),x(2)) \in EXP 
    expdomain,err := task.AppendPrimalExpConeDomain(); if err != nil { panic(err) }
    if err := task.AppendAcc(expdomain,            /* Domain */
                             []int64{0, 1, 2},     /* Rows from F */
                             nil); err != nil { panic(err) } 

    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MINIMIZE); err != nil { panic(err) }

    fmt.Println("optimize")
    // Solve the problem 
    trm,err := task.Optimize(); if err != nil { panic(err) }
    fmt.Printf(" Mosek warning: %s\n",trm)
    // Print a summary containing information
    // about the solution for debugging purposes
    if err := task.SolutionSummary(mosek.MSK_STREAM_MSG); err != nil { panic(err) }

    /* Get status information about the solution */
    solsta,err := task.GetSolSta(mosek.MSK_SOL_ITR); if err != nil { panic(err) }

    xx,err := task.GetXx(mosek.MSK_SOL_ITR); if err != nil { panic(err) }

    switch solsta {
        case mosek.MSK_SOL_STA_OPTIMAL:
            fmt.Print("Optimal primal solution: %v\n",xx)
        case mosek.MSK_SOL_STA_DUAL_INFEAS_CER,mosek.MSK_SOL_STA_PRIM_INFEAS_CER:
            fmt.Println("Primal or dual infeasibility.");
        case mosek.MSK_SOL_STA_UNKNOWN:
            fmt.Println("Unknown solution status.");
        default:
            fmt.Printf("Other solution status: %s\n",solsta);
    }
}
