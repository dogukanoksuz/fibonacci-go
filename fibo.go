package main

import "fmt"

func fibo() func() int {
	first, last := 1, 0
	return func () int {
		first, last = last, first + last
		return first
	}
}

func main() {
	fibonacci := fibo()
	for i := 0; i < 20; i++ {
		fmt.Println(fibonacci())
	}
}
