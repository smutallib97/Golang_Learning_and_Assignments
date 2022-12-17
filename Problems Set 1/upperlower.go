/*
Write a program which takes input of a place name where
you would like to visit most and displays that with place
with uppercase once and then all lower case once.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var place string

	fmt.Print("Enter Your Favourite Place Name: ")
	fmt.Scanln(&place)

	var upper_case = strings.ToUpper(place)
	fmt.Print("Your Favourite Place is: ", upper_case)
	//fmt.Print(upper_case)

	var lower_case = strings.ToLower(place)
	fmt.Print("\nYour Favourite Place is: ", lower_case)
	//fmt.Printf(lower_case)

}
