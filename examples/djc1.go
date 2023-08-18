//
//   Copyright: Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//   File:      djc1.go
//
//   Purpose: Demonstrates how to solve the problem with two disjunctions:
//
//      minimize    2x0 + x1 + 3x2 + x3
//      subject to   x0 + x1 + x2 + x3 >= -10
//                  (x0-2x1<=-1 and x2=x3=0) or (x2-3x3<=-2 and x1=x2=0)
//                  x0=2.5 or x1=2.5 or x2=2.5 or x3=2.5
//
package main
import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
)

const inf = 0.0 // Infinity for symbolic purposes

func main() {
    task,err := mosek.NewTask(); if err != nil { panic(err) }
    defer task.Delete()

    task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })

    // Append free variables
    const numvar = 4
    if err := task.AppendVars(numvar); err != nil { panic(err) }
    if err := task.PutVarBoundSliceConst(0, numvar, mosek.MSK_BK_FR, -inf, inf); err != nil { panic(err) }

    // The linear part: the linear constraint
    if err := task.AppendCons(1); err != nil { panic(err) }
    if err := task.PutARow(0, []int32{0, 1, 2, 3}, []float64{1, 1, 1, 1}); err != nil { panic(err) }
    if err := task.PutConBound(0, mosek.MSK_BK_LO, -10.0, -10.0); err != nil { panic(err) }

    // The linear part: objective
    if err := task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MINIMIZE); err != nil { panic(err) }
    if err := task.PutCList([]int32{0, 1, 2, 3}, []float64{2, 1, 3, 1}); err != nil { panic(err) }

    // Fill in the affine expression storage F, g
    var numafe int64 = 10
    if err := task.AppendAfes(numafe); err != nil { panic(err) }

    fafeidx := []int64{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fvaridx := []int32{0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3}
    fval    := []float64{1.0, -2.0, 1.0, -3.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}
    g       := []float64{1.0, 2.0, 0.0, 0.0, 0.0, 0.0, -2.5, -2.5, -2.5, -2.5}

    if err := task.PutAfeFEntryList(fafeidx, fvaridx, fval); err != nil { panic(err) }
    if err := task.PutAfeGSlice(0, numafe, g); err != nil { panic(err) }

    // Create domains
    zero1,_   := task.AppendRzeroDomain(1)
    zero2,_   := task.AppendRzeroDomain(2)
    rminus1,_ := task.AppendRminusDomain(1)

    // Append disjunctive constraints
    var numdjc int64 = 2;
    if err := task.AppendDjcs(numdjc); err != nil { panic(err) }

    // First disjunctive constraint
    task.PutDjc(0,                                       // DJC index
                []int64{rminus1, zero2, rminus1, zero2}, // Domains     (domidxlist)
                []int64{0, 4, 5, 1, 2, 3},               // AFE indices (afeidxlist)
                nil,                                     // Unused
                []int64{2, 2} )                          // Term sizes  (termsizelist)

    // Second disjunctive constraint
    task.PutDjc(1,                                       // DJC index
                []int64{zero1, zero1, zero1, zero1},     // Domains     (domidxlist)
                []int64{6, 7, 8, 9},                     // AFE indices (afeidxlist)
                nil,                                     // Unused
                []int64{1, 1, 1, 1} )                    // Term sizes  (termidxlist)

    // Useful for debugging
    task.WriteData("djc.ptf")                            // Write file in human-readable format

    // Solve the problem
    if _,err := task.Optimize(); err != nil { panic(err) }

    // Print a summary containing information
    // about the solution for debugging purposes
    if err := task.SolutionSummary(mosek.MSK_STREAM_MSG); err != nil { panic(err) }

    // Get status information about the solution
    solsta,err := task.GetSolSta(mosek.MSK_SOL_ITG); if err != nil { panic(err) }

    switch solsta {
      case mosek.MSK_SOL_STA_INTEGER_OPTIMAL:
          xx,err := task.GetXx(mosek.MSK_SOL_ITG); if err != nil { panic(err) }

          fmt.Printf("Optimal primal solution: %v\n",xx)
        break;
      default:
        fmt.Println("Another solution status")
        break;
    }
}

