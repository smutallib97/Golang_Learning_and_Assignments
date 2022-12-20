package main

import "fmt"

func main() {
	var n int

	fmt.Printf("Enter a Number: ")
	fmt.Scanf("%d", &n)
	// compute 1/1 + 1/2 + 1/3 + ... + 1/n
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	// print the nth harmonic number
	fmt.Println(sum)
}
