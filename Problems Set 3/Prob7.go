package main

import (
	"fmt"
)

func findMultiples(n int) int {
	var i int
	for i = 0; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
		}

	}
	return 0
}

func main() {
	//var n int
	// 	fmt.Printf("Enter a number: ")
	// 	fmt.Scanf("%d",&n)

	//fmt.Printf("Multiples of 3 or 5 are %d: ",n, findMultiples(18))

	fmt.Println(findMultiples(18))

}
