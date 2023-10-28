package main

import (
	"adventure/game"
	"flag"
)

func main() {
	filename := flag.String("filename", "adventure.json", "input json file")
	flag.Parse()
	a := game.NewAdventureFromFile(*filename)
	//a.PlayTextAdventure()
	a.PlayHtmlAdventure()

}
