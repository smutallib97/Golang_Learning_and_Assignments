/*
Develop a program that reads a number from user as
string and print 10 times of that number as string only.
Suppose the number entered is 12 Output will be 120
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 string
	var zero int
	fmt.Println("Enter a number : ")
	fmt.Scanln(&n1)
	zero = 0
	fmt.Println("Printing 10 times of given number :")
	fmt.Println(n1 + strconv.Itoa(zero))

}
