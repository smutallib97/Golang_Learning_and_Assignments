package main

import (
	"fmt"
)

func main() {
	inpSlice := make([]uint8, 3, 5)
	fmt.Printf("Length of inpSlice is %d and Capacity is %d  \n", len(inpSlice), cap(inpSlice))

	inpSlice = append(inpSlice, 2, 4, 8)
	fmt.Println("after adding values in slice ", inpSlice)

}
