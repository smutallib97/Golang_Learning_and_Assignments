package main

import (
	"fmt"
	"regexp"
)

func isPasswordValid(password string) bool {
	passwordPattern := regexp.MustCompile("^(?=.*\b)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$")
	return passwordPattern.MatchString(password)
}

func main() {
	var password string
	fmt.Println("Enter Password :")
	fmt.Scanln(&password)
	checkPassword := isPasswordValid(password)
	fmt.Println(checkPassword)
}
