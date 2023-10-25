package game

import "fmt"

type Game struct {
	filename string
	start    *Node
}

func (g *Game) Play() {
	fmt.Printf("Inizia l'avventura '%v' \n", g.filename)
}

func NewGame(filename string) (game *Game) {
	st := ParseFile(filename)
	g := Game{
		filename: filename,
		start:    st,
	}
	return &g
}
