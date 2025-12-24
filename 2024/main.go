package main

import (
	"fmt"
	"sync"
	"time"
)

type SolutionFuncSignature = func() string

func main() {
	wg := new(sync.WaitGroup)
	traceAndRecord(wg, "D1C1", D1P1)
	wg.Wait()
}

func traceAndRecord(wg *sync.WaitGroup, name string, f SolutionFuncSignature) {
	wg.Add(1)

	start := time.Now()
	go func() {
		defer wg.Done()
		fmt.Printf("Name: %s Answer %v: Execution Time: %v\n", name, f(), time.Since(start))
	}()
}
