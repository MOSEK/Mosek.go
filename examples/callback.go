/*
   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.

   File :      callback.go

   Purpose :   To demonstrate how to use the progress
               callback.

               Use this script as follows:

                 callback psim  25fv47.mps
                 callback dsim  25fv47.mps
                 callback intpnt 25fv47.mps

               The first argument tells which optimizer to use
               i.e. psim is primal simplex, dsim is dual simplex
               and intpnt is interior-point.
*/
package main

import (
	mosek "github.com/mosek/mosek.go"
	"fmt"
        "os"
)

func makeUserCallback(maxtime float64) func(mosek.Callbackcode,[]float64,[]int32,[]int64)bool {
    return func(caller  mosek.Callbackcode,
                douinf  []float64,
                intinf  []int32,
                lintinf []int64) (stop bool) {

        var opttime float64
        var itrn    int32
        var pobj    float64
        var dobj    float64
        var stime   float64

        switch caller {
          case mosek.MSK_CALLBACK_BEGIN_INTPNT:
            fmt.Printf("Starting interior-point optimizer\n")
            break
          case mosek.MSK_CALLBACK_INTPNT:
            itrn    = intinf[mosek.MSK_IINF_INTPNT_ITER      ]
            pobj    = douinf[mosek.MSK_DINF_INTPNT_PRIMAL_OBJ]
            dobj    = douinf[mosek.MSK_DINF_INTPNT_DUAL_OBJ  ]
            stime   = douinf[mosek.MSK_DINF_INTPNT_TIME      ]
            opttime = douinf[mosek.MSK_DINF_OPTIMIZER_TIME   ]

            fmt.Printf("Iterations: %-3d\n", itrn)
            fmt.Printf("  Time: %6.2f(%.2f) ", opttime, stime)
            fmt.Printf("  Primal obj.: %-18.6e  Dual obj.: %-18.6e\n", pobj, dobj)
            break
          case mosek.MSK_CALLBACK_END_INTPNT:
            fmt.Printf("Interior-point optimizer finished.\n")
            break
          case mosek.MSK_CALLBACK_BEGIN_PRIMAL_SIMPLEX:
            fmt.Printf("Primal simplex optimizer started.\n")
            break
          case mosek.MSK_CALLBACK_UPDATE_PRIMAL_SIMPLEX:
            itrn    = intinf[mosek.MSK_IINF_SIM_PRIMAL_ITER  ]
            pobj    = douinf[mosek.MSK_DINF_SIM_OBJ          ]
            stime   = douinf[mosek.MSK_DINF_SIM_TIME         ]
            opttime = douinf[mosek.MSK_DINF_OPTIMIZER_TIME   ]

            fmt.Printf("Iterations: %-3d\n", itrn)
            fmt.Printf("  Elapsed time: %6.2f(%.2f\n", opttime, stime)
            fmt.Printf("  Obj.: %-18.6e", pobj )
            break
          case mosek.MSK_CALLBACK_END_PRIMAL_SIMPLEX:
            fmt.Printf("Primal simplex optimizer finished.\n")
            break
          case mosek.MSK_CALLBACK_BEGIN_DUAL_SIMPLEX:
            fmt.Printf("Dual simplex optimizer started.\n")
            break
          case mosek.MSK_CALLBACK_UPDATE_DUAL_SIMPLEX:
            itrn    = intinf[mosek.MSK_IINF_SIM_DUAL_ITER    ]
            pobj    = douinf[mosek.MSK_DINF_SIM_OBJ          ]
            stime   = douinf[mosek.MSK_DINF_SIM_TIME         ]
            opttime = douinf[mosek.MSK_DINF_OPTIMIZER_TIME   ]
            fmt.Printf("Iterations: %-3d\n", itrn)
            fmt.Printf("  Elapsed time: %6.2f(%.2f)\n", opttime, stime)
            fmt.Printf("  Obj.: %-18.6e\n", pobj)
            break
          case mosek.MSK_CALLBACK_END_DUAL_SIMPLEX:
            fmt.Printf("Dual simplex optimizer finished.\n")
            break
          case mosek.MSK_CALLBACK_BEGIN_BI:
            fmt.Printf("Basis identification started.\n")
            break
          case mosek.MSK_CALLBACK_END_BI:
            fmt.Printf("Basis identification finished.\n")
            break
          default:
        }

        stop = opttime >= maxtime
        return
    }
}

func main() {
    if len(os.Args) < 3 {
      fmt.Println("Usage: ( psim | dsim | intpnt ) filename")
      return
    }
    slvr     := os.Args[1]
    filename := os.Args[2]
    fmt.Printf("filename = %s",filename)

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    defer task.Delete()
  
    if err := task.ReadData(filename); err != nil { panic(err) }

    switch slvr {
        case "psim":   task.PutIntParam(mosek.MSK_IPAR_OPTIMIZER, int32(mosek.MSK_OPTIMIZER_PRIMAL_SIMPLEX))
        case "dsim":   task.PutIntParam(mosek.MSK_IPAR_OPTIMIZER, int32(mosek.MSK_OPTIMIZER_DUAL_SIMPLEX))
        case "intpnt": task.PutIntParam(mosek.MSK_IPAR_OPTIMIZER, int32(mosek.MSK_OPTIMIZER_INTPNT))
    }

    // Turn all MOSEK logging off (note that errors and other messages
    // are still sent through the log stream)

    maxtime := 0.05
    task.PutInfoCallbackFunc(makeUserCallback(maxtime))
    task.Optimize()
    task.PutIntParam(mosek.MSK_IPAR_LOG, 1)
    task.SolutionSummary(mosek.MSK_STREAM_LOG)
}
