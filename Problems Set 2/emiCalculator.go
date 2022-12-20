package main

import (
	"fmt"
	"math"
)

func emi_calculator(p, y, r float64) float64 {
	var emi float64

	r = r / (12 * 100) // one month interest
	y = 12 * y         // one month period
	emi = (p * r * math.Pow(1+r, y)) / (math.Pow(1+r, y) - 1)

	return emi
}

func main() {
	var principal, rate, years, emi float64
	fmt.Printf("Enter Pricipal Amount, Number of Years and Rate of Interest: \n")
	//fmt.Scanf("%f \t , %f \t , %f \t", &principal, &years, &rate)
	fmt.Scanf("%f \n", &principal)
	fmt.Scanf("%f \n", &years)
	fmt.Scanf("%f \n", &rate)
	emi = emi_calculator(principal, years, rate)
	fmt.Printf("\nMonthly EMI is= %f\n", emi)

}
