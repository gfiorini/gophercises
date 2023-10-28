package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InvalidOptionError struct{}

func (m *InvalidOptionError) Error() string {
	return "Selezione non valida\n"
}
func (e Entry) Show() {
	fmt.Println(e)
}

func (e Entry) String() string {

	s := e.Title + "\n\n"
	for i := range e.Story {
		s = s + e.Story[i] + "\n"
	}
	s = s + "\n"
	for i := range e.Options {
		s = s + strconv.Itoa(i+1) + ") " + e.Options[i].Text + "\n"
	}
	return s
}

func (a *Adventure) PlayTextAdventure() {
	entry, ok := a.FindStartEntry()
	if !ok {
		panic("Impossibile iniziare l'avventura!\n")
	}
	for ok {
		entry.Show()
		if !entry.HasOptions() {
			ok = false
		} else {
			//come ottimizzare questo snippet ?
			opt, e := entry.GetPlayerOption()
			for e != nil {
				fmt.Println(e)
				opt, e = entry.GetPlayerOption()
			}
			entry, ok = a.GetNext(opt)
		}
	}
	fmt.Println("Fine!")
}

func (e Entry) GetPlayerOption() (Option, error) {
	o := Option{"ERROR", "ERROR"}
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		return o, err
	}
	s = strings.Trim(strings.TrimRight(s, "\r\n"), " ")
	i, err := strconv.Atoi(s)
	i = i - 1
	if err != nil {
		return o, err
	}
	if i < 0 || i >= len(e.Options) {
		return o, &InvalidOptionError{}
	}
	return e.Options[i], nil
}
