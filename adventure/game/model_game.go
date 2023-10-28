package game

import (
	"encoding/json"
	"os"
)

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
