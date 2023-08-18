/*
   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.

   File :      acc1.go

   Purpose :   Tutorial example for affine conic constraints.
               Models the problem:
 
               maximize c^T x
               subject to  sum(x) = 1
                           gamma >= |Gx+h|_2
*/
package main 

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
)
const infinity = 0.0;

func main() {
    // create task object
    task,err := mosek.NewTask()
    if err != nil { panic(err) }

    /* Directs the log task stream to the 'printstr' function. */
    task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })
    
    var n int32 = 3
    var k int64 = 2

    // Create n free variables
    if err := task.AppendVars(n); err != nil { panic(err) }
    if err := task.PutVarBoundSliceConst(0, n, mosek.MSK_BK_FR, -infinity, infinity); err != nil { panic(err) }

    // Set up the objective
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE); err != nil { panic(err) }
    if err := task.PutCList([]int32{0, 1, 2}, []float64{2,3,-1}); err != nil { panic(err) }

    // One linear constraint - sum(x) = 1
    if err := task.AppendCons(1); err != nil { panic(err) }
    if err := task.PutConBound(0, mosek.MSK_BK_FX, 1.0, 1.0); err != nil { panic(err) }
    for i := int32(0); i < n; i++ {
        if err := task.PutAij(0, i, 1.0); err != nil { panic(err) }
    }

    // Append empty AFE rows for affine expression storage
    if err := task.AppendAfes(k + 1); err != nil { panic(err) }

    // F matix in sparse form
    Fsubi := []int64{1, 1, 2, 2}   // The G matrix starts in F from row 1
    Fsubj := []int32{0, 1, 0, 2}
    Fval  := []float64{1.5, 0.1, 0.3, 2.1}
    // Other data
    h     := []float64{0, 0.1}
    gamma := 0.03

    // Fill in F storage
    if err := task.PutAfeFEntryList(Fsubi, Fsubj, Fval); err != nil { panic(err) }

    // Fill in g storage;
    if err := task.PutAfeG(0, gamma); err != nil { panic(err) }
    if err := task.PutAfeGSlice(1, k+1, h); err != nil { panic(err) }

    // Define a conic quadratic domain
    quadDom,err := task.AppendQuadraticConeDomain(k + 1); if err != nil { panic(err) }

    // Create the ACC
    afeidx := []int64{0, 1, 2}
    if err := task.AppendAcc(quadDom, afeidx, nil); err != nil { panic(err) }

   
    task.WriteData("acc1.ptf")
    /* Solve the problem */
    trm,_ := task.Optimize()
    fmt.Printf(" Termination code: %s\n",trm)
    // Print a summary containing information
    // about the solution for debugging purposes
    task.SolutionSummary(mosek.MSK_STREAM_MSG)
    /* Get status information about the solution */
    solsta,_ := task.GetSolSta(mosek.MSK_SOL_ITR)

    switch (solsta) {
      case mosek.MSK_SOL_STA_OPTIMAL:
        // Fetch solution
        xx,_ := task.GetXx(mosek.MSK_SOL_ITR);
        fmt.Printf("Optimal primal solution: %v\n",xx);

        // Fetch doty dual for the ACC
        doty,_ := task.GetAccDotY(mosek.MSK_SOL_ITR,0)
        fmt.Printf("Dual doty value for the ACC: %v\n",doty);

        // Fetch ACC activity
        activity,_ := task.EvaluateAcc(mosek.MSK_SOL_ITR, 0)
        fmt.Printf("Activity for the ACC: %v\n",activity);
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
