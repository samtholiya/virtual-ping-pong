package refree

import (
	"github.com/samtholiya/virtual-ping-pong/core/game"
	"github.com/samtholiya/virtual-ping-pong/core/player"
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

//Refree ... A refree decides who is the winner
type Refree interface {

	//SetPlayerCount ... Set Players for entire Game
	SetPlayerCount(int) error

	//StartRounds ... Start Game rounds
	StartRounds(chan player.Player, chan []result.Result)

	//SetGameCreator ... Game Config
	SetGameCreator(game.Creator) error

	//GetGameCreator ... Game Creator
	GetGameCreator() game.Creator
}
