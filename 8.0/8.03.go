package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	inc := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i += 1 {
		go func() {
			x := inc
			runtime.Gosched()
			x += 1
			inc = x
			fmt.Println(inc)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(inc)
}
