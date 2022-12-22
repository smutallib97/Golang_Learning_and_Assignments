package main

import (
	"fmt"
	"sort"
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
	fmt.Println("--------------------------------------------------------")
	for states := range stateCapitals {
		fmt.Println(states, "->", stateCapitals[states])
	}
	fmt.Println("--------------------------------------------------------")
	for state, capital := range stateCapitals {
		fmt.Printf(" %s =  %s\n", state, capital)
	}
	fmt.Println("--------------------------------------------------------")
	var state1 []string
	for eachState := range stateCapitals {
		state1 = append(state1, eachState)
	}
	fmt.Println("--------------------------------------------------------")
	sort.Strings(state1)
	fmt.Println("Sorted States : ", state1)

}
