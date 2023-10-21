package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Score int

type QuizEntry struct {
	question string
	answer   string
}

type Quiz struct {
	score   Score
	entries []QuizEntry
}

func NewQuiz(records [][]string) *Quiz {

	q := Quiz{
		entries: make([]QuizEntry, len(records)),
	}
	for i := range records {
		q.entries[i].question = records[i][0]
		q.entries[i].answer = records[i][1]
	}
	return &q
}

func main() {

	records := readCSV()
	q := NewQuiz(records)

	c := make(chan bool)
	t := make(chan bool)

	go doQuiz(q, c)
	go timer(t)

	select {
	case <-c:
	case <-t:
		fmt.Println("Tempo scaduto!")
	}

	fmt.Printf("Hai risposto correttamente a %v domanda/e su %v\n", q.score, len(q.entries))
}

func readCSV() [][]string {
	filename := flag.String("filename", "problems.csv", "input problems file")

	b, err := os.ReadFile(*filename)
	if err != nil {
		s := fmt.Sprintf("file %v non trovato!", *filename)
		panic(s)
	}

	r := csv.NewReader(bytes.NewReader(b))
	r.FieldsPerRecord = 2

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}

func doQuiz(q *Quiz, c chan bool) {

	nq := len(q.entries)
	for i := range q.entries {
		fmt.Printf("Quiz %v/%v) - %v ?\n ", i+1, nq, q.entries[i].question)
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		if err != nil {
			c <- false
		}

		if q.entries[i].answer == strings.Trim(strings.TrimRight(s, "\r\n"), " ") {
			q.score++
		}
	}
	c <- true
}

func timer(t chan bool) {
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	t <- true
}
