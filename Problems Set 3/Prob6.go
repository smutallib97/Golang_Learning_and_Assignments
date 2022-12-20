package main

import "fmt"

func vowelsconsonantsCount(str string) (int, int) {

	var vCount, cCount int

	vCount = 0
	cCount = 0

	for i := 0; i < len(str); i++ {

		// if str[i] == ' ' {
		// 	continue
		// }

		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' ||
			str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U' {

			vCount++
		} else {
			cCount++
		}
	}
	return vCount, cCount
}
func main() {
	var str string

	var vCount, cCount int

	str = "Hello Go 1"

	vCount, cCount = vowelsconsonantsCount(str)
	fmt.Println("Sentence:-\n", str)

	fmt.Println("Result:- \nThe total number of vowels in the above sentence are", vCount)
	fmt.Println("The total number of consonants in the above sentence are", cCount)
}
