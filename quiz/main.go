package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type quizEntry struct {
	question string
	answer   string
}

type score int

func main() {

	score := 0

	filename := flag.String("filename", "problems.csv", "input problems file")

	b, err := os.ReadFile(*filename)
	if err != nil {
		s := fmt.Sprintf("file %v non trovato!", *filename)
		panic(s)
	}

	log.Printf("Parsing file %v\n", *filename)
	r := csv.NewReader(bytes.NewReader(b))
	r.FieldsPerRecord = 2

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var quiz = make([]quizEntry, len(records))
	for i := range records {
		quiz[i].question = records[i][0]
		quiz[i].answer = records[i][1]
	}

	nq := len(quiz)
	for i := range quiz {
		fmt.Printf("Quiz %v/%v) - %v ?\n ", i+1, nq, quiz[i].question)
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		if quiz[i].answer == strings.Trim(strings.TrimRight(s, "\r\n"), " ") {
			score++
		}
	}

	fmt.Printf("Hai risposto correttamente a %v domande su %v", score, nq)

}
