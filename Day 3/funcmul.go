package main

import "fmt"

// function definition
func mulFloatNumbers(numb1, numb2 float32) float32 {
	product := numb1 * numb2
	return product
	//return numb1 * numb2
}

func main() {
	// function call
	result := mulFloatNumbers(1.5, 5.0)
	fmt.Println("1.5 X 5.0 = ", result)
	//fmt.Println("1.5 X 5.0 = ", mulFloatNumbers(1.5, 5.0))
}
