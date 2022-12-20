package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var arr [5]int
	sum := 0
	// lower := 10
	// upper := 50
	fmt.Println("Random numbers between 10 to 50 are: ")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn((50 - 10 + 1) + 10)
		fmt.Println(arr[i])
		sum += arr[i]
	}
	avg := float32(sum) / 5.0
	fmt.Println("Average is : ", avg)
}
