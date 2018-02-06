package main

import (
	"fmt"
	"sort"
)

func main() {
	//taskOne()
	taskMany()
}

func taskOne()  {
	c := make(chan int)  // 分配一个信道
	list := []int{1, 3, 5, 7, 2}

	// 在Go程中启动排序。当它完成后，在信道上发送信号。
	go func() {
		sort.Ints(list)
		c <- 1  // 发送信号，什么值无所谓。
	}()

	doSomethingForAWhile()
	<-c   // 等待排序结束，丢弃发来的值。

	fmt.Println(list)
}

func doSomethingForAWhile(){
	fmt.Println("ok")
}

func taskMany(){

	 list := []int{1, 3, 5, 7, 2}

	 Serve(createTask(list))

}


func handle(r int) {
	MaxOutstanding := 10
	sem := make(chan int, MaxOutstanding)

	// 同步处理任务
	sem <- 1 // 等待活动队列清空。
	process(r)  // 可能需要很长时间。
	<-sem    // 完成；使下一个请求可以运行。
}

func Serve(queue <- chan int) {
	 // 分发任务到handle
	 for v := range queue{
		 go handle(v)  // 无需等待 handle 结束。
	 }
}

// 处理任务进程
func process(r int){
		fmt.Println( "handle task: ",r)
}

func createTask(r []int) <-chan int{
	out:= make(chan int)

	 go func() {
		 for _,v:= range r{
			 out <- v
		 }
		 close(out)
	 }()

	 return out
}
