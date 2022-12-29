package main

import (
	"fmt"
	"strings"
)

func Filter(stringSlice []string, funcName func(string) bool) []string {
	filteredstring := []string{}
	for _, Words := range stringSlice {
		if funcName(Words) {
			filteredstring = append(filteredstring, Words)
		}
	}
	return filteredstring
}

func main() {
	stringSlice := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}
	fmt.Println("Input Slice is: ", stringSlice)

	wordsStartWithB := Filter(stringSlice, func(word string) bool {
		return strings.HasPrefix(word, "b")
	})
	fmt.Println("Words Start with B: ", wordsStartWithB)

	wordsHave3Letters := Filter(stringSlice, func(word string) bool {
		return len(word) == 3
	})
	fmt.Println("Words Which having 3 letters: ", wordsHave3Letters)
}
