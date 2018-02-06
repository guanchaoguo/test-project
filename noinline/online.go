package main

import (
	"fmt"
	"runtime"
	"time"
)

//go:noinline
func addTest(a, b int) int {
	return a + b
}

func call() {
	addTest(1, 3)
}

func main() {
	cpuNu := runtime.NumCPU()

	for i := 0; i < cpuNu; i++ {
		go func() {
			for {
				// 改为 addTest 依然不会打印
				addTest(1, 3)
				//call()
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Not print...why???")
}