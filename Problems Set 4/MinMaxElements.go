package main

import "fmt"

func maxAndNextHighestNum(arr []int, n int) {
	maxNum := arr[0]
	var nextHighestNum int

	for i := 0; i < n; i++ {
		if nextHighestNum < arr[i] {
			maxNum = nextHighestNum
			nextHighestNum = arr[i]
		} else if maxNum < arr[i] {
			maxNum = arr[i]
		}
	}
	fmt.Print("Max valued element from array is: ", maxNum)
	fmt.Print("Next Highest valued element array is:", nextHighestNum)
}

func minAndNextLeastNum(arr []int, n int) {
	minNum := arr[0]
	var nextLeastNum int

	for i := 0; i < n; i++ {
		if nextLeastNum > arr[i] {
			minNum = nextLeastNum
			nextLeastNum = arr[i]
		} else if minNum > arr[i] {
			minNum = arr[i]
		}
	}
	fmt.Print("Min valued element from array is: ", minNum)
	fmt.Print("Next Minimum valued element array is:", nextLeastNum)
}

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
	//fmt.Println(arr)
	maxAndNextHighestNum(arr, n)
	minAndNextLeastNum(arr, n)
}
