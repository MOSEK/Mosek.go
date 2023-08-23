//
//   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//   File :      pow1.go
//
//   Purpose: Demonstrates how to solve the problem
//
//   maximize x^0.2*y^0.8 + z^0.4 - x
//          st x + y + 0.5z = 2
//             x,y,z >= 0
//
package main

import (
    mosek "github.com/mosek/mosek.go"
    "fmt"
)

func main() {
    numcon := 1
    numvar := 5   // x,y,z and 2 auxiliary variables for conic constraints

    // Since the value infinity is never used, we define
    // 'infinity' symbolic purposes only
    infinity := 0.0

    val  := []float64{ 1.0, 1.0, -1.0 }
    sub  := []int32{ 3, 4, 0 }

    aval := []float64{ 1.0, 1.0, 0.5 }
    asub := []int32{ 0, 1, 2 }

    // create a new environment object
    task,err := mosek.NewTask(); if err != nil { panic(err) }
    // Directs the log task stream to the user specified
    // method task_msg_obj.stream
    if err := task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) }); err != nil { panic(err) }

    // Append 'numcon' empty constraints.
    // The constraints will initially have no bounds.
    if err := task.AppendCons(int32(numcon)); err != nil { panic(err) }

    // Append 'numvar' variables.
    // The variables will initially be fixed at zero (x=0).
    if err := task.AppendVars(int32(numvar)); err != nil { panic(err) }

    // Define the linear part of the problem
    if err := task.PutCList(sub, val); err != nil { panic(err) }
    if err := task.PutARow(0, asub, aval); err != nil { panic(err) }
    if err := task.PutConBound(0, mosek.MSK_BK_FX, 2.0, 2.0); err != nil { panic(err) }
    if err := task.PutVarBoundSliceConst(0, int32(numvar), mosek.MSK_BK_FR, -infinity, infinity); err != nil { panic(err) }

    // Add conic constraints
    // Append two power cone domains
    pc1,err := task.AppendPrimalPowerConeDomain(3, []float64{0.2, 0.8}); if err != nil { panic(err) }
    pc2,err := task.AppendPrimalPowerConeDomain(3, []float64{4.0, 6.0}); if err != nil { panic(err) }

    // Create data structures F,g so that
    //
    // F * x + g = (x(0), x(1), x(3), x(2), 1.0, x(4)) 
    if err := task.AppendAfes(6); err != nil { panic(err) }
    if err := task.PutAfeFEntryList([]int64{0, 1, 2, 3, 5},         // Rows 
                                    []int32{0, 1, 3, 2, 4},         // Columns 
                                    []float64{1.0, 1.0, 1.0, 1.0, 1.0}); err != nil { panic(err) }
    if err := task.PutAfeG(4, 1.0); err != nil { panic(err) }

    // Append the two conic constraints 
    if err := task.AppendAcc(pc1,                  // Domain 
                             []int64{0, 1, 2},     // Rows from F 
                             nil); err != nil { panic(err) }
    if err := task.AppendAcc(pc2,                  // Domain 
                             []int64{3, 4, 5},     // Rows from F 
                             nil); err != nil { panic(err) }

    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE); err != nil { panic(err) }

    fmt.Println ("optimize")
    // Solve the problem 
    if trm,err := task.Optimize(); err != nil { panic(err) 
    } else if trm != mosek.MSK_RES_OK {
        fmt.Printf(" Mosek warning: %s\n",trm)
    }
    // Print a summary containing information
    // about the solution for debugging purposes
    if err := task.SolutionSummary(mosek.MSK_STREAM_MSG); err != nil { panic(err) }

    // Get status information about the solution 
    if solsta,err := task.GetSolSta(mosek.MSK_SOL_ITR); err != nil {
        panic(err)
    } else if xx,err := task.GetXx(mosek.MSK_SOL_ITR);  err != nil { // Interior solution
        panic(err)
    } else {
        switch solsta {
          case mosek.MSK_SOL_STA_OPTIMAL:
              fmt.Println("Optimal primal solution: %v\n",xx);
          case mosek.MSK_SOL_STA_DUAL_INFEAS_CER,mosek.MSK_SOL_STA_PRIM_INFEAS_CER:
            fmt.Println("Primal or dual infeasibility.\n");
          case mosek.MSK_SOL_STA_UNKNOWN:
            fmt.Println("Unknown solution status.\n")
          default:
            fmt.Println("Other solution status")
        }
    }
}
