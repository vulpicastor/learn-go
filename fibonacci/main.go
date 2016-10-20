package main

import (
	"fmt"
)

func FibonacciIter(n int) int {
	a, b := 1, 0
	for i := 0; i < n; i++ {
		a, b = a+b, a
	}
	return b
}

func FibonacciRecursiveIter(a, b, count int) (fibonacci int) {
	if count == 0 {
		fibonacci = b
	} else {
		fibonacci = FibonacciRecursiveIter(a+b, a, count-1)
	}
	return
}

func FibonacciRecursive(n int) int {
	return FibonacciRecursiveIter(1, 0, n)
}

var fibMemoCache [100]int

func FibonacciMemo(n int) (fibonacci int) {
	switch n {
	case 0:
		fibonacci = 0
	case 1:
		fibonacci = 1
	default:
		if fibMemoCache[n] == 0 {
			fibonacci = FibonacciMemo(n-1) + FibonacciMemo(n-2)
			fibMemoCache[n] = fibonacci
		} else {
			fibonacci = fibMemoCache[n]
		}
	}
	return
}

// Inspired by SICP Exercise 1.19
func FibonacciSICP(n int) int {
	var a, b, s, t, count int = 1, 0, 1, 0, n
	for count > 0 {
		if count%2 == 0 {
			s, t = s*s+2*s*t, s*s+t*t
			count /= 2
		} else {
			a, b = s*a+t*a+s*b, s*a+t*b
			count--
		}
	}
	return b
}

func FibonacciSICPIter(a, b, s, t, count int) (fibonacci int) {
	if count == 0 {
		fibonacci = b
	} else {
		if count%2 == 0 {
			s, t = s*s+2*s*t, s*s+t*t
			fibonacci = FibonacciSICPIter(a, b, s, t, count/2)
		} else {
			a, b = s*a+t*a+s*b, s*a+t*b
			fibonacci = FibonacciSICPIter(a, b, s, t, count-1)
		}
	}
	return
}

// Recursive version of FibonacciSICP
func FibonacciSICP2(n int) int {
	return FibonacciSICPIter(1, 0, 1, 0, n)
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(FibonacciSICP(i))
	}
}
