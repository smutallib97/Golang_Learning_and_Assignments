package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//var stateCapitals map[string]string

	stateCapitals := make(map[string]string)

	stateCapitals["Gujrat"] = "Gandhinagar"
	stateCapitals["Maharashtra"] = "Mumbai"
	stateCapitals["Karnataka"] = "Bengaluru"
	stateCapitals["Jharkhand"] = "Ranchi"
	stateCapitals["Bihar"] = "Patna"
	stateCapitals["Madhya Pradesh"] = "Bhopal"
	stateCapitals["Punjab"] = "Chandigarh"
	stateCapitals["Rajasthan"] = "Jaipur"
	stateCapitals["Tamil Nadu"] = "Chennai"
	stateCapitals["Goa"] = "Panaji"

	fmt.Printf("The data type of stateCapital is %T\n", stateCapitals)
	fmt.Println("10 states of india and their state capitals are:  \n", stateCapitals)

	for states := range stateCapitals {
		fmt.Println(states, "->", stateCapitals[states])
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println("Displaying by using JSON")

	stateMap, err := json.MarshalIndent(stateCapitals, " ", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("10 states of india and their state capitals are:  \n", string(stateMap))
}
