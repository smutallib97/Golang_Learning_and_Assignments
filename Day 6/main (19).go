package main

import (
	"fmt"
	"regexp"
)

func main() {
	words := [...]string{"sunday", "Monday", "TueSday", "Saturday"}

	pattern := "^[MT].*y"

	re, _ := regexp.Compile(pattern)

	for _, word := range words {

		found := re.MatchString(word)

		if found {

			fmt.Printf("%s matches\n", word)
		} else {

			fmt.Printf("%s does not match\n", word)
		}
	}
}
