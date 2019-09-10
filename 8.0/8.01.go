package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(`Main Go Routine`)
	go func() {
		fmt.Println(`New Go Routine 1`)
		wg.Done()
	}()
	go func() {
		fmt.Println(`New Go Routine 2`)
		wg.Done()
	}()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}
