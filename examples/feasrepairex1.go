/*
   Copyright : Copyright (c) MOSEK ApS, Denmark. All rights reserved.

   File :      feasrepairex1.go

   Purpose :   To demonstrate how to use the MSK_relaxprimal function to
               locate the cause of an infeasibility.

   Syntax :     On command line

                   java  feasrepairex1.feasrepairex1 feasrepair.lp

                feasrepair.lp is located in mosek\<version>\tools\examples.
*/
package main

import (
    mosek "github.com/mosek/mosek.go"
    "os"
    "fmt"
)

func main() {
    var filename string
    if len(os.Args) < 2 {
        filename = "examples/feasrepair.lp"
    } else {
        filename = os.Args[1]
    }

    task,err := mosek.NewTask(); if err != nil { panic(err) }
    defer task.Delete()
    
    if err := task.PutStreamFunc(mosek.MSK_STREAM_LOG,func(msg string) { fmt.Print(msg) }); err != nil { panic(err) }
    if err := task.ReadData(filename); err != nil { panic(err) }

    if err := task.PutIntParam(mosek.MSK_IPAR_LOG_FEAS_REPAIR, 3); err != nil { panic(err) }

    if err := task.PrimalRepair(nil,nil,nil,nil); err != nil { panic(err) }

    sum_viol,err := task.GetDouInf(mosek.MSK_DINF_PRIMAL_REPAIR_PENALTY_OBJ)

    fmt.Printf("Minimized sum of violations = %g", sum_viol)

    if _,err := task.Optimize(); err != nil { panic(err) }

    if err := task.SolutionSummary(mosek.MSK_STREAM_MSG); err != nil { panic(err) }
}

