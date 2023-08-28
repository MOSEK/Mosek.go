//
//   Copyright: Copyright (c) MOSEK ApS, Denmark. All rights reserved.
//
//   File:      parallel.java
//
//   Purpose: Demonstrates parallel optimization using optimizebatch()
//
package main

import (
    mosek "github.com/mosek/mosek.go"
    "os"
    "fmt"
)

// Example of how to use env.optimizebatch(). 
// Optimizes tasks whose names were read from command line.
func main() {
    n := len(os.Args)-1
    
    tasks := make([]*mosek.Task,n)

    // Size of thread pool available for all tasks
    const threadpoolsize = 6 

    // Create an example list of tasks to optimize
    for i,p := range(os.Args[1:]) {
        t,err := mosek.NewTask(); if err != nil { panic(err) }
        tasks[i] = t
        if err := tasks[i].ReadData(p); err != nil { panic(err) }
        // We can set the number of threads for each task
        if err := tasks[i].PutIntParam(mosek.MSK_IPAR_NUM_THREADS, 2); err != nil { panic(err) }
    }

    // Optimize all the given tasks in parallel
    trm,res,err := mosek.OptimizeBatch(false,          // No race
                        -1.0,           // No time limit
                        threadpoolsize,
                        tasks);         // Array of tasks to optimize
    if err != nil { panic(err) }

    for i,r := range(res) {
        pobj,err := tasks[i].GetDouInf(mosek.MSK_DINF_INTPNT_PRIMAL_OBJ); if err != nil { panic(err) }
        opttime,err := tasks[i].GetDouInf(mosek.MSK_DINF_OPTIMIZER_TIME); if err != nil { panic(err) }

      fmt.Printf("Task  %d  res %s   trm %s   obj_val  %f  time %f\n", i, r, trm[i], pobj, opttime)
    }
}
