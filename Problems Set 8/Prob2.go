package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type MyFriend struct {
	Name                  string
	MobileNumber          string
	AlternateMobileNumber string
	Address               string
	City                  string
}

func main() {

	file, err := os.Create("MyFriendsAddressBook.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	friendDetails := [][]string{
		{"Name", "MobileNumber", "AlternateMobileNumber", "Address", "City"},
		{"Tauqeer", "8180883610", "9875154823", "Mumbra", "Mumbai"},
		{"Mujahid", "9623154878", "8996364548", "Taj Nagar", "Malkapur"},
		{"Sahil", "8885852547", "7858451263", "Ajuba Nagar", "Nagpur"},
		{"Owais", "8999634875", "9673451852", "Hinjewadi", "Pune"},
		{"Junaid", "9518487563", "9875154825", "Mumbra", "Mumbai"},
		{"Nitish", "9848756318", "0", "Kashi Mandir", "Patna"},
		{"Ullas", "8885524818", "0", "Kochipawli", "Bangluru"},
		{"Sailesh", "9875154823", "0", "Amirpeth", "Hydrabbad"},
		{"Sunil", "8180883611", "0", "Sangram Nagar", "Surat"},
		{"Nadeem", "9875484823", "0", "Shahdara", "Delhi"},
	}

	writeToFile := csv.NewWriter(file)
	err = writeToFile.WriteAll(friendDetails)
	if err != nil {
		panic(err)
	}

	fmt.Println("File Successfully Created")
}
