package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter the element : \n")
		fmt.Scan(&arr[i])
	}
	fmt.Println("Array Elements are: ", arr)
}
