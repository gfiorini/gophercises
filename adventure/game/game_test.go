package game_test

import (
	"adventure/game"
	"testing"
)

func TestParse(t *testing.T) {
	game.ParseFile("../adventure.json")
}
