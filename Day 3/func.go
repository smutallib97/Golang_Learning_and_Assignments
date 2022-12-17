package main

import "fmt"

// function definition
func mulFloatNumbers(numb1, numb2 float32) float32 {
	return numb1 * numb2
}

func main() {
	fmt.Println("1.5 X 5.0 = ", mulFloatNumbers(1.5, 5.0))
}
