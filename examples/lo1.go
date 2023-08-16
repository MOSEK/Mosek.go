package main
/*
  Purpose:   To demonstrate how to solve a small linear
             optimization problem using the MOSEK C API,
             and handle the solver result and the problem
             solution.
*/

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
	"math"
)

func main() {
        const numvar int32 = 4
        const numcon int32 = 3

        c := []float64 {3.0, 1.0, 5.0, 1.0}
        /* Below is the sparse representation of the A
           matrix stored by column. */

        aptrb := []int32 {0, 2, 5, 7}
        aptre := []int32 {2, 5, 7, 9}
        asub  := [9]int32 {
		0, 1,
		0, 1, 2,
		0, 1,
		1, 2 }
        aval  := [9]float64 {
		3.0, 2.0,
		1.0, 1.0, 2.0,
		2.0, 3.0,
		1.0, 3.0};

        /* Bounds on constraints. */
        bkc := []mosek.Boundkey { mosek.MSK_BK_FX, mosek.MSK_BK_LO, mosek.MSK_BK_UP    }
        blc := []float64{ 30.0,        15.0,        math.Inf(-1)}
        buc := []float64{ 30.0,        math.Inf(+1),25.0         }
        /* Bounds on variables. */
        bkx := []mosek.Boundkey { mosek.MSK_BK_LO,  mosek.MSK_BK_RA, mosek.MSK_BK_LO,  mosek.MSK_BK_LO }
        blx := []float64 { 0.0,          0.0,         0.0,          0.0 };
        bux := []float64 { math.Inf(+1), 10.0,        math.Inf(+1), math.Inf(+1) };

        task,err := mosek.NewTask()
        if err != nil { panic(err) }
        defer task.Delete()

	/* Directs the log task stream to the 'printstr' function. */
	task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) })

	/* Append 'numcon' empty constraints. The constraints will
	/* initially have no bounds. */
	task.AppendCons(numcon)

	/* Append 'numvar' variables. The variables will initially be
	/* fixed at zero (x=0). */
	task.AppendVars(numvar)

	for j := int32(0) ; j < numvar ; j++ {
		/* Set the linear term c_j in the objective.*/
		task.PutCJ(j,c[j])

		/* Set the bounds on variable j. blx[j] <= x_j <=
		/* bux[j] */
		task.PutVarBound(j,bkx[j], blx[j], bux[j])
		/* Input column j of A */
		task.PutACol(j, asub[aptrb[j]:aptre[j]], aval[aptrb[j]:aptre[j]])
	}

	/* Set the bounds on constraints. for i=1, ...,numcon : blc[i]
	/* <= constraint i <= buc[i] */
	for i := int32(0) ; i < numcon; i++ {
            task.PutConBound(i, bkc[i], blc[i], buc[i])
	}

	/* Maximize objective function. */
	task.PutObjSense(mosek.MSK_OBJECTIVE_SENSE_MAXIMIZE)

	/* Run optimizer */
	trmcode,err := task.Optimize()
        if err != nil { panic(err) }

	/* Print a summary containing information about the solution
	/* for debugging purposes. */
	task.SolutionSummary(mosek.MSK_STREAM_LOG)

        if solsta,err := task.GetSolSta(mosek.MSK_SOL_BAS); err != nil {
            panic(err)
        } else {
            switch solsta {
            case mosek.MSK_SOL_STA_OPTIMAL:
                if xx,err := task.GetXx(mosek.MSK_SOL_BAS); err != nil {
                    panic(err)
                } else {
                    fmt.Println("Optimal primal solution")
                    fmt.Println("  x = ",xx)
                }
            case mosek.MSK_SOL_STA_DUAL_INFEAS_CER: fallthrough
            case mosek.MSK_SOL_STA_PRIM_INFEAS_CER:
                fmt.Println("Primal or dual infeasibility certificate found.")
            case mosek.MSK_SOL_STA_UNKNOWN:
                /* If the solutions status is unknown, print the
                /* termination code indicating why the optimizer
                /* terminated prematurely. */

                symname,_,_ := mosek.GetCodeDesc(trmcode)

                fmt.Println("The solution status is unknown.")
                fmt.Printf("The optimizer terminitated with code: %s\n",symname)
            default:
                fmt.Println("Other solution status.")
            }
        }
}
