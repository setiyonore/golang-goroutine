package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu: ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread: ", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine: ", totalGoRoutine)
}
