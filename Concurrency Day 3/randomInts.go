package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(chnl chan int) {
	for {
		rand.Seed(time.Now().UnixNano())
		chnl <- rand.Intn(100)
		time.Sleep(time.Second)
	}
}

func main() {
	chnl := make(chan int, 3)

	go randomGenerator(chnl)
	fmt.Println("15 random generated intergers are: ")

	for i := 0; i < 15; i++ {
		fmt.Println(<-chnl)
	}
}
