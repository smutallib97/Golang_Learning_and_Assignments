package main

import "fmt"

func facto(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * facto(n-1)
	}
}

func main() {
	var n int
	fmt.Printf("Enter a Number: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Factorial of %d = %d", n, facto(n))

}
