package main

import "fmt"

func average(numbers ...int) float64 {
	var add int
	count := 0
	for _, value := range numbers {
		add += value
		count++
	}
	avg := float64(add) / float64(count)
	return avg
}

func main() {
	threeIntegerNumbers := average(12, 40, 30)
	fmt.Println("Three Integer Numbers Average: ", threeIntegerNumbers)

	fiveIntegerNumbers := average(12, 40, 30, 18, 26)
	fmt.Println("Five Integer Numbers Average: ", fiveIntegerNumbers)

}
