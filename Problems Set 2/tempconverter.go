package main

import "fmt"

func main() {
	var n int
	var temp_c float32
	var temp_f float32
	fmt.Printf("Enter your choice : ")
	fmt.Scanf("%d", &n)

	switch n {
	case 0:
		fmt.Printf("Enter Temperature in centigrade: \n")
		fmt.Scanf("%f", &temp_c)
		temp_f = ((9.0 / 5.0) * temp_c) + 32.0
		fmt.Printf("%f degrees Fahrenheit ", temp_f)
	case 1:
		fmt.Printf("Enter Temperature in Fahrenheit: \n")
		fmt.Scanf("%f", &temp_f)
		temp_c = ((temp_f - 32) * 5) / 9 //[(°F-32)×5]/9
		fmt.Printf("%f degrees centigrade ", temp_c)

	default:
		fmt.Printf("Invalid Input")
	}
}
