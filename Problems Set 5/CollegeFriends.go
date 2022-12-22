package main

import "fmt"

type MyFriend struct {
	fname                 string
	mobileNumber          int
	alternateMobileNumber int
	address               string
	city                  string
	collegeFriend         bool
}

func main() {

	friend1 := MyFriend{"Tauqeer", 8180883610, 9875154823, "Mumbra", "Mumbai", true}
	friend2 := MyFriend{"Mujahid", 9623154878, 8996364548, "Taj Nagar", "Malkapur", false}
	friend3 := MyFriend{"Sahil", 8885852547, 7858451263, "Ajuba Nagar", "Nagpur", true}
	friend4 := MyFriend{"Owais", 8999634875, 9673451852, "Hinjewadi", "Pune", true}
	friend5 := MyFriend{"Junaid", 9518487563, 9875154825, "Mumbra", "Mumbai", false}
	friend6 := MyFriend{"Nitish", 9848756318, 0, "Kashi Mandir", "Patna", false}
	friend7 := MyFriend{"Ullas", 8885524818, 0, "Kochipawli", "Bangluru", false}
	friend8 := MyFriend{"Sailesh", 9875154823, 0, "Amirpeth", "Hydrabbad", false}
	friend9 := MyFriend{"Sunil", 8180883611, 0, "Sangram Nagar", "Surat", true}
	friend10 := MyFriend{"Nadeem", 9875484823, 0, "Shahdara", "Delhi", true}

	Details := []MyFriend{
		friend1,
		friend2,
		friend3,
		friend4,
		friend5,
		friend6,
		friend7,
		friend8,
		friend9,
		friend10,
	}
	for _, MyfriendDetails := range Details {
		fmt.Println("[", MyfriendDetails.fname, " , ", MyfriendDetails.mobileNumber, " , ", MyfriendDetails.alternateMobileNumber, " , ", MyfriendDetails.address, " , ", MyfriendDetails.city, "]")
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println("Displaying Names and Mobile numbers of MyFriends")
	fmt.Println("")
	for _, MyfriendDetails := range Details {
		fmt.Println("[", MyfriendDetails.fname, " - ", MyfriendDetails.mobileNumber, "]")
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("Displaying friend names Who are having alternate mobile number")
	for _, MyfriendDetails := range Details {
		if MyfriendDetails.alternateMobileNumber != 0 {
			fmt.Println(MyfriendDetails.fname)
		}
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println("Displaying friend names Who are not having alternate mobile number")
	for _, MyfriendDetails := range Details {
		if MyfriendDetails.alternateMobileNumber == 0 {
			fmt.Println(MyfriendDetails.fname)
		}
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println("Displaying college friends names")
	for _, friendDetails := range Details {
		if friendDetails.collegeFriend != false {
			fmt.Println(friendDetails.fname, " - ", friendDetails.city)
		}
	}
}
