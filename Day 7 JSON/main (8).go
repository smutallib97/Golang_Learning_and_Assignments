package main

import (
    "fmt"
    "encoding/json"
)
type Player struct {
    Name string
    Period string
    MatchesPlayed, RunsScored int
    AvgScore float64
}
func main() {
    jsonInput := ` {
    "Dhoni": 347,
    "Dravid": 340,
    "Ganguly": 308,
    "Kohli": 262,
    "Sachin": 463
    }`

    matchesPlayed := make(map[string]int)
    
    err := json.Unmarshal([]byte(jsonInput), &matchesPlayed)
    if err != nil {
        fmt.Println("JSON decode error!")
        return
    }
    fmt.Println( matchesPlayed)
}