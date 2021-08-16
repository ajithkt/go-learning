package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println("Main Go Routine")
	wg.Add(2)
	go func() {
		fmt.Println(" Second Go Routine")
		fmt.Println(runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {
		fmt.Println("Third Go ROutine")
		fmt.Println(runtime.NumGoroutine())
		wg.Done()

	}()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}
