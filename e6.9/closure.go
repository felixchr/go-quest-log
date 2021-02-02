package main

func main() {
	var f = fib()
	for i := 1; i < 10; i++ {
		println(i, f(i))
	}
}

func fib() func(num int) int {
	a, b := 1, 0
	return func(num int) int {
		if num > 1 {
			a, b = b, a+b
		}
		return b
	}
}
