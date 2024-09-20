package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	How can we reduce the memory conumption of the following process?
*/

func main() {
	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)
	fmt.Printf("Before get(): Alloc = %v MiB\n", mem.Alloc/1024/1024)

	resOfRes := make([][]int, 0, 100)

	for i := 0; i < 50; i++ {
		res := getLastElem()

		runtime.GC()
		runtime.ReadMemStats(&mem)
		fmt.Printf("After getAll(): Alloc = %v MiB, slice: %v\n", mem.Alloc/1024/1024, len(resOfRes))

		resOfRes = append(resOfRes, res)
	}

	fmt.Println()
	_ = resOfRes

	for i := 0; i < 20000; i++ {
		runtime.GC()
		runtime.ReadMemStats(&mem)
		fmt.Printf("Alloc = %v MiB", mem.Alloc/1024/1024)
		time.Sleep(time.Second * 2)
	}

	fmt.Println(resOfRes)
}

func getLastElem() []int {
	sl := make([]int, 0, 100000)

	for i := 0; i < 100000; i++ {
		sl = append(sl, i)
	}

	return sl[99999:]
}
