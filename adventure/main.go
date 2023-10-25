package main

import (
	"adventure/game"
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("filename", "adventure.json", "input json file")
	flag.Parse()
	a := game.NewAdventureFromFile(*filename)
	st := a.FindStartEntry()
	fmt.Println(st.Story[0])

}
