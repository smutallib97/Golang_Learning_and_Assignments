package main

import "fmt"

var factVal = 1
var i int = 1
var n int

func facto(n int) int {
	if n == 0 {
		return 1
	} else {
		for i := 1; i <= n; i++ {
			factVal *= i
		}
	}
	return factVal
}

func main() {
	fmt.Printf("Enter a Number: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Factorial of %d = %d", n, facto(n))

}
