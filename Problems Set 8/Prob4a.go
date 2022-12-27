package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	readContent, err := ioutil.ReadFile("Cricket_Players_Stats.csv")

	if err != nil {
		panic(err)
	}
	fmt.Println("Reading contents from Cricket_Players_Stats.csv file")
	fmt.Println(string(readContent))
}
