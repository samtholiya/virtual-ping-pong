package championship

import (
	"github.com/samtholiya/virtual-ping-pong/core/game"
	"github.com/samtholiya/virtual-ping-pong/core/player"
	"github.com/samtholiya/virtual-ping-pong/core/refree"
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

//Championship ... A tournament
type Championship interface {

	//AddRefree ... Refree for the championship
	AddRefree(refree.Refree) error

	//SetGameCreator ... Add the game initializer
	SetGameCreator(game.Creator) error

	//GetRegisterationChannel ... registeration channel for players
	GetRegisterationChannel() chan player.Player

	//StartTournament ... start the tournament
	StartTournament(result chan []result.Result) error

	//GetRegisterationChannel ... get the tournament channel
	GetResultsChannel() chan []result.Result
}
