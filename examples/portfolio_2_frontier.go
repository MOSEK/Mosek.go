//
//   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//   File :      portfolio_2_frontier.go
//
//   Purpose :   Implements a basic portfolio optimization model.
//               Computes points on the efficient frontier.
//
package main

import (
    mosek "github.com/mosek/mosek.go"
    "fmt"
    "math"
)

func main() {
    // Since the value infinity is never used, we define
    // 'infinity' symbolic purposes only
    const infinity = 0.0;
    const n = 8
    mu := []float64{0.07197, 0.15518, 0.17535, 0.08981, 0.42896, 0.39292, 0.32171, 0.18379}
    GT := []float64{
      0.30758, 0.12146, 0.11341, 0.11327, 0.17625, 0.11973, 0.10435, 0.10638,
      0.     , 0.25042, 0.09946, 0.09164, 0.06692, 0.08706, 0.09173, 0.08506,
      0.     , 0.     , 0.19914, 0.05867, 0.06453, 0.07367, 0.06468, 0.01914,
      0.     , 0.     , 0.     , 0.20876, 0.04933, 0.03651, 0.09381, 0.07742,
      0.     , 0.     , 0.     , 0.     , 0.36096, 0.12574, 0.10157, 0.0571 ,
      0.     , 0.     , 0.     , 0.     , 0.     , 0.21552, 0.05663, 0.06187,
      0.     , 0.     , 0.     , 0.     , 0.     , 0.     , 0.22514, 0.03327,
      0.     , 0.     , 0.     , 0.     , 0.     , 0.     , 0.     , 0.2202 }
    //int   k = GT.length;
    k := len(GT) / n

    x0 := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    w := 1.0
    alphas := []float64{0.0, 0.01, 0.1, 0.25, 0.30, 0.35, 0.4, 0.45, 0.5, 0.75, 1.0, 1.5, 2.0, 3.0, 10.0}
   
    //Offset of variables into the API variable.
    numvar := int32(n + 1)
    voff_x := int32(0)
    voff_s := int32(n)

    // Offset of constraints
    const coff_bud = 0

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })
    // Directs the log task stream to the user specified
    // method task_msg_obj.stream
    
    if err := task.AppendVars(numvar); err != nil { panic(err) }

    // Setting up variable x 
    for j := 0; j < n; j++ {
        /* Optionally we can give the variables names */
        if err := task.PutVarName(voff_x + int32(j), fmt.Sprintf("x[%d]",j + 1)); err != nil { panic(err) }
        /* No short-selling - x^l = 0, x^u = inf */
        if err := task.PutVarBound(voff_x + int32(j), mosek.MSK_BK_LO, 0.0, infinity); err != nil { panic(err) }
    }
    if err := task.PutVarName(voff_s, "s"); err != nil { panic(err) }
    if err := task.PutVarBound(voff_s, mosek.MSK_BK_FR, -infinity, infinity); err != nil { panic(err) }

    // One linear constraint: total budget
    if err := task.AppendCons(1); err != nil { panic(err) }
    if err := task.PutConName(coff_bud, "budget"); err != nil { panic(err) }
    for j := 0; j < n; j++ {
        /* Coefficients in the first row of A */
        if err := task.PutAij(coff_bud, voff_x + int32(j), 1.0); err != nil { panic(err) }
    }

    totalBudget := w
    for _,x0j := range x0 { totalBudget += x0j }
    if err := task.PutConBound(coff_bud, mosek.MSK_BK_FX, totalBudget, totalBudget); err != nil { panic(err) }

    // Input (gamma, GTx) in the AFE (affine expression) storage
    // We build the following F and g for variables [x, s]:
    //     [0, 1]      [0  ]
    // F = [0, 0], g = [0.5]
    //     [GT,0]      [0  ]
    // We need k+2 rows
    if err := task.AppendAfes(int64(k) + 2); err != nil { panic(err) }
    // The first affine expression is variable s (last variable, index n)
    if err := task.PutAfeFEntry(0, n, 1.0); err != nil { panic(err) }        
    // The second affine expression is constant 0.5
    if err := task.PutAfeG(1, 0.5); err != nil { panic(err) }
    // The remaining k expressions comprise GT*x, we add them row by row
    // In more realisic scenarios it would be better to extract nonzeros and input in sparse form

    vslice_x := make([]int32,n)
    for i := 0; i < n; i++ { vslice_x[i] = voff_x + int32(i) } 
    for i := 0; i < k; i++ {
        if err := task.PutAfeFRow(int64(i) + 2, vslice_x, GT[i*n:(i+1)*n]); err != nil { panic(err) }
    }

    // Input the affine conic constraint (gamma, GT*x) \in QCone
    // Add the quadratic domain of dimension k+1
    if rqdom,err := task.AppendRQuadraticConeDomain(int64(k) + 2); err != nil { 
        panic(err) 
    // Add the constraint
    } else if err := task.AppendAccSeq(rqdom, 0, nil); err != nil { panic(err) }
    if err := task.PutAccName(0, "risk"); err != nil { panic(err) }

    // Objective: maximize expected return mu^T x
    for j,muj := range mu {
        if err := task.PutCJ(voff_x + int32(j), muj); err != nil { panic(err) }
    }
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE); err != nil { panic(err) }

    //Turn all log output off.
    if err := task.PutIntParam(mosek.MSK_IPAR_LOG, 0); err != nil { panic(err) }

    fmt.Printf("%-12s  %-12s  %-12s\n", "alpha", "exp ret", "std. dev.");
    for _,alpha := range alphas {
        if err := task.PutCJ(voff_s, -alpha); err != nil { panic(err) }

        if _,err := task.Optimize(); err != nil { panic(err) }

        if err := task.SolutionSummary(mosek.MSK_STREAM_LOG); err != nil { panic(err) }

        if xx,err := task.GetXx(mosek.MSK_SOL_ITR); err != nil { 
            panic(err) 
        } else {
            expret := 0.0
            for jj := 0; jj < n; jj++ { expret += mu[jj] * xx[jj + int(voff_x)] }
            fmt.Printf("%-12.3e  %-12.3e  %-12.3e\n", alpha, expret, math.Sqrt(xx[voff_s]))
        }
    }
}
