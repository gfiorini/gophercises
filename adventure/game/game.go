package game

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type InvalidOptionError struct{}

func (m *InvalidOptionError) Error() string {
	return "Selezione non valida\n"
}

//type Node struct {
//	key   string
//	value string
//	nodes *[]Node
//}

type Option struct {
	Text string
	Arc  string
}
type Entry struct {
	Title   string
	Story   []string
	Options []Option
}
type Adventure struct {
	EntryMap map[string]Entry
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

func (e Entry) Show() {
	fmt.Println(e)
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

func (e Entry) HasOptions() bool {
	return len(e.Options) > 0
}

func (a *Adventure) FindStartEntry() (Entry, bool) {
	o := Option{"", "intro"}
	return a.GetNext(o)
}

func (a *Adventure) GetNext(opt Option) (Entry, bool) {
	ent, ok := a.EntryMap[opt.Arc]
	return ent, ok
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

func (e Entry) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var tmplFile = "adventure_template.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, e)
	if err != nil {
		panic(err)
	}

}

func (a *Adventure) PlayHtmlAdventure() {
	st, _ := a.FindStartEntry()
	http.Handle("/", st)
	for k := range a.EntryMap {
		entry := a.EntryMap[k]
		http.Handle("/"+k, entry)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func NewAdventureFromFile(filename string) *Adventure {
	b, e := os.ReadFile(filename)
	if e != nil {
		panic(e)
	}

	m := make(map[string]Entry)
	e = json.Unmarshal(b, &m)
	if e != nil {
		panic(e)
	}

	return &Adventure{
		EntryMap: m,
	}

}
