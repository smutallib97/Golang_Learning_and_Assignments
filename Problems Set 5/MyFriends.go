/*
4. Address Book of MyFriends:
Create an Address-Book of your friends.
It will have fields like - name, Mobile Number, Alternate Mobile number,
Address, City/Town Name. Some friends will have Alternate Mobile numbers
and some will not have.
Populate the same with 10 of your friend's details.
Keep these data in an iterable data type
Iterate through this struct and print name and Mobile number s
Also print names of friends who have alternate phone numbers
*/

package main

import "fmt"

type MyFriend struct {
	fname                 string
	mobileNumber          int
	alternateMobileNumber int
	address               string
	city                  string
}

func main() {

	friend1 := MyFriend{"Tauqeer", 8180883610, 9875154823, "Mumbra", "Mumbai"}
	friend2 := MyFriend{"Mujahid", 9623154878, 8996364548, "Taj Nagar", "Malkapur"}
	friend3 := MyFriend{"Sahil", 8885852547, 7858451263, "Ajuba Nagar", "Nagpur"}
	friend4 := MyFriend{"Owais", 8999634875, 9673451852, "Hinjewadi", "Pune"}
	friend5 := MyFriend{"Junaid", 9518487563, 9875154825, "Mumbra", "Mumbai"}
	friend6 := MyFriend{"Nitish", 9848756318, 0, "Kashi Mandir", "Patna"}
	friend7 := MyFriend{"Ullas", 8885524818, 0, "Kochipawli", "Bangluru"}
	friend8 := MyFriend{"Sailesh", 9875154823, 0, "Amirpeth", "Hydrabbad"}
	friend9 := MyFriend{"Sunil", 8180883611, 0, "Sangram Nagar", "Surat"}
	friend10 := MyFriend{"Nadeem", 9875484823, 0, "Shahdara", "Delhi"}

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

	//MyFriends := make(map[int]MyFriend)
	// for city, mobileNumber := range MyFriends {
	// 	fmt.Printf("City -  %s = Mobile Number  %d\n", city, mobileNumber)
	// }

	// fmt.Println(MyFriends)

}
