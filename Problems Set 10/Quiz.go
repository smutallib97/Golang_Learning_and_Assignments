package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Questions struct {
	Question       string
	AnswerOption_A string
	AnswerOption_B string
	AnswerOption_C string
	AnswerOption_D string
	CorrectOption  string
}

var totalQuestion int

func startQuiz(options [][]string) int {
	fmt.Println("==================Quiz Started=================")
	for _, optionsIndex := range options {
		totalQuestion += 1
		data := Questions{
			Question:       optionsIndex[0],
			AnswerOption_A: optionsIndex[1],
			AnswerOption_B: optionsIndex[2],
			AnswerOption_C: optionsIndex[3],
			AnswerOption_D: optionsIndex[4],
			CorrectOption:  optionsIndex[5],
		}
		fmt.Println(data.Question)

		fmt.Println("A.", data.AnswerOption_A, "\nB.", data.AnswerOption_B, "\nC.", data.AnswerOption_C, "\nD.", data.AnswerOption_D)

		fmt.Printf("Please Enter Your Option (A,B,C or D): ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		selectedOption := sc.Text()

		fmt.Println("Correct Answer :", data.CorrectOption, "\n")
		finalScore(selectedOption, data.CorrectOption)
	}
	return totalQuestion
}

var totalScore int

func finalScore(selectedOption string, CorrectOption string) {
	result := selectedOption
	answer := CorrectOption
	if result == answer {
		totalScore++
	}
}
func fileOperation(fileName *string) [][]string {
	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}
func main() {
	fileName := flag.String("csv", "Quiz.csv", "a csv file in the format of 'question & answer'")
	options := fileOperation(fileName)

	totalQuestion = startQuiz(options)
	fmt.Println("You Scored :", totalScore, "Total question :", totalQuestion)
	fmt.Println()
}
