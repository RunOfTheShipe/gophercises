package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {

	// 0. get CSV file path from CLI arg
	var fileName string
	flag.StringVar(&fileName, "file", "./problems.csv", "path to Q&A file")
	flag.Parse()

	fmt.Printf("file: %v\n", fileName)

	// 1. read a CSV file; parse into q/a
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("error reading file! ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("error parsing CSV file! ", err)
	}

	// 2. for each q/a, display Q, wait for A

	var questions int = 0
	var correct int = 0

	for _, line := range records {

		fmt.Printf("Question: %v\nAnswer: ", line[0])
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')

		trimAnswer := strings.TrimSpace(answer)
		trimCorrect := strings.TrimSpace(line[1])

		// 3. after a, keep track of successes/failures
		questions += 1
		if strings.EqualFold(trimAnswer, trimCorrect) {
			correct += 1
		}

		fmt.Println()
	}

	var percent float64 = (float64(correct) / float64(questions)) * 100
	fmt.Printf("Results %d/%d (%d%%)\n", correct, questions, int(math.Round(percent)))
}
