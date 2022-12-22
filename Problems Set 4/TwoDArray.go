package main

import "fmt"

func main() {
	var twoarr [5][2]int

	twoarr[0][0] = 10
	twoarr[1][0] = 12
	twoarr[2][0] = 15
	twoarr[3][0] = 19
	twoarr[4][0] = 24

	for i := 0; i < len(twoarr); i++ {
		twoarr[i][1] = twoarr[i][0] * 2
	}
	fmt.Println(twoarr)
}
