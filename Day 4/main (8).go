package main

import "fmt"

func main() {
	myArray := [4]int{2, 3, 4, 5}
	sum := 0
	for i, sum := range myArray {
		fmt.Println(i, sum)
	}
}
