/*
   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.

   File :      solutionquality.go

   Purpose :   To demonstrate how to examine the quality of a solution.
*/
package main
import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
	"math"
        "os"
)

func main() {
    if len(os.Args) < 2 { 
      fmt.Println ("Missing argument, syntax is:");
      fmt.Println ("  solutionquality inputfile");
    } else {
        task,err := mosek.NewTask()
        if err != nil { panic(err) }
        defer task.Delete()
        
        // We assume that a problem file was given as the first command
        // line argument (received in `args')
        if err := task.ReadData(os.Args[1]); err != nil { panic(err) }

        // Solve the problem
        if _,err := task.Optimize(); err != nil { panic(err) }

        if err := task.SolutionSummary(mosek.MSK_STREAM_LOG); err != nil { panic(err) }

        solsta,err := task.GetSolSta(mosek.MSK_SOL_BAS); if err != nil { panic(err) }

        pobj,pviolcon,pviolvar,pviolbarvar,pviolcones,_,dobj,dviolcon,dviolvar,dviolbarvar,dviolcones,err := task.GetSolutionInfo(mosek.MSK_SOL_BAS)

        switch (solsta) {
        case mosek.MSK_SOL_STA_OPTIMAL:
            abs_obj_gap := math.Abs(dobj - pobj)
            rel_obj_gap := abs_obj_gap / (1.0 + math.Min(math.Abs(pobj), math.Abs(dobj)))
            max_primal_viol := math.Max(pviolcon, pviolvar)
            max_primal_viol = math.Max(max_primal_viol, pviolbarvar)
            max_primal_viol = math.Max(max_primal_viol, pviolcones)

            max_dual_viol := math.Max(dviolcon, dviolvar)
            max_dual_viol = math.Max(max_dual_viol, dviolbarvar)
            max_dual_viol = math.Max(max_dual_viol, dviolcones)

            // Assume the application needs the solution to be within
            //    1e-6 ofoptimality in an absolute sense. Another approach
            //   would be looking at the relative objective gap

            fmt.Printf("Customized solution information.\n");
            fmt.Printf("  Absolute objective gap: %.2e\n", abs_obj_gap)
            fmt.Printf("  Relative objective gap: %.2e\n", rel_obj_gap)
            fmt.Printf("  Max primal violation  : %.2e\n", max_primal_viol)
            fmt.Printf("  Max dual violation    : %.2e\n", max_dual_viol)

            accepted := true

            if rel_obj_gap > 1e-6 {
              fmt.Println("Warning: The relative objective gap is LARGE.")
              accepted = false
            }

            // We will accept a primal infeasibility of 1e-8 and
            // dual infeasibility of 1e-6. These number should chosen problem
            // dependent.
            if max_primal_viol > 1e-8 {
              fmt.Println("Warning: Primal violation is too LARGE")
              accepted = false
            }

            if max_dual_viol > 1e-6 {
              fmt.Println("Warning: Dual violation is too LARGE.")
              accepted = false;
            }

            if accepted {
                xx,err := task.GetXx(mosek.MSK_SOL_BAS); if err != nil { panic(err) }
                fmt.Println("Optimal primal solution")

                fmt.Printf("x = %v\n",xx)
            } else {
              // print etailed information about the solution
              if err := task.AnalyzeSolution(mosek.MSK_STREAM_LOG, mosek.MSK_SOL_BAS); err != nil { panic(err) }
            }
          case mosek.MSK_SOL_STA_DUAL_INFEAS_CER,mosek.MSK_SOL_STA_PRIM_INFEAS_CER:
            fmt.Println("Primal or dual infeasibility certificate found.")
          case mosek.MSK_SOL_STA_UNKNOWN:
            fmt.Println("The status of the solution is unknown.")
          default:
            fmt.Println("Other solution status")
        }
    }
}
