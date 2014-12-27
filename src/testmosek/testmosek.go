package main

import (
	"mosek"
	"os"
	"fmt"
)

func main() {
        env,err := mosek.MakeEnv()
        if err != 0 { return } 
        defer env.DeleteEnv()
        task,err := env.MakeTask()
        if err != 0 { return }
        
        task.ReadData(os.Args[1])
        var numcon, numvar int
        numcon,_ = task.GetNumCon()
        numvar,_ = task.GetNumVar()
        fmt.Println("Numcon : ", numcon)
        fmt.Println("Numvar : ", numvar)

        for i := 0; i < numcon; i++ {
		var subj []int32
		var valj []float64
		subj,valj,_ = task.GetARow(i, nil, nil)

		fmt.Println("row :",subj,valj)
	}
        

        task.WriteData(os.Args[2])
}
