package main
/*
   Purpose: To demonstrate how to solve a quadratic optimization
              problem using the MOSEK API.
 */


import (
        "mosek"
        "fmt"
        "math"
        "os"
)

func main() {
        const numcon int32 = 1 /* Number of constraints.             */
        const numvar int32 = 3 /* Number of variables.               */


        c := []float64 { 0.0, -1.0, 0.0 }

        bkc := []int32   { mosek.BK_LO }
        blc := []float64 { 1.0 }
        buc := []float64 { math.Inf(+1) }

        bkx := []int32   { mosek.BK_LO,  mosek.BK_LO,  mosek.BK_LO }
        blx := []float64 { 0.0,          0.0,          0.0}
        bux := []float64 { math.Inf(+1), math.Inf(+1), math.Inf(+1) }

        aptrb := []int32   {0,   1,   2}
        aptre := []int32   {1,   2,   3}
        asub  := []int32   {0,   0,   0}
        aval  := []float64 {1.0, 1.0, 1.0};

        qsubi := []int32   { 0, 1, 2, 2 }
        qsubj := []int32   { 0, 1, 0, 2 }
        qval  := []float64 { 2.0, 0.2, -1.0, 2.0 }

        var r int32

        /* Create the mosek environment. */
        env,r := mosek.MakeEnv()
        if r != 0 { os.Exit(1) }
        defer env.DeleteEnv()
        task,r := env.MakeTask()
        if r != 0 { os.Exit(1) }

        /* Directs the log task stream to the 'printstr' function. */
        task.PutStreamFunc(mosek.STREAM_LOG,func(msg string) { fmt.Print(msg) })

        /* Append 'numcon' empty constraints. The constraints will initially have no bounds. */
        task.AppendCons(numcon)

        /* Append 'numvar' variables. The variables will initially be fixed at zero (x=0). */
        task.AppendVars(numvar)

        /* Optionally add a constant term to the objective. */
        task.PutCfix(0.0)

        for j := int32(0); j < numvar; j++ {
                /* Set the linear term c_j in the objective.*/
                task.PutCJ(j,c[j])

                /* Set the bounds on variable j. blx[j] <= x_j <= bux[j] */
                task.PutVarBound( j, bkx[j], blx[j], bux[j])

                /* Input column j of A */
                task.PutACol(j, asub[aptrb[j]:aptre[j]], aval[aptrb[j]:aptre[j]])
        }

        /* Set the bounds on constraints.
           for i=1, ...,NUMCON : blc[i] <= constraint i <= buc[i] */
        for i := int32(0) ; i < numcon ; i++ {
                task.PutConBound( i, bkc[i], blc[i], buc[i])
        }

        /* Input the Q for the objective. */
        task.PutQObj(qsubi,qsubj,qval)

        /* Run optimizer */
        trmcode := task.Optimize()

        /* Print a summary containing information
           about the solution for debugging purposes*/
        task.SolutionSummary(mosek.STREAM_MSG)

        solsta := task.GetSolSta (mosek.SOL_ITR)
        if task.GetRes() != mosek.RES_OK { os.Exit(1) }

        switch solsta {
        case mosek.SOL_STA_OPTIMAL:
                xx := task.GetXx(mosek.SOL_BAS, nil)

                fmt.Println("Optimal primal solution")
                fmt.Println("  x = ",xx)

        case mosek.SOL_STA_DUAL_INFEAS_CER: fallthrough
        case mosek.SOL_STA_PRIM_INFEAS_CER:
                fmt.Println("Primal or dual infeasibility certificate found.")
        case mosek.SOL_STA_UNKNOWN:
                /* If the solutions status is unknown, print the
                /* termination code indicating why the optimizer
                /* terminated prematurely. */

                symname,_,_ := mosek.GetCodeDesc(trmcode)

                fmt.Println("The solution status is unknown.")
                fmt.Printf("The optimizer terminitated with code: %s\n",symname)
        default:
                fmt.Println("Other solution status.")
        }
 } /* main */
