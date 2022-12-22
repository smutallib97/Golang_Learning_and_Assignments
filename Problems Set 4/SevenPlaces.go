package main

import "fmt"

func main() {
	var places [7]string

	for i := 0; i < len(places); i++ {
		fmt.Println("Enter the favourite Place names : ")
		fmt.Scanln(&places[i])
	}
	fmt.Println("Your favourite places are: ", places)
	for i := 0; i < len(places); i++ {
		count := 0
		for j := 0; j < len(places[i]); j++ {
			count++
		}
		fmt.Println("Place name : ", places[i], " its characters  count is ", count)
	}

}
