package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi")
	file, error := os.Open("problems.csv")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("file opened")
	defer file.Close()

	filereader := csv.NewReader(file)
	rec, error := filereader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(rec)
}
