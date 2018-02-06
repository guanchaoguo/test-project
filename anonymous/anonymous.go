package main

import(
	"fmt"
)

func main(){
	result1 := haha1()()
	fmt.Println(result1)

	result2 := haha2()()
	fmt.Println(result2)

	result3 := haha3()()
	fmt.Println(result3)
}

//返回匿名局部函数
func haha1() func() string {
	return func() string {
		return "000"
	}
}

//返回具名局部函数
func haha2() func() string {
	ha := func() string {
		return "000"
	}
	return ha
}

//返回具名包级函数
func haha3() func() string {
	return ha
}

func ha() string {
	return "000"
}