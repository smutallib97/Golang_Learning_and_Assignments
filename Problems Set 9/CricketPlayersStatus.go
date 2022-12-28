package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	csvFile, err := os.Open("Cricket_Players_Stats.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	_, err = csvReader.Read()
	if err != nil {
		panic(err)
	}

	modifiedFile, err := os.Create("updated_Cricket_Players_Stats.csv")
	if err != nil {
		panic(err)
	}
	defer modifiedFile.Close()

	csvWriter := csv.NewWriter(modifiedFile)

	csvWriter.Write([]string{"Name", "MatchesPlayed", "TotalScore"})

	for {
		row, err := csvReader.Read()
		if err != nil {
			break
		}

		name := row[0]
		matchesPlayed, err := strconv.Atoi(row[2])
		if err != nil {
			panic(err)
		}
		totalScore, err := strconv.Atoi(row[3])
		if err != nil {
			panic(err)
		}

		matchesPlayed += 10
		totalScore += 500

		csvWriter.Write([]string{name, strconv.Itoa(matchesPlayed), strconv.Itoa(totalScore)})
	}
	fmt.Println("Updated File Created Successfully....!")

	csvWriter.Flush()
}
