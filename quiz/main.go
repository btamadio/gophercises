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


func askQuestions(records [][]string, result chan<- bool){
	for i, record := range records {
		problem := record[0]
		solution := record[1]
		rd := bufio.NewReader(os.Stdin)
		fmt.Printf("Problem #%d: %s: = ", i+1, problem)
		text, err := rd.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.Trim(text, "\n ") == solution {
			result <- true
		} else {
			result <- false
		}
	}
}

func main(){
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	csvfile, err := os.Open(*fileName)
	if err != nil{
		log.Fatal(err)
		return
	}

	csvReader := csv.NewReader(csvfile)
	records, err := csvReader.ReadAll()
	if err != nil{
		log.Fatal(err)
	}

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Print("Press enter to begin quiz")
	_, _ = inputReader.ReadString('\n')

	numCorrect := 0
	result := make(chan bool)

	go askQuestions(records, result)
	timeout := time.After(time.Duration(*timeLimit) * time.Second)

	for _, _ = range records{
		select{
		case correct := <- result:
			if correct {
				numCorrect++
			}
		case <- timeout:
			fmt.Printf("\nTime's Up! You scored %d out of %d.\n", numCorrect, len(records))
			os.Exit(0)
		}
	}
	fmt.Printf("You scored %d out of %d.\n", numCorrect, len(records))
}