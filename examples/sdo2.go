/*
  Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.

  File :      sdo2.go

  Purpose :   Solves the semidefinite problem with two symmetric variables:

                 min   <C1,X1> + <C2,X2>
                 st.   <A1,X1> + <A2,X2> = b
                             (X2)_{1,2} <= k
                
                 where X1, X2 are symmetric positive semidefinite,

                 C1, C2, A1, A2 are assumed to be constant symmetric matrices,
                 and b, k are constants.
*/
package main

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
)


func main() {
    /* Input data */
    const numcon int32 = 2              /* Number of constraints. */
    const numbarvar int32 = 2
    dimbarvar := []int32{3, 4}         /* Dimension of semidefinite variables */

    /* Objective coefficients concatenated */
    Cj := []int32{ 0, 0, 1, 1, 1, 1 }   /* Which symmetric variable (j) */
    Ck := []int32{ 0, 2, 0, 1, 1, 2 }   /* Which entry (k,l)->v */
    Cl := []int32{ 0, 2, 0, 0, 1, 2 }
    Cv := []float64{ 1.0, 6.0, 1.0, -3.0, 2.0, 1.0 }

    /* Equality constraints coefficients concatenated */
    Ai := []int32{ 0, 0, 0, 0, 0, 0 }   /* Which constraint (i = 0) */
    Aj := []int32{ 0, 0, 0, 1, 1, 1 }   /* Which symmetric variable (j) */
    Ak := []int32{ 0, 2, 2, 1, 1, 3 }   /* Which entry (k,l)->v */
    Al := []int32{ 0, 0, 2, 0, 1, 3 }
    Av := []float64{ 1.0, 1.0, 2.0, 1.0, -1.0, -3.0 }

    /* The second constraint - one-term inequality */
    A2i := []int32{ 1 }                        /* Which constraint (i = 1) */
    A2j := []int32{ 1 }                        /* Which symmetric variable (j = 1) */
    A2k := []int32{ 1 }                        /* Which entry A(1,0) = A(0,1) = 0.5 */
    A2l := []int32{ 0 }
    A2v := []float64{ 0.5 }

    bkc := []mosek.Boundkey{ mosek.MSK_BK_FX,
                             mosek.MSK_BK_UP }
    blc := []float64{ 23.0, 0.0 }
    buc := []float64{ 23.0, -3.0 }

    task,err := mosek.NewTask()
    if err != nil { panic(err) }
    defer task.Delete()

    // Directs the log task stream to the 'printstr' function. 
    task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })

    /* Append numcon empty constraints.
       The constraints will initially have no bounds. */
    if err := task.AppendCons(numcon); err != nil { panic(err) }

    /* Append numbarvar semidefinite variables. */
    if err := task.AppendBarvars(dimbarvar); err != nil { panic(err) }

    /* Set objective (6 nonzeros).*/
    if err := task.PutBarcBlockTriplet(Cj, Ck, Cl, Cv); err != nil { panic(err) }

    /* Set the equality constraint (6 nonzeros).*/
    if err := task.PutBaraBlockTriplet(Ai, Aj, Ak, Al, Av); err != nil { panic(err) }

    /* Set the inequality constraint (1 nonzero).*/
    if err := task.PutBaraBlockTriplet(A2i, A2j, A2k, A2l, A2v); err != nil { panic(err) }

    /* Set constraint bounds */
    if err := task.PutConBoundSlice(0, 2, bkc, blc, buc); err != nil { panic(err) }

    /* Run optimizer */
    if _,err := task.Optimize(); err != nil { panic(err) }
    if err := task.SolutionSummary(mosek.MSK_STREAM_MSG); err != nil { panic(err) }

    solsta,err := task.GetSolSta(mosek.MSK_SOL_ITR); if err != nil { panic(err) }

    switch solsta { 
    case mosek.MSK_SOL_STA_OPTIMAL:
        /* Retrieve the soution for all symmetric variables */
        fmt.Println("Solution (lower triangular part vectorized):")
        for i := int32(0); i < numbarvar; i++ {
            barx,err := task.GetBarxJ(mosek.MSK_SOL_ITR, i); if err != nil { panic(err) }

            fmt.Printf("X %d: %v\n",i+1,barx);
        }
    case mosek.MSK_SOL_STA_DUAL_INFEAS_CER,mosek.MSK_SOL_STA_PRIM_INFEAS_CER:
        fmt.Println("Primal or dual infeasibility certificate found.")
    case mosek.MSK_SOL_STA_UNKNOWN:
        fmt.Println("The status of the solution could not be determined.")
    default:
          fmt.Println("Other solution status: ",solsta)
    }
}
