package main

import (
	"fmt"
	//"time"
	//"time"
	//"github.com/name5566/leaf/timer"
)

func main() {
	var (
		chaA  = make(chan int)
		chaB  = make(chan int)
		//timer = time.NewTicker(time.Second)
	)

	go func() {
		chaA <- 1
		chaB <- 2
	}()

	for {
		loop:
		select {
		case <-chaA:
			fmt.Println("1111")
			break
		case <-chaB:
			fmt.Println("2222")
		/*case time_out := <-timer.C:
			fmt.Println("nihao", time_out)
			break*/
		default:
			break loop
		}
		fmt.Println("break here----> for ")
	}

	fmt.Println("break there ----over")
}
