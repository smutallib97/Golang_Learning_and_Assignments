package main

import "fmt"

func isOdd(arr1 []int) {
	var oddSlice []int

	for i := 0; i < len(arr1); i++ {
		if arr1[i]%2 != 0 {
			oddSlice = append(oddSlice, arr1[i])
		}
	}
	fmt.Println("Odd Elements Slice is : ", oddSlice)

}

func isEven(arr1 []int) {
	var evenSlice []int

	for i := 0; i < len(arr1); i++ {
		if arr1[i]%2 == 0 {
			evenSlice = append(evenSlice, arr1[i])
		}
	}
	fmt.Println("Even Elements Slice is :", evenSlice)
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Array of 20 Interger number is:", arr)
	slice := arr[5:15]
	fmt.Println("10 elements slice from original array:", slice)

	// 	isOdd(arr)
	// 	isEven(arr)

	isOdd(slice)
	isEven(slice)
}
