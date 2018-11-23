package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	//read the problems.csv file or the filename provided from the flag
	filename := flag.String("file", "problems.csv", "The file containing questions")
	timeLimit := flag.Int("timelimit", 30, "Amount of seconds to answer quiz")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	askQuestions(records, *timeLimit)

}
