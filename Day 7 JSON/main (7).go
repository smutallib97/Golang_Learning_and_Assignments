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
    jsonInput := `{ 
    "Name": "Sachin Tendulkar",
    "Period": "1989-2012",
    "MatchesPlayed": 463,
    "RunsScored": 8426,
    "AvgScore": 44.83
    }`

    var player Player
    
    err := json.Unmarshal([]byte(jsonInput), &player)

    if err != nil {
        fmt.Println("JSON decode error!")
        return
    }

    fmt.Println(player)
    
}