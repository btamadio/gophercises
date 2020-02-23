package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func askQuestions(records [][]string, result chan<- bool, finished chan<- bool) {
	rd := bufio.NewReader(os.Stdin)

	for i, record := range records {

		fmt.Printf("Problem #%d: %s: = ", i+1, record[0])

		text, err := rd.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		result <- strings.Trim(text, "\n ") == record[1]
	}
	finished <- true
}

func main() {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	csvfile, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	records, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Press enter to begin quiz")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')


	problemResult := make(chan bool)
	finished := make(chan bool)
	go askQuestions(records, problemResult, finished)
	timeout := time.After(time.Duration(*timeLimit) * time.Second)

	numCorrect := 0
	Loop:
		for {
			select {

			case correctAnswer := <-problemResult:
				if correctAnswer {
					numCorrect++
				}

			case <-timeout:
				fmt.Printf("\nTime's Up! You scored %d out of %d.\n", numCorrect, len(records))
				os.Exit(0)

			case <-finished:
				break Loop

			}
		}
	fmt.Printf("Quiz completed! You scored %d out of %d.\n", numCorrect, len(records))
}
