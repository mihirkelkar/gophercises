package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func askQuestions(records [][]string, timeLimit int) {
	correct, incorrect := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	// declare a channel to gather scanned answers from.
	ansChan := make(chan string)

mainloop:
	for _, record := range records {
		fmt.Printf("What is %s ?\n", record[0])
		go func() {
			if scanner.Scan() {
				ansChan <- scanner.Text()
			}
		}()
		// select to see if you see the timer timeout first or the answers
		// get done first
		select {
		case <-timer.C:
			break mainloop
		case answer := <-ansChan:
			if answer == record[1] {
				correct++
			} else {
				incorrect++
			}
		}
	}
	fmt.Printf("You got %d correct out of %d questions\n", correct, len(records))
}
