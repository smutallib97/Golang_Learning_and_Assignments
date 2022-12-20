package main

import "fmt"

func evenFibo(n int) int {
	if n < 1 {
		return n
	}
	if n == 1 {
		return 2
	}

	// calculation of
	// Fn = 4*(Fn-1) + Fn-2
	return ((4 * evenFibo(n-1)) +
		evenFibo(n-2))
}

func main() {
	var n int
	fmt.Printf("Enter number of terms: ")
	fmt.Scanf("%d", &n)

	fmt.Printf("Count of Even Fibonacci Numbers are %d", evenFibo(n))
}
