package main

import "fmt"

func main() {
	var sampleNum int = 100

	fmt.Println(sampleNum)

	var test int
	fmt.Println(test)

	var ptr2Int *int

	ptr2Int = &sampleNum // address of sampleNum is stored in ptr2Int

	test = *ptr2Int. // value at the address ptr2Int is stored at test

				fmt.Println(test)
}
