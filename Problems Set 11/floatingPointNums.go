package main

import (
	"fmt"
	"math"
)

func floatNumbersFilter(floatNumbers []float64, funcName func(float64) bool) []float64 {
	filteredFloatNumbers := []float64{}
	for _, floatNum := range floatNumbers {
		if funcName(floatNum) {
			filteredFloatNumbers = append(filteredFloatNumbers, floatNum)
		}
	}
	return filteredFloatNumbers
}

func main() {
	floatNumbers := []float64{5.12, 15.0, 3.39, 8.0, 9.0, 14.0, 6.63}
	fmt.Println("Float Numbers Slice is : ", floatNumbers)

	notDigitAfterDecimal := floatNumbersFilter(floatNumbers, func(floatNums float64) bool {
		return math.Mod(floatNums, 1.0) == 0.0
	})
	fmt.Println("Not Have Any Digit After Decimal", notDigitAfterDecimal)

	divisibleBy3 := floatNumbersFilter(floatNumbers, func(floatNums float64) bool {
		return math.Mod(floatNums, 3.0) == 0.0
	})
	fmt.Println("Devisible by 3", divisibleBy3)

}
