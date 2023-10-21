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
	"time"
)

type Quiz struct {
}
type quizEntry struct {
	question string
	answer   string
}

type score int

func main() {

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

	c := make(chan bool)
	t := make(chan bool)

	go doQuiz(records, c)
	go timer(t)

	select {
	case <-c:
	case <-t:
		fmt.Println("Tempo scaduto!")
	}
}

func doQuiz(records [][]string, c chan bool) {

	var score score
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
			c <- false
		}

		if quiz[i].answer == strings.Trim(strings.TrimRight(s, "\r\n"), " ") {
			score++
		}
	}

	fmt.Printf("Hai risposto correttamente a %v domanda/e su %v", score, nq)
	c <- true
}

func timer(t chan bool) {
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	t <- true
}
