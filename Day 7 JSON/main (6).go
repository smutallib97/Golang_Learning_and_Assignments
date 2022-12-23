package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Name                      string
	Period                    string
	MatchesPlayed, RunsScored int
	AvgScore                  float64
}

func main() {
	fmt.Println("Creating an variable of Type Player and print the same")
	player1 := Player{"Sachin Tendulkar", "1989-2012", 463, 18426, 44.83}
	player2 := Player{"Saurav Ganguly", "1992-2007", 308, 11221, 40.95}
	player4 := Player{"Rahul Dravid", "1996-2011", 340, 10768, 39.15}
	players := []Player{
		player1,
		player2,
		player4,
	}
	for _, plyr := range players {
		fmt.Println(plyr)
	}
	fmt.Println(player1)
	fmt.Println("")
	fmt.Println("Converting this variable to json")
	bytes, _ := json.Marshal(player1)
	fmt.Printf("%T %s", bytes, string(bytes))

	fmt.Println("------------")

	fmt.Println("Creating a map and printing the same")
	matchesPlayed := make(map[string]int)

	matchesPlayed["Sachin"] = 463
	matchesPlayed["Ganguly"] = 308
	matchesPlayed["Dhoni"] = 347
	matchesPlayed["Kohli"] = 262
	matchesPlayed["Dravid"] = 340
	fmt.Println(matchesPlayed)
	fmt.Println("")

	fmt.Println("Converting this variable to json")
	bytes, _ = json.Marshal(matchesPlayed)
	fmt.Println(matchesPlayed)
	fmt.Println("")

	fmt.Println("Converting this variable to json")
	bytes, _ = json.Marshal(matchesPlayed)
	fmt.Println(string(bytes))

	fmt.Printf("\nConverting slice of structs to JSON \n")
	fmt.Printf("Used Pretty printing \n")
	//bytesStructSlice, _ := json.Marshal(players)
	bytesStructSlice, _ := json.MarshalIndent(players, " ", "\t")
	fmt.Println(string(bytesStructSlice))

}
