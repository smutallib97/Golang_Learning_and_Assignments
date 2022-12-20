package main

import "fmt"

func sum(n1, n2 int) int {
	//var sum int
	sum := n1 + n2
	return sum
}

func diff(n1, n2 int) int {
	//var diff int
	diff := n1 - n2
	return diff
}

func getSum_n_Diff(n1, n2 int) (int, int) {

	sum := n1 + n2
	diff := n1 - n2
	return sum, diff
}

func main() {
	var n1, n2 int
	fmt.Println("Enter Two Numbers:")
	fmt.Scanln(&n1, &n2)

	fmt.Println("Sum Function Called")
	fmt.Println(sum(n1, n2))

	fmt.Println("Diff Function Called")
	fmt.Println(diff(n1, n2))

	fmt.Println("getSum_n_Diff Function Called")
	fmt.Println(getSum_n_Diff(n1, n2))
}
