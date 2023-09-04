//
//   Copyright: Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//   File:      gp1.go
//
//   Purpose:   Demonstrates how to solve a simple Geometric Program (GP)
//              cast into conic form with exponential cones and log-sum-exp.
//
//              Example from
//                https://gpkit.readthedocs.io/en/latest/examples.html//maximizing-the-volume-of-a-box
//

package main

import (
	"fmt"
        "math"
	mosek "github.com/mosek/mosek.go"
)

// Since the value of infinity is ignored, we define it solely
// for symbolic purposes
const inf float64 = 0.0

  // maximize     h*w*d
  // subjecto to  2*(h*w + h*d) <= Awall
  //              w*d <= Afloor
  //              alpha <= h/w <= beta
  //              gamma <= d/w <= delta
  //
  // Variable substitutions:  h = exp(x), w = exp(y), d = exp(z).
  //
  // maximize     x+y+z
  // subject      log( exp(x+y+log(2/Awall)) + exp(x+z+log(2/Awall)) ) <= 0
  //                              y+z <= log(Afloor)
  //              log( alpha ) <= x-y <= log( beta )
  //              log( gamma ) <= z-y <= log( delta )
func max_volume_box(Aw    float64,
                    Af    float64,
                    alpha float64,
                    beta  float64,
                    gamma float64,
                    delta float64) []float64 {
    // Basic dimensions of our problem
    var numvar int32   = 3;  // Variables in original problem
    var x int32 = 0 
    var y int32 = 1
    var z int32 = 2  // Indices of variables
    var numcon int32 = 3;  // Linear constraints in original problem

    // Linear part of the problem involving x, y, z
    cval  := []float64{1, 1, 1}
    asubi := []int32{0, 0, 1, 1, 2, 2}
    asubj := []int32{y, z, x, y, z, y}
    aval  := []float64{1.0, 1.0, 1.0, -1.0, 1.0, -1.0}
    bkc   := []mosek.Boundkey{mosek.MSK_BK_UP,mosek.MSK_BK_RA,mosek.MSK_BK_RA}
    blc   := []float64{-inf, math.Log(alpha), math.Log(gamma)}
    buc   := []float64{math.Log(Af), math.Log(beta), math.Log(delta)}

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    defer task.Delete()
    // Directs the log task stream to the user specified
    // method task_msg_obj.stream
    if err := task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) }); err != nil { panic(err) }
    // Add variables and constraints
    if err := task.AppendVars(numvar); err != nil { panic(err) }
    if err := task.AppendCons(numcon); err != nil { panic(err) }

    // Objective is the sum of three first variables
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MINIMIZE); err != nil { panic(err) }
    if err := task.PutCSlice(0, numvar, cval); err != nil { panic(err) }
    if err := task.PutVarBoundSliceConst(0, numvar, mosek.MSK_BK_FR, -inf, inf); err != nil { panic(err) }

    // Add the linear constraints
    if err := task.PutAijList(asubi, asubj, aval); err != nil { panic(err) }
    if err := task.PutConBoundSlice(0, numcon, bkc, blc, buc); err != nil { panic(err) }

    // Affine expressions appearing in affine conic constraints
    // in this order:
    // u1, u2, x+y+log(2/Awall), x+z+log(2/Awall), 1.0, u1+u2-1.0
    const numafe int64 = 6
    const u1 int32 = 3
    const u2 int32 = 4     // Indices of slack variables
    afeidx := []int64{0, 1, 2, 2, 3, 3, 5, 5}
    varidx := []int32{u1, u2, x, y, x, z, u1, u2}
    fval   := []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}
    gfull  := []float64{0, 0, math.Log(2/Aw), math.Log(2/Aw), 1.0, -1.0}

    // New variables u1, u2
    if err := task.AppendVars(2); err != nil { panic(err) }
    if err := task.PutVarBoundSliceConst(u1, u2+1, mosek.MSK_BK_FR, -inf, inf); err != nil { panic(err) }

    // Append affine expressions
    if err := task.AppendAfes(numafe); err != nil { panic(err) }
    if err := task.PutAfeFEntryList(afeidx, varidx, fval); err != nil { panic(err) }
    if err := task.PutAfeGSlice(0, numafe, gfull); err != nil { panic(err) }

    // Two affine conic constraints
    if expdom,err := task.AppendPrimalExpConeDomain(); err != nil { panic(err) 
        // (u1, 1, x+y+log(2/Awall)) \in EXP
    } else if err := task.AppendAcc(expdom, []int64{0, 4, 2}, nil); err != nil { panic(err) 
        // (u2, 1, x+z+log(2/Awall)) \in EXP
    } else if err := task.AppendAcc(expdom, []int64{1, 4, 3}, nil);      err != nil { panic(err) }

    // The constraint u1+u2-1 \in \ZERO is added also as an ACC
    if domidx,err := task.AppendRzeroDomain(1); err != nil { panic(err)
    } else if err := task.AppendAcc(domidx, []int64{5}, nil); err != nil { panic(err) }

    // Solve and map to original h, w, d
    if _,err := task.Optimize(); err != nil { panic(err) }
    xyz,err := task.GetXxSlice(mosek.MSK_SOL_ITR, 0, numvar); if err != nil { panic(err) }
    hwd := make([]float64,numvar)
    for i,xyz_i := range(xyz) { hwd[i] = math.Exp(xyz_i) }

    return hwd
}
 
func main() {
    Aw    := 200.0
    Af    := 50.0
    alpha := 2.0
    beta  := 10.0
    gamma := 2.0
    delta := 10.0
    
    hwd := max_volume_box(Aw, Af, alpha, beta, gamma, delta)

    fmt.Printf("h=%.4f w=%.4f d=%.4f\n", hwd[0], hwd[1], hwd[2])
}
