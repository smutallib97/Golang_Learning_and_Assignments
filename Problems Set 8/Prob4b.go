package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readSample(rs io.ReadSeeker) ([][]string, error) {
	row1, err := bufio.NewReader(rs).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = rs.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(rs)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func main() {
	f, err := os.Open("Cricket_Players_Stats.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rows, err := readSample(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}
