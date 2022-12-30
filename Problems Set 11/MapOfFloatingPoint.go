package main

import (
	"fmt"
)

func Imap2(floatNumbers []float64, myFunc func(float64) float64) []float64 {
	floatMap := make([]float64, len(floatNumbers))
	for i, floatNum := range floatNumbers {
		floatMap[i] = myFunc(floatNum)
	}
	return floatMap
}

func main() {
	floatNumberSlice := []float64{0.6, 0.3, 0.84, 0.04}

	fmt.Println("Input float numbers slice is: ", floatNumberSlice)

	valueAsPercentage := Imap2(floatNumberSlice, func(value float64) float64 {
		return value * 100
	})
	fmt.Println("Value as percentage slice is: ", valueAsPercentage)

	halfOfValue := Imap2(floatNumberSlice, func(value float64) float64 {
		return value / 2
	})
	fmt.Println("Value as percentage slice is: ", halfOfValue)

}
