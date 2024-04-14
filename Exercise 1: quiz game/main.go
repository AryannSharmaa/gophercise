package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, error := os.Open("problems.csv")
	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()
	filereader := csv.NewReader(file)
	rec, error := filereader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}

	correct, incorrect := 0, 0
	fmt.Println("Write correct answers")
	for _, v := range rec {
		result, error := (strconv.Atoi(v[1]))
		if error != nil {
			fmt.Println(error)
		}
		fmt.Print(v[0], "=")
		var answer int
		fmt.Scanln(&answer)
		if answer == result {
			correct++
		} else {
			incorrect++
		}
	}
	fmt.Printf("Total: %v correct: %v, incorrect: %v\n", correct+incorrect, correct, incorrect)
}
