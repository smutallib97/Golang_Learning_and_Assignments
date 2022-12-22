package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n int
	var str string
	inputSlice := make([]int, n)
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter the element : \n")
		fmt.Scan(&str)
		if str == "X" {
			break
		}
		//fmt.Println(str)

		slice, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
			continue
		}
		inputSlice = append(inputSlice, slice)

		sort.Ints(inputSlice)
		fmt.Println(inputSlice)
	}
	fmt.Println("Sorted Inputslice is : ", inputSlice)

}
