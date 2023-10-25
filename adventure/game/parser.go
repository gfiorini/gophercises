package game

import (
	"encoding/json"
	"os"
)

type Node struct {
	key   string
	value string
	nodes *[]Node
}

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

func (a *Adventure) FindStartEntry() Entry {
	return a.EntryMap["intro"]
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
