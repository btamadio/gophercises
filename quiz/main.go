package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
//	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

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

	numCorrect := 0
	numProblems := len(records)

	for i, record := range records {
		problem := record[0]
		solution := record[1]

		fmt.Printf("Problem #%d: %s: = ", i+1, problem)
		text, err := inputReader.ReadString('\n')
		if err != nil{
			log.Fatal(err)
		}

		if strings.Trim(text, "\n ") == solution{
			numCorrect +=1
			fmt.Println("good job!")
		}
	}
	fmt.Printf("You scored %d out of %d\n", numCorrect, numProblems)
}