package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
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
	timer   time.Duration
	score   Score
	entries []QuizEntry
}

func NewQuiz(records [][]string, timer time.Duration, shuffle bool) *Quiz {

	q := Quiz{
		entries: make([]QuizEntry, len(records)),
		timer:   timer,
	}
	for i := range records {
		q.entries[i].question = records[i][0]
		q.entries[i].answer = records[i][1]
	}

	if shuffle {
		rand.Shuffle(len(q.entries), func(i, j int) {
			q.entries[i], q.entries[j] = q.entries[j], q.entries[i]
		})
	}

	return &q
}

func main() {

	filename := flag.String("filename", "problems.csv", "input problems file")
	qt := flag.Int("timer", 30, "Quiz timer")
	shuffle := flag.Bool("shuffle", true, "Shuffle questions")

	flag.Parse()

	records := readCSV(filename)

	q := NewQuiz(records, time.Duration(*qt)*time.Second, *shuffle)

	c := make(chan bool)
	t := make(chan bool)

	go doQuiz(q, c, t)

	select {
	case <-c:
	case <-t:
		fmt.Println("Tempo scaduto!")
	}

	fmt.Printf("Hai risposto correttamente a %v domanda/e su %v\n", q.score, len(q.entries))

}

func readCSV(filename *string) [][]string {

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

func doQuiz(q *Quiz, c chan bool, t chan bool) {

	fmt.Println("Premi un tasto per partire")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		c <- false
	}
	go timer(t, q.timer)

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

func timer(tc chan bool, d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
	tc <- true
}
