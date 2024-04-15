package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of problems.csv")
	timelimit := flag.Int("t", 30, "timer for quiz in seconds")
	flag.Parse()
	file, error := os.Open(*filename)
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
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	for _, v := range rec {
		fmt.Print(v[0], "=")
		answerCh := make(chan int)
		result, error := (strconv.Atoi(v[1]))
		if error != nil {
			fmt.Println(error)
		}

		go func() {
			var answer int
			fmt.Scanln(&answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTotal: %v correct: %v, incorrect: %v\n", correct+incorrect, correct, incorrect)
			return
		case answer := <-answerCh:
			if answer == result {
				correct++
			} else {
				incorrect++
			}
		}
	}
	fmt.Printf("Total: %v correct: %v, incorrect: %v\n", correct+incorrect, correct, incorrect)
}
