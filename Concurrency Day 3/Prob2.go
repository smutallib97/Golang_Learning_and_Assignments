package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {

	var nums int
	var userNum float64
	//var wg sync.WaitGroup
	sliceOfFloats := make([]float64, 0, 3)

	fmt.Println("Please enter a series of float numbers into the array.")
	fmt.Println("Enter the size of the array")
	fmt.Scan(&nums)

	for i := 0; i < nums; i++ {

		fmt.Println("Please Enter a float number to fill the array")
		_, err := fmt.Scan(&userNum)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Invalid user input")
		}

		sliceOfFloats = append(sliceOfFloats, userNum)

	}
	//var sliceSize float64
	sliceSize := nums / 4.0
	slice1 := sliceOfFloats[:sliceSize]
	slice2 := sliceOfFloats[sliceSize : 2*(sliceSize)]
	slice3 := sliceOfFloats[2*(sliceSize) : 3*(sliceSize)]
	slice4 := sliceOfFloats[3*(sliceSize):]

	fmt.Println("Here are your arrays when partitioned", slice1, slice2, slice3, slice4)

	// wg.Add(4)
	// go sortList(slice1)
	// wg.Done()
	// go sortList(slice2)
	// wg.Done()
	// go sortList(slice3)
	// wg.Done()
	// go sortList(slice4)
	// wg.Done()
	// wg.Wait()
	chnl := make(chan []float64, 4)
	go sortList(slice1, chnl)
	go sortList(slice2, chnl)
	go sortList(slice3, chnl)
	go sortList(slice4, chnl)

	// slice1 = <-chnl
	// slice2 = <-chnl
	// slice3 = <-chnl
	// slice4 = <-chnl

	newSlice := mergeAndSort(slice1, slice2, slice3, slice4)
	fmt.Println("Sorted Array is: ", newSlice)

}

func sortList(unsortedSlice []float64, chnl chan []float64) []float64 {
	sort.Float64s(unsortedSlice)
	//chnl <- unsortedSlice
	return unsortedSlice

}

func mergeAndSort(list1 []float64, list2 []float64, list3 []float64, list4 []float64) []float64 {
	newSlice := []float64{}
	newSlice = append(list1, list2...)
	newSlice = append(newSlice, list3...)
	newSlice = append(newSlice, list4...)
	sort.Float64s(newSlice)
	chnl := make(chan []float64, 4)
	chnl <- newSlice
	return newSlice

}
