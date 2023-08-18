/*
   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.

   File :      acc2.go

   Purpose :   Tutorial example for affine conic constraints.
               Models the problem:
 
               maximize c^T x
               subject to  sum(x) = 1
                           gamma >= |Gx+h|_2

               This version inputs the linear constraint as an affine conic constraint.
*/
package main

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
)

/* Data dimensions */
const n int32 = 3
const k int64 = 2
const infinity = 0.0

func main() {
    task,err := mosek.NewTask(); if err != nil { panic(err) }

    /* Directs the log task stream to the 'printstr' function. */
    task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })

    // Create n free variables
    if err := task.AppendVars(n); err != nil { panic(err) }
    if err := task.PutVarBoundSliceConst(0, n, mosek.MSK_BK_FR, -infinity, infinity); err != nil { panic(err) }

    // Set up the objective
    var c    = []float64{2, 3, -1}
    var cind = []int32{0, 1, 2}  
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE); err != nil { panic(err) }
    if err := task.PutCList(cind, c); err != nil { panic(err) }

    // Set AFE rows representing the linear constraint
    if err := task.AppendAfes(1); err != nil { panic(err) }
    if err := task.PutAfeG(0, -1.0); err != nil { panic(err) }
    for i := int32(0); i < n; i++ {  task.PutAfeFEntry(0, i, 1.0); }

    // Set AFE rows representing the quadratic constraint
    // F matix in sparse form
    Fsubi := []int64{2, 2, 3, 3}   // The G matrix starts in F from row 2
    Fsubj := []int32{0, 1, 0, 2}
    Fval  := []float64{1.5, 0.1, 0.3, 2.1}
    // Other data
    h     := []float64{0, 0.1}
    var gamma float64 = 0.03

    if err := task.AppendAfes(k + 1); err != nil { panic(err) }
    if err := task.PutAfeFEntryList(Fsubi, Fsubj, Fval); err != nil { panic(err) }
    if err := task.PutAfeG(1, gamma); err != nil { panic(err) }
    if err := task.PutAfeGSlice(2, k+2, h); err != nil { panic(err) }

    // Define domains
    zeroDom,err := task.AppendRzeroDomain(1); if err != nil { panic(err) }
    quadDom,err := task.AppendQuadraticConeDomain(k + 1); if err != nil { panic(err) }

    // Create the linear ACC
    afeidxZero := []int64{0};
    if err := task.AppendAcc(zeroDom,    // Domain index
                             afeidxZero, // Indices of AFE rows
                             nil); err != nil { panic(err) }

    // Create the quadratic ACC
    afeidxQuad := []int64{1, 2, 3}
    if err := task.AppendAcc(quadDom,    // Domain index
                             afeidxQuad, // Indices of AFE rows
                             nil); err != nil { panic(err) }

    
    /* Solve the problem */
    trm,err := task.Optimize(); if err != nil { panic(err) }
    fmt.Printf(" Termination code: %s",trm)
    // Print a summary containing information
    // about the solution for debugging purposes
    if err := task.SolutionSummary(mosek.MSK_STREAM_LOG); err != nil { panic(err) }

    /* Get status information about the solution */
    solsta,err := task.GetSolSta(mosek.MSK_SOL_ITR); if err != nil { panic(err) }

    switch solsta {
    case mosek.MSK_SOL_STA_OPTIMAL:
        // Fetch solution
        if xx,err := task.GetXx(mosek.MSK_SOL_ITR); err != nil {
            panic(err)
        } else if doty,err := task.GetAccDotY(mosek.MSK_SOL_ITR, 1); err != nil { 
            panic(err)
        } else if activity,err := task.EvaluateAcc(mosek.MSK_SOL_ITR,1); err != nil {
            panic(err)
        } else {
            fmt.Printf("Optimal primal solution: %v\n",xx);
            fmt.Printf("Dual doty value for the ACC: %v\n",doty);
            fmt.Printf("Activity for the ACC: %v\n",activity);
        }
        break;
    case mosek.MSK_SOL_STA_DUAL_INFEAS_CER:
    case mosek.MSK_SOL_STA_PRIM_INFEAS_CER:
        fmt.Println("Primal or dual infeasibility.");
        break;
    case mosek.MSK_SOL_STA_UNKNOWN:
        fmt.Println("Unknown solution status.");
        break;
    default:
        fmt.Println("Other solution status");
        break;
    }
}

