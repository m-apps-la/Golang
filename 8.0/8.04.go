package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	inc := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i += 1 {
		go func() {
			m.Lock()
			x := inc
			x += 1
			inc = x
			fmt.Println(inc)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(inc)
}
