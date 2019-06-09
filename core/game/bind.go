package game

import (
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

//NewGame ... Creates a new game
func NewGame() Game {
	return &game{
		result: make([]result.Result, 0),
	}
}
