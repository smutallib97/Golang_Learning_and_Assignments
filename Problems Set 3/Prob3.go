package main

import "fmt"

func printFibonacii(n1, n2, n int) int {
	var nextN int = 0
	if n > 0 {
		nextN = n1 + n2
		fmt.Printf("%d ", nextN)
		n1 = n2
		n2 = nextN
		printFibonacii(n1, n2, n-1)
	}
	return 0
}

func main() {
	var n1, n2, n int

	n1 = 0
	n2 = 1

	fmt.Printf("Enter total number of terms: ")
	fmt.Scanf("%d", &n)

	fmt.Printf("Fibonacii series is : ")
	fmt.Printf("%d\t%d ", n1, n2)

	printFibonacii(n1, n2, n-2)
	fmt.Printf("\n")
}
