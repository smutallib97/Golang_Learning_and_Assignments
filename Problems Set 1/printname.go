/*
From a user, read his/her friend’s first name and last
name. Print the full name of this friend with space
between first name and last name.
*/

package main

import "fmt"

func main() {
	var firstName string
	var lastName string

	fmt.Print("Enter Your Friend's First Name : \n")
	fmt.Scanln(&firstName)
	fmt.Print("Enter Your Friend's Last Name : \n")
	fmt.Scanln(&lastName)

	fmt.Print("Your Friend's Full Name is : ")
	fmt.Print(firstName, " "+lastName)
}
