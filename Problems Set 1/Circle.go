/*
Write a program which takes diameter (float32) of a circle
and computes perimeter and area of the circle
*/

package main

import "fmt"

func main() {
	var d float32
	var area float32
	var perimeter float32

	fmt.Print("Enter Diameter of a circle: \n")
	fmt.Scanln(&d)

	perimeter = d * 3.14
	fmt.Print("Perimeter of a circle is = ", perimeter)

	area = 0.25 * 3.14 * d * d
	fmt.Print("\n Area of circle = ", area)

}
