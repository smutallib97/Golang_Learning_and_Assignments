package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(c chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		c <- rand.Intn(100)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan int, 3)

	go randomGenerator(c)
	fmt.Println("15 random generated intergers are: ")

	for i := 0; i < 15; i++ {
		fmt.Println(<-c)
	}
}
