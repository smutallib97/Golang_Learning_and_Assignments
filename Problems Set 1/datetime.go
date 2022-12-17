package main

import (
	"fmt"
	"time"
)

const (
	DDMMYYYY = "14-12-2022"
)

func main() {
	//fmt.Println("Current Date and time is: ", DDMMYYYY)

	y := time.Now().Year()
	fmt.Println("Current Year is: ", y)

	m := time.Now().Month()
	fmt.Println("Current Month is: ", m)

	d := time.Now().Day()
	fmt.Println("Current Date is: ", d)

	//now := time.Now().UTC()
	//fmt.Println(time.Now().Format(DDMMYYYY))

	now := time.Now().UTC()
	fmt.Println(now.Format(DDMMYYYY))
}
