package lib

//斐波那契数列
//求出第n个数的值
func Fibonacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}