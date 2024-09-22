package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(wgP *sync.WaitGroup) chan int {
	ch := make(chan int)

	go func() {
		defer (*wgP).Done()
		fmt.Println("Worker has started.")
		time.Sleep(3 * time.Second)
		ch <- 42
		fmt.Println("Worker has finished.")
	}()

	return ch
}

func main() {
	timeStart := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(2)

	fmt.Printf("GOMAXPROCS = %d, while the number of CPUs is: %d\n", runtime.GOMAXPROCS(0), runtime.NumCPU()) // Checking if there is a cpu limitation that messes up the concurrency.

	ch1 := worker(&wg)
	ch2 := worker(&wg)

	_, _ = <-ch1, <-ch2

	wg.Wait()

	fmt.Printf("The number of goroutines running: %d\n", runtime.NumGoroutine())

	println(int(time.Since(timeStart).Seconds()))
}
