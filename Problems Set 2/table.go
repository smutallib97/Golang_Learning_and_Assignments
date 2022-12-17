package main

import "fmt"

func main() {
	var n int
	fmt.Printf("Enter a Number between 2 to 25: ")
	fmt.Scanf("%d", &n)

	if n >= 2 && n <= 25 {
		for i := 1; i <= 10; i++ {
			fmt.Println(n, " * ", i, "=", n*i)
		}
	}
}
