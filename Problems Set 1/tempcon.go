/*
Develop a Temperature Unit Converter Program. It can
convert temperature from centigrade to Fahrenheit.
○ It asks user to enter temperature in degree c
○ It converts in Fahrenheit and displays the same
*/

package main

import "fmt"

func main() {
	var temp_c float32
	var temp_f float32

	fmt.Print("Enter Temperature in centigrade: \n")
	fmt.Scanln(&temp_c)

	temp_f = ((9.0 / 5.0) * temp_c) + 32.0

	fmt.Print(temp_f, " degrees Fahrenheit ")

}
