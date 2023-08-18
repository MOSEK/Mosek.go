/*
   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.

   File :      portfolio_1_basic.go

   Purpose :   Implements a basic portfolio optimization model.
*/
package main

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
)

func main() {
    // Since the value infinity is never used, we define
    // 'infinity' for symbolic purposes only

    var n int32 = 8
    const infinity = 0.0
    const gamma = 36.0
    var mu = []float64{0.07197349, 0.15518171, 0.17535435, 0.0898094 , 0.42895777, 0.39291844, 0.32170722, 0.18378628}
    var GT = []float64{
        0.30758, 0.12146, 0.11341, 0.11327, 0.17625, 0.11973, 0.10435, 0.10638,
        0.     , 0.25042, 0.09946, 0.09164, 0.06692, 0.08706, 0.09173, 0.08506,
        0.     , 0.     , 0.19914, 0.05867, 0.06453, 0.07367, 0.06468, 0.01914,
        0.     , 0.     , 0.     , 0.20876, 0.04933, 0.03651, 0.09381, 0.07742,
        0.     , 0.     , 0.     , 0.     , 0.36096, 0.12574, 0.10157, 0.0571 ,
        0.     , 0.     , 0.     , 0.     , 0.     , 0.21552, 0.05663, 0.06187,
        0.     , 0.     , 0.     , 0.     , 0.     , 0.     , 0.22514, 0.03327,
        0.     , 0.     , 0.     , 0.     , 0.     , 0.     , 0.     , 0.2202 }
    var k int64 = int64(len(GT) / int(n))
    var x0 = []float64{8.0, 5.0, 3.0, 5.0, 2.0, 9.0, 3.0, 6.0}
    const w = 59;

    // Offset of variables into the API variable.
    numvar := n
    const voff_x int32 = 0

    // Constraints offsets
    const coff_bud int32 = 0;

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    defer task.Delete()
        
    // Directs the log task stream to the user specified
    // method task_msg_obj.stream
    task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })
    
    // Holding variable x of length n
    // No other auxiliary variables are needed in this formulation
    if err := task.AppendVars(numvar); err != nil { panic(err) }
    
    // Setting up variable x 
    for j := int32(0); j < n; j++ {
        /* Optionally we can give the variables names */
        if err := task.PutVarName(voff_x + j, fmt.Sprintf("x[%d]",j + 1)); err != nil { panic(err) }
        /* No short-selling - x^l = 0, x^u = inf */
        if err := task.PutVarBound(voff_x + j, mosek.MSK_BK_LO, 0.0, infinity); err != nil { panic(err) }
    }

    // One linear constraint: total budget
    if err := task.AppendCons(1); err != nil { panic(err) }
    if err := task.PutConName(coff_bud, "budget"); err != nil { panic(err) }
    for j := int32(0); j < n; j++ {
      /* Coefficients in the first row of A */
      if err := task.PutAij(coff_bud, voff_x + j, 1.0); err != nil { panic(err) }
    }
    var totalBudget float64 = w
    for _,x0j := range(x0) { totalBudget += x0j }
    if err := task.PutConBound(coff_bud, mosek.MSK_BK_FX, totalBudget, totalBudget); err != nil { panic(err) }

    // Input (gamma, GTx) in the AFE (affine expression) storage
    // We need k+1 rows
    if err := task.AppendAfes(k + 1); err != nil { panic(err) }
    // The first affine expression = gamma
    if err := task.PutAfeG(0, gamma); err != nil { panic(err) }
    // The remaining k expressions comprise GT*x, we add them row by row
    // In more realisic scenarios it would be better to extract nonzeros and input in sparse form
    var vslice_x = make([]int32,n);
    for i := int32(0); i < n; i++ { vslice_x[i] = voff_x + int32(i) }
    for i := 0; i < int(k); i++ { 
        GTi := GT[i*int(n):(i+1)*int(n)]
        if err := task.PutAfeFRow(int64(i) + 1, vslice_x, GTi); err != nil { panic(err) }
    }
    
    // Input the affine conic constraint (gamma, GT*x) \in QCone
    // Add the quadratic domain of dimension k+1
    qdom,err := task.AppendQuadraticConeDomain(k + 1)
    if err != nil { panic(err) }
    // Add the constraint
    if err := task.AppendAccSeq(qdom, 0, nil); err != nil { panic(err) }
    if err := task.PutAccName(0, "risk"); err != nil { panic(err) }
    
    // Objective: maximize expected return mu^T x
    for j,mu_j := range mu {
        if err := task.PutCJ(voff_x + int32(j), mu_j); err != nil { panic(err) }
    }
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE); err != nil { panic(err) }
    
    if _,err := task.Optimize(); err != nil { panic(err) }

    /* Display solution summary for quick inspection of results */
    if err := task.SolutionSummary(mosek.MSK_STREAM_LOG); err != nil { panic(err) }

    task.WriteData("portfolio-1.ptf")

    /* Read the results */
    if xx,err := task.GetXxSlice(mosek.MSK_SOL_ITR, voff_x, voff_x + int32(n)); err != nil { 
        panic(err) 
    } else {
        expret := 0.0
        for j,mu_j := range mu { expret += mu_j * xx[j] }
        fmt.Printf("\nExpected return %.4e for gamma %.4e\n", expret, gamma)
    }
}
