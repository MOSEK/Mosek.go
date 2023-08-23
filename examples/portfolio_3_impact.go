//
//   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//   File :      portfolio_3_impact.java
//
//   Purpose :   Implements a basic portfolio optimization model
//               with transaction costs of the form x^(3/2).
//
package main
import (
    mosek "github.com/mosek/mosek.go"
    "fmt"
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
    k := len(GT)/n
    x0 := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0};
    w := 1.0;
    gamma := 0.36;
  
    m := make([]float64,n)
    for i := 0; i < n; i++ { m[i] = 0.01 } 

    // Offset of variables into the API variable.
    const numvar = n*3
    const voff_x = 0
    const voff_c = n
    const voff_z = 2 * n

    // Offset of constraints.
    const numcon = 2 * n + 1
    const coff_bud = 0
    const coff_abs1 = 1
    const coff_abs2 = 1 + n

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    // Directs the log task stream to the user specified
    // method task_msg_obj.stream
    if err := task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) }); err != nil { panic(err) }

    // Variables (vector of x, c, z)
    if err := task.AppendVars(numvar); err != nil { panic(err) }
    for j := 0; j < n; j++ {
      /* Optionally we can give the variables names */
      if err := task.PutVarName(voff_x + int32(j), fmt.Sprintf("x[%d]",j + 1)); err != nil { panic(err) }
      if err := task.PutVarName(voff_c + int32(j), fmt.Sprintf("c[%d]",j + 1)); err != nil { panic(err) }
      if err := task.PutVarName(voff_z + int32(j), fmt.Sprintf("z[%d]",j + 1)); err != nil { panic(err) }
      /* Apply variable bounds (x >= 0, c and z free) */
      if err := task.PutVarBound(voff_x + int32(j), mosek.MSK_BK_LO, 0.0, infinity); err != nil { panic(err) }
      if err := task.PutVarBound(voff_c + int32(j), mosek.MSK_BK_FR, -infinity, infinity); err != nil { panic(err) }
      if err := task.PutVarBound(voff_z + int32(j), mosek.MSK_BK_FR, -infinity, infinity); err != nil { panic(err) }
    }
            
    // Linear constraints
    // - Total budget
    if err := task.AppendCons(1); err != nil { panic(err) }
    if err := task.PutConName(coff_bud, "budget"); err != nil { panic(err) }
    for j := 0; j < n; j++ {
      /* Coefficients in the first row of A */
      if err := task.PutAij(coff_bud, voff_x + int32(j), 1.0); err != nil { panic(err) }
      if err := task.PutAij(coff_bud, voff_c + int32(j), m[j]); err != nil { panic(err) }
    }
    totalBudget := w
    for _,x0j := range x0 { totalBudget += x0j; }
    if err := task.PutConBound(coff_bud, mosek.MSK_BK_FX, totalBudget, totalBudget); err != nil { panic(err) }

    // - Absolute value
    if err := task.AppendCons(2 * n); err != nil { panic(err) }
    for i := 0; i < n; i++ {
      if err := task.PutConName(coff_abs1 + int32(i), fmt.Sprintf("zabs1[%d]", 1 + i)); err != nil { panic(err) }
      if err := task.PutAij(coff_abs1 + int32(i), voff_x + int32(i), -1.0); err != nil { panic(err) }
      if err := task.PutAij(coff_abs1 + int32(i), voff_z + int32(i), 1.0); err != nil { panic(err) }
      if err := task.PutConBound(coff_abs1 + int32(i), mosek.MSK_BK_LO, -x0[i], infinity); err != nil { panic(err) }
      if err := task.PutConName(coff_abs2 + int32(i), fmt.Sprintf("zabs2[%d]",1 + i)); err != nil { panic(err) }
      if err := task.PutAij(coff_abs2 + int32(i), voff_x + int32(i), 1.0); err != nil { panic(err) }
      if err := task.PutAij(coff_abs2 + int32(i), voff_z + int32(i), 1.0); err != nil { panic(err) }
      if err := task.PutConBound(coff_abs2 + int32(i), mosek.MSK_BK_LO, x0[i], infinity);   err != nil { panic(err) }        
    }

    // ACCs
    const aoff_q = 0
    aoff_pow := int64(k + 1)
    // - (gamma, GTx) in Q(k+1)
    // The part of F and g for variable x:
    //     [0,  0, 0]      [gamma]
    // F = [GT, 0, 0], g = [0    ]    
    if err := task.AppendAfes(int64(k + 1)); err != nil { panic(err) }
    if err := task.PutAfeG(aoff_q, gamma); err != nil { panic(err) }
    vslice_x := make([]int32,n)
    for i := 0; i < n; i++ { vslice_x[i] = voff_x + int32(i) } 
    for i := 0; i < k; i++ {
        if err := task.PutAfeFRow(aoff_q + int64(i) + 1, vslice_x, GT[i*n:(i+1)*n]); err != nil { panic(err) }
    }
    if qdom,err := task.AppendQuadraticConeDomain(int64(k + 1)); err != nil { 
        panic(err) 
    } else if err := task.AppendAccSeq(qdom, aoff_q, nil); err != nil { 
        panic(err) 
    }
    if err := task.PutAccName(aoff_q, "risk"); err != nil { panic(err) }

    // - (c_j, 1, z_j) in P3(2/3, 1/3)
    // The part of F and g for variables [c, z]:
    //     [0, I, 0]      [0]
    // F = [0, 0, I], g = [0]
    //     [0, 0, 0]      [1]
    if err := task.AppendAfes(2 * n + 1); err != nil { panic(err) }
    for i := 0; i < n; i++ {
        if err := task.PutAfeFEntry(aoff_pow + int64(i), voff_c + int32(i), 1.0); err != nil { panic(err) }
        if err := task.PutAfeFEntry(aoff_pow + int64(n + i), voff_z + int32(i), 1.0); err != nil { panic(err) }
    } 
    if err := task.PutAfeG(aoff_pow + 2 * n, 1.0); err != nil { panic(err) }
    // We use one row from F and g for both c_j and z_j, and the last row of F and g for the constant 1.
    // NOTE: Here we reuse the last AFE and the power cone n times, but we store them only once.
    exponents := []float64{2, 1};
    if powdom,err := task.AppendPrimalPowerConeDomain(3, exponents); err != nil { 
        panic(err) 
    } else {
        flat_afe_list := make([]int64,3*n)
        dom_list := make([]int64,n)
        for i := 0; i < n; i++ {
            flat_afe_list[3 * i + 0] = aoff_pow + int64(i);
            flat_afe_list[3 * i + 1] = aoff_pow + int64(2 * n);
            flat_afe_list[3 * i + 2] = aoff_pow + int64(n + i);
            dom_list[i] = powdom
        }
        if err := task.AppendAccs(dom_list, flat_afe_list, nil); err != nil { panic(err) }
    }
    for i := 0; i < n; i++ {
        if err := task.PutAccName(int64(i + 1), fmt.Sprintf("market_impact[%d]",i+1)); err != nil { panic(err) }
    }             

    // Objective: maximize expected return mu^T x
    for j,muj := range mu {
        if err := task.PutCJ(voff_x + int32(j), muj); err != nil { panic(err) }
    }
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE); err != nil { panic(err) }
    
    /* Solve the problem */
    //Turn all log output off.
    //task.putintparam(mosek.iparam.log,0);

    if _,err := task.Optimize(); err != nil { panic(err) }

    if err := task.SolutionSummary(mosek.MSK_STREAM_LOG); err != nil { panic(err) }


    if xx,err := task.GetXx(mosek.MSK_SOL_ITR); err != nil { 
        panic(err) 
    } else {
        expret := 0.0
        for j,muj := range mu { expret += muj * xx[j + int(voff_x)] }

        fmt.Printf("Expected return %e for gamma %e\n\n", expret, gamma)
        // Expected return: 4.1657e-01 Std. deviation: 3.6000e-01 Market impact cost: 7.6305e-03
    }
}
