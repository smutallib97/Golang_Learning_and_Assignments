package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Logging examples: ")
	log.Println("We are logging in Golang!")
	log.Panic("A panic call ...!")
	log.Fatal("Exception occured!")
	log.Println("Are we reaching here?")
}
