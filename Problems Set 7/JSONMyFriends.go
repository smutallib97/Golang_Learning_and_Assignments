package main

import (
	"encoding/json"
	"fmt"
)

type MyFriend struct {
	fname                 string
	mobileNumber          int
	alternateMobileNumber int
	address               string
	city                  string
}

func main() {

	Details := []MyFriend{
		{fname: "Tauqeer", mobileNumber: 8180883610, alternateMobileNumber: 9875154823, address: "Mumbra", city: "Mumbai"},
		{fname: "Mujahid", mobileNumber: 9623154878, alternateMobileNumber: 8996364548, address: "Taj Nagar", city: "Malkapur"},
		{fname: "Sahil", mobileNumber: 8885852547, alternateMobileNumber: 7858451263, address: "Ajuba Nagar", city: "Nagpur"},
		{fname: "Owais", mobileNumber: 8999634875, alternateMobileNumber: 9673451852, address: "Hinjewadi", city: "Pune"},
		{fname: "Junaid", mobileNumber: 9518487563, alternateMobileNumber: 9875154825, address: "Mumbra", city: "Mumbai"},
		{fname: "Nitish", mobileNumber: 9848756318, alternateMobileNumber: 0, address: "Kashi Mandir", city: "Patna"},
		{fname: "Ullas", mobileNumber: 8885524818, alternateMobileNumber: 0, address: "Kochipawli", city: "Bangluru"},
		{fname: "Sailesh", mobileNumber: 9875154823, alternateMobileNumber: 0, address: "Amirpeth", city: "Hydrabbad"},
		{fname: "Sunil", mobileNumber: 8180883611, alternateMobileNumber: 0, address: "Sangram Nagar", city: "Surat"},
		{fname: "Nadeem", mobileNumber: 9875484823, alternateMobileNumber: 0, address: "Shahdara", city: "Delhi"},
	}
	json_addressbook, err := json.Marshal(Details)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json_addressbook))
}
