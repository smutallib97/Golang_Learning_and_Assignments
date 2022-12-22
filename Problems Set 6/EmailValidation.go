package main

import (
	"fmt"
	"regexp"
)

func main() {
	emails := [...]string{
		"abc@yahoo.com",
		"abc-100@yahoo.com",
		"ç$€§/az@gmail.com",
		"abcd@gmail_yahoo.com",
		"abc.566@yahoo.com",
		"abc111@abc.com",
		"abc-111@abc.net",
	}

	//emailPattern := "^[a-z]+[a-z0-9+_.-]*[@][a-z0-9]+[.][a-z]{2,4}[.]*([a-z]{2,3})*$"
	emailPattern := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

	re, _ := regexp.Compile(emailPattern)

	for _, emails := range emails {

		found := re.MatchString(emails)

		if found {

			fmt.Printf("%s correct \n", emails)
		} else {

			fmt.Printf("%s Incorrect \n", emails)
		}
	}
}
