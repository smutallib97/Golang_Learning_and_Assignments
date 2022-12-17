package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	const (
		IS_FULL_TIME          = 2
		IS_PART_TIME          = 1
		RATE_PER_HOUR         = 125
		WORKING_DAYS_IN_MONTH = 20
	)

	var empHourInDay, totalEmpHour int

	for i := 0; i < WORKING_DAYS_IN_MONTH; i++ {
		empCheck := rand.Intn(3)
		switch empCheck {
		case IS_FULL_TIME:
			fmt.Println("Employee is Full Time Present")
			empHourInDay = 8
		case IS_PART_TIME:
			fmt.Println("Employee is Part Time Present")
			empHourInDay = 4
		default:
			fmt.Println("Employee is Absent")
			empHourInDay = 0
		}
		totalEmpHour += empHourInDay
		fmt.Println(i, empCheck, empHourInDay, totalEmpHour)
	}
	fmt.Println("--")
	empWage := RATE_PER_HOUR * totalEmpHour

	fmt.Printf("Employee has spent %d hours & thus wage is Rs: %d \n", totalEmpHour, empWage)

}
