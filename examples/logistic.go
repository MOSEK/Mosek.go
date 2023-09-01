//
//  Copyright: Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//  File:      logistic.go
//
// Purpose: Implements logistic regression with regulatization.
//
//          Demonstrates using the exponential cone and log-sum-exp in Optimizer API.

package main

import (
	"fmt"
	mosek "github.com/mosek/mosek.go"
)


// Adds ACCs for t_i >= log ( 1 + exp((1-2*y[i]) * theta' * X[i]) )
// Adds auxiliary variables, AFE rows and constraints
func SoftPlus(task *mosek.Task, d int, n int, theta int, t int, X [][]float64, y []bool) {
    const inf = 0.0
    nvar,err := task.GetNumVar(); if err != nil { panic(err) }
    ncon,err := task.GetNumCon(); if err != nil { panic(err) }
    nafe,err := task.GetNumAfe(); if err != nil { panic(err) }
    if err := task.AppendVars(int32(2*n)); err != nil { panic(err) } // z1, z2
    if err := task.AppendCons(int32(n));   err != nil { panic(err) } // z1 + z2 = 1
    if err := task.AppendAfes(int64(4*n)); err != nil { panic(err) } //theta * X[i] - t[i], -t[i], z1[i], z2[i]
    z1 := nvar
    z2 := nvar+int32(n)
    zcon     := ncon
    thetaafe := nafe 
    tafe     := nafe+int64(n)
    z1afe    := nafe+2*int64(n)
    z2afe    := nafe+3*int64(n)
    k := 0

    // Linear constraints
    subi := make([]int32,2*n)
    subj := make([]int32,2*n)
    aval := make([]float64,2*n);

    for i := 0; i < n; i++ {
      // z1 + z2 = 1
      subi[k] = zcon+int32(i);  subj[k] = z1+int32(i);  aval[k] = 1;  k++;
      subi[k] = zcon+int32(i);  subj[k] = z2+int32(i);  aval[k] = 1;  k++;
    }
    
    if err := task.PutAijList(subi, subj, aval); err != nil { panic(err) }
    if err := task.PutConBoundSliceConst(zcon, zcon+int32(n), mosek.MSK_BK_FX, 1, 1); err != nil { panic(err) }
    if err := task.PutVarBoundSliceConst(nvar, nvar+int32(2*n), mosek.MSK_BK_FR, -inf, inf); err != nil { panic(err) }
    
    // Affine conic expressions
    afeidx := make([]int64,d*n+4*n);
    varidx := make([]int32,d*n+4*n);
    fval   := make([]float64,d*n+4*n);
    k = 0

    // Thetas
    for i := 0; i < n; i++ {
        for j := 0; j < d; j++ {
            afeidx[k] = thetaafe + int64(i); varidx[k] = int32(theta + j); 
            if y[i] {
                fval[k] = - X[i][j]
            } else {
                fval[k] = X[i][j]
            }
            k++
        }
    }

    // -t[i]
    for i := 0; i < n; i++ {
        afeidx[k] = thetaafe + int64(i); varidx[k] = int32(t + i); fval[k] = -1; k++;
        afeidx[k] = tafe + int64(i);     varidx[k] = int32(t + i); fval[k] = -1; k++;
    }

    // z1, z2
    for i := 0; i < n; i++ {
        afeidx[k] = z1afe + int64(i); varidx[k] = z1 + int32(i); fval[k] = 1; k++;
        afeidx[k] = z2afe + int64(i); varidx[k] = z2 + int32(i); fval[k] = 1; k++;
    }

    // Add the expressions
    if err := task.PutAfeFEntryList(afeidx, varidx, fval); err != nil { panic(err) }

    // Add a single row with the constant expression "1.0"
    oneafe,err := task.GetNumAfe(); if err != nil { panic(err) }
    if err := task.AppendAfes(1); err != nil { panic(err) }
    if err := task.PutAfeG(oneafe, 1.0); err != nil { panic(err) }

    // Add an exponential cone domain
    expdomain,err := task.AppendPrimalExpConeDomain(); if err != nil { panic(err) }
    
    // Conic constraints
    numacc,err := task.GetNumAcc(); if err != nil { panic(err) }
    for i := 0; i < n; i++ {
        if     err := task.AppendAcc(expdomain, []int64{z1afe+int64(i), oneafe, thetaafe+int64(i)}, nil); err != nil { panic(err)
        } else if err := task.AppendAcc(expdomain, []int64{z2afe+int64(i), oneafe, tafe+int64(i)}, nil); err != nil { panic(err)
        } else if err := task.PutAccName(numacc+int64(i)*2,  fmt.Sprintf("z1:theta[%d]",i)); err != nil { panic(err)
        } else if err := task.PutAccName(numacc+int64(i)*2+1,fmt.Sprintf("z2:t[%d]",i)); err != nil { panic(err) 
        }
    }
}

// Model logistic regression (regularized with full 2-norm of theta)
// X - n x d matrix of data points
// y - length n vector classifying training points
// lamb - regularization parameter
func LogisticRegression(X [][]float64,
                        y []bool,
                        lamb float64) []float64 { 
    n := len(X)
    d := len(X[0])

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    defer task.Delete()
    // Variables [r; theta; t]
    nvar := 1+d+n
    if err := task.AppendVars(int32(nvar)); err != nil { panic(err) }
    if err := task.PutVarBoundSliceConst(0, int32(nvar), mosek.MSK_BK_FR, -inf, inf); err != nil { panic(err) }
    r := 0
    theta := 1
    t := 1+d
    if err := task.PutVarName(int32(r),"r"); err != nil { panic(err) }
    for i := 0; i < d; i++ { if err := task.PutVarName(int32(theta+i),fmt.Sprintf("theta[%d]",i)); err != nil { panic(err) } }
    for i := 0; i < n; i++ { if err := task.PutVarName(int32(t+i),    fmt.Sprintf("t[%d]",i)); err != nil { panic(err) } }

    // Objective lambda*r + sum(t)
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MINIMIZE); err != nil { panic(err) } 
    if err := task.PutCJ(int32(r), lamb); err != nil { panic(err) }
    for i := 0; i < n; i++ { if err := task.PutCJ(int32(t+i), 1.0); err != nil { panic(err) } }
    
    // Softplus function constraints
    SoftPlus(task, d, n, theta, t, X, y)

    // Regularization
    // Append a sequence of linear expressions (r, theta) to F
    numafe,err := task.GetNumAfe(); if err != nil { panic(err) }
    if err := task.AppendAfes(int64(1+d)); err != nil { panic(err) }
    if err := task.PutAfeFEntry(numafe, int32(r), 1.0); err != nil { panic(err) }
    for i := 0; i < d; i++ {
        if err := task.PutAfeFEntry(numafe + int64(i) + 1, int32(theta + i), 1.0); err != nil { panic(err) }
    }
    
    // Add the constraint
    if domidx,err := task.AppendQuadraticConeDomain(int64(1+d)); err != nil { panic(err) 
    } else if err := task.AppendAccSeq(domidx, numafe, nil); err != nil { panic(err) }

    // Solution
    if _,err := task.Optimize(); err != nil { panic(err) }
    xx,err := task.GetXxSlice(mosek.MSK_SOL_ITR, int32(theta), int32(theta+d)); if err != nil { panic(err) }
    return xx
}

func main() {
    // Test: detect and approximate a circle using degree 2 polynomials
    n  := 30
    
    X := make([][]float64,n*n); for i := 0; i < n*n; i++ { X[i] = make([]float64,6) }
    Y := make([]bool,n*n)

    for i := 0; i<n; i++ {
        for j := 0; j<n; j++ {
            k := i*n+j
            x := -1 + 2.0*float64(i)/float64(n-1);
            y := -1 + 2.0*float64(j)/float64(n-1);
            X[k][0] = 1.0; X[k][1] = x; X[k][2] = y; X[k][3] = x*y;
            X[k][4] = x*x; X[k][5] = y*y;
            Y[k] = (x*x+y*y>=0.69);
        }
    }

    theta := LogisticRegression(X, Y, 0.1)

    fmt.Printf("Thera: %v\n",theta)
}


