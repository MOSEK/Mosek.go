// Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
// File :      sdo1.go
//
// Purpose :   Solves the following small semidefinite optimization problem
//             using the MOSEK API.
//
//             minimize    Tr [2, 1, 0; 1, 2, 1; 0, 1, 2]*X + x0
//
//             subject to  Tr [1, 0, 0; 0, 1, 0; 0, 0, 1]*X + x0           = 1
//                         Tr [1, 1, 1; 1, 1, 1; 1, 1, 1]*X      + x1 + x2 = 0.5
//                         (x0,x1,x2) \in Q,  X \in PSD


package main

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
)

func main() {
    numcon    := 2  // Number of constraints.              
    numvar    := 3  // Number of scalar variables 
    dimbarvar := []int32{3}         // Dimension of semidefinite cone 

    bkc := []mosek.Boundkey{ mosek.MSK_BK_FX, mosek.MSK_BK_FX }
    blc := []float64{ 1.0, 0.5 }
    buc := []float64{ 1.0, 0.5 }

    barc_i := []int32{0, 1, 1, 2, 2}
    barc_j := []int32{0, 0, 1, 1, 2}
    barc_v := []float64{2.0, 1.0, 2.0, 1.0, 2.0}

    asub   := [][]int32{ []int32{0}, []int32{1, 2}} // column subscripts of A 
    aval   := [][]float64{ []float64{1.0}, []float64{1.0, 1.0} }

    bara_i := [][]int32{ []int32{0,   1,   2}, []int32{0,   1 ,  2,   1,   2,   2 } }
    bara_j := [][]int32{ []int32{0,   1,   2}, []int32{0,   0 ,  0,   1,   1,   2 } }
    
    bara_v := [][]float64{ []float64{1.0, 1.0, 1.0}, []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}}

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    // Directs the log task stream to the user specified
    // method task_msg_obj.stream
    if err := task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) }); err != nil { panic(err) }

    // Append 'NUMCON' empty constraints.
    // The constraints will initially have no bounds. 
    if err := task.AppendCons(int32(numcon)); err != nil { panic(err) }

    // Append 'NUMVAR' variables.
    // The variables will initially be fixed at zero (x=0). 
    if err := task.AppendVars(int32(numvar)); err != nil { panic(err) }

    // Append 'NUMBARVAR' semidefinite variables. 
    if err := task.AppendBarvars(dimbarvar); err != nil { panic(err) }

    // Optionally add a constant term to the objective. 
    if err := task.PutCfix(0.0); err != nil { panic(err) }

    // Set the linear term c_j in the objective.
    if err := task.PutCJ(0, 1.0); err != nil { panic(err) }

    for j := 0; j < numvar; j++ {
        if err := task.PutVarBound(int32(j), mosek.MSK_BK_FR, -0.0, 0.0); err != nil { panic(err) }
    }

    // Set the linear term barc_j in the objective.
    {
        idx,err := task.AppendSparseSymMat(dimbarvar[0],barc_i,barc_j,barc_v); if err != nil { panic(err) }
        if err := task.PutBarcJ(0, []int64{idx}, []float64{1.0}); err != nil { panic(err) }
    }

    // Set the bounds on constraints.
    // for i=1, ...,numcon : blc[i] <= constraint i <= buc[i] 

    for i := 0; i < numcon; i++ {
        if err := task.PutConBound(int32(i),   // Index of constraint.
                                   bkc[i],     // Bound key.
                                   blc[i],     // Numerical value of lower bound.
                                   buc[i]); err != nil {   // Numerical value of upper bound.
            panic(err)
        }
    }

    // Input A row by row 
    for i := 0; i < numcon; i++ {
        if err := task.PutARow(int32(i), asub[i], aval[i]); err != nil { panic(err) }
    }

    // Append the conic quadratic constraint 
    if err := task.AppendAfes(3); err != nil { panic(err) }
    // Diagonal F matrix
    if err := task.PutAfeFEntryList([]int64{0, 1, 2}, []int32{0, 1, 2}, []float64{1.0,1.0,1.0}); err != nil { panic(err) }
    if dimidx,err := task.AppendQuadraticConeDomain(3); err != nil { panic(err)
    } else if err := task.AppendAccSeq(dimidx, 0, nil); err != nil { panic(err) }

    // Add the first row of barA 
    if idx,err := task.AppendSparseSymMat(dimbarvar[0], bara_i[0], bara_j[0], bara_v[0]); err != nil { 
        panic(err) 
    } else if err := task.PutBaraIj(0, 0, []int64{idx}, []float64{1.0}); err != nil {
        panic(err) 
    }

    // Add the second row of barA 
    if idx,err := task.AppendSparseSymMat(dimbarvar[0], bara_i[1], bara_j[1], bara_v[1]); err != nil { 
        panic(err) 
    } else if err := task.PutBaraIj(1, 0, []int64{idx}, []float64{1.0}); err != nil { 
        panic(err) 
    }

    // Run optimizer 
    if _,err := task.Optimize(); err != nil { panic(err) }
    
    // Print a summary containing information
    // about the solution for debugging purposes
    if err := task.SolutionSummary (mosek.MSK_STREAM_MSG); err != nil { panic(err) }

    if solsta,err := task.GetSolSta(mosek.MSK_SOL_ITR); err != nil { 
        panic(err) 
    } else {
        switch solsta {
            case mosek.MSK_SOL_STA_OPTIMAL:
                xx,err := task.GetXx(mosek.MSK_SOL_ITR); if err != nil { panic(err) }
                barx,err := task.GetBarxJ(mosek.MSK_SOL_ITR, 0); if err != nil { panic(err) }    // Request the interior solution. 
                fmt.Printf("Optimal primal solution: \n  xx   = %v\n  barx = %v\n",xx,barx);
                break;
            case mosek.MSK_SOL_STA_DUAL_INFEAS_CER:
            case mosek.MSK_SOL_STA_PRIM_INFEAS_CER:
                fmt.Println("Primal or dual infeasibility certificate found.");
            case mosek.MSK_SOL_STA_UNKNOWN:
                fmt.Println("The status of the solution could not be determined.");
                break;
            default:
                fmt.Println("Other solution status.");
                break;
        }
    }
}
