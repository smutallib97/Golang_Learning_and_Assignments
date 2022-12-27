package main

import (
	"bufio"
	"fmt"
	"os"
)

// function for file printLinebyLine
func printLinebyLine(fileName string) {

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {

		fmt.Println("Reading and Priting the file line by line", sc.Text())
	}

}

func main() {
	fileName := "LearnerData.txt"
	readFile, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println("Reading the complete file")
	fmt.Println(string(readFile))

	fmt.Println("******************************************************")
	printLinebyLine(fileName)
}
