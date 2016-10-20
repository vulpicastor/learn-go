package main

import (
	"fmt"
)

var fibRecurseCache [100]int

func FibonacciRecurse(n int) (fibonacci int) {
	switch n {
	case 0:
		fibonacci = 0
	case 1:
		fibonacci = 1
	default:
		if fibRecurseCache[n] == 0 {
			fibonacci = FibonacciRecurse(n-1) + FibonacciRecurse(n-2)
			fibRecurseCache[n] = fibonacci
		} else {
			fibonacci = fibRecurseCache[n]
		}
	}
	return
}

func FibonacciMat() {
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(FibonacciRecurse(i))
	}
}
