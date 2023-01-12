/*
Write two goroutines which have a race condition when executed concurrently

*/

package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup

func increment(s string) {
	for i := 0; i < 10; i++ {
		x := count
		fmt.Println(s, "Incrementing ", x)
		x++
		count = x
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go increment("First")
	go increment("Second")
	wg.Wait()
	fmt.Println("Final Count: ", count)
}
