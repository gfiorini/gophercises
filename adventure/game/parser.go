package game

import "os"

type Node struct {
	key   string
	value string
	nodes *[]Node
}

func ParseFile(filename string) *Node {
	_, e := os.ReadFile(filename)
	if e != nil {
		panic(e)
	}
	return nil
}
