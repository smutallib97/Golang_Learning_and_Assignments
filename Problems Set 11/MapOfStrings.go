package main

import (
	"fmt"
	"strings"
)

func Imap(data []string, myFunc func(string) string) []string {
	mapped := make([]string, len(data))
	for i, value := range data {
		mapped[i] = myFunc(value)
	}
	return mapped
}

func main() {
	stringSlice := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}
	fmt.Println("Input Slice is: ", stringSlice)

	firstCharacterEachWordAsCapital := Imap(stringSlice, func(word string) string {
		return strings.Title(word)
	})
	fmt.Println("First character of each word as capital slice is: ", firstCharacterEachWordAsCapital)

	colonAtEndOfWord := Imap(stringSlice, func(word string) string {
		addcolon := word + ":"
		return addcolon
	})
	fmt.Println("Colon at end of words slice is: ", colonAtEndOfWord)
}
