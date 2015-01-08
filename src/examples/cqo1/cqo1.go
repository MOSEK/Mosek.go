package main
/* 
   Purpose:   To demonstrate how to solve a small conic quadratic
              optimization problem using the MOSEK API.
 */

import "mosek"
import "fmt"
import "os"
import "math"


func main() {
	const numvar  int32 = 6
	const numcon  int32 = 1
	const numcone int32 = 2	
      
	bkc := []int32   { mosek.BK_FX }
	blc := []float64 { 1.0 }
	buc := []float64 { 1.0 }
  
	bkx := []int32 { 
		mosek.BK_LO,
		mosek.BK_LO,
		mosek.BK_LO,
		mosek.BK_FR,
		mosek.BK_FR,
		mosek.BK_FR }

	blx := []float64 { 
		0.0,
		0.0,
		0.0,
		math.Inf(-1),
		math.Inf(-1),
		math.Inf(-1) }
	bux := []float64 { 
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1) }
  
	c        := []float64 { 0.0, 0.0, 0.0, 1.0, 1.0, 1.0 }

	aptrb    := []int32   { 0, 1, 2, 3, 3, 3 }
	aptre    := []int32   { 1, 2, 3, 3, 3, 3 }
	asub     := []int32   {   0,   0,   0 }
	aval     := []float64 { 1.0, 1.0, 2.0 }

	conetype := []int32   { mosek.CT_QUAD, mosek.CT_RQUAD }
	coneptrb := []int32   { 0, 3 }
	coneptre := []int32   { 3, 6 }
	conesub := []int32 { 
		3, 0, 1,
		4, 5, 2 }
	
	var r int32
	
        env,r := mosek.MakeEnv()
        if r != 0 { os.Exit(1) } 
        defer env.DeleteEnv()
        task,r := env.MakeTask()
        if r != 0 { os.Exit(1) }


	/* Create the mosek environment. */
	task.PutStreamFunc(mosek.STREAM_LOG,func(msg string) { fmt.Print(msg) })

	/* Append 'numcon' empty constraints. The constraints will initially have no bounds. */	
        if mosek.RES_OK != task.AppendCons(numcon) { os.Exit(1) }

	/* Append 'numvar' variables. The variables will initially be fixed at zero (x=0). */
        if mosek.RES_OK != task.AppendVars(numvar) { os.Exit(1) }

	for j := int32(0) ; j < numvar ; j++ {
		/* Set the linear term c_j in the objective.*/  
		if mosek.RES_OK != task.PutCJ(j,c[j]) { os.Exit(1) }
		
		/* Set the bounds on variable j. blx[j] <= x_j <= bux[j] */
		if mosek.RES_OK !=  task.PutVarBound( j, bkx[j], blx[j], bux[j]) { os.Exit(1) }

		/* Input column j of A */   
		if mosek.RES_OK != task.PutACol(j, asub[aptrb[j]:aptre[j]], aval[aptrb[j]:aptre[j]]) { os.Exit(1) }
	}

	/* Set the bounds on constraints. for i=1, ...,numcon : blc[i] <= constraint i <= buc[i] */
	for i := int32(0) ; i < numcon ; i++ {
		if mosek.RES_OK != task.PutConBound( i, bkc[i], blc[i], buc[i])  { os.Exit(1) }
	}
	/* Append the cones. */
	for i := int32(0) ; i < numcone ; i++ {
		if mosek.RES_OK != task.AppendCone(conetype[i], 0.0, conesub[coneptrb[i]:coneptre[i]]) { os.Exit(1) }
	}
	

	/* Run optimizer */
	var trmcode int32
	trmcode,r = task.Optimize()
	if r != mosek.RES_OK { os.Exit(1) }
	

	/* Print a summary containing information about the solution
	/* for debugging purposes. */
	task.SolutionSummary(mosek.STREAM_LOG)
     
	var solsta int32

	solsta,r = task.GetSolSta (mosek.SOL_ITR)
	if mosek.RES_OK != r { os.Exit(1) }

        switch solsta {
	case mosek.SOL_STA_OPTIMAL: fallthrough
	case mosek.SOL_STA_NEAR_OPTIMAL:
		xx,_ := task.GetXx(mosek.SOL_ITR, nil)
        
		fmt.Println("Optimal primal solution")
		fmt.Println("  x = ",xx)

	case mosek.SOL_STA_DUAL_INFEAS_CER: fallthrough
	case mosek.SOL_STA_PRIM_INFEAS_CER: fallthrough
	case mosek.SOL_STA_NEAR_DUAL_INFEAS_CER: fallthrough
	case mosek.SOL_STA_NEAR_PRIM_INFEAS_CER:  
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
