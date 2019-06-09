package game

import (
	"github.com/samtholiya/virtual-ping-pong/core/player"
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

//Game ... Game has attacker, player
type Game interface {

	//SetWinningScore ... Set the winning score of the game
	SetWinningScore(int) error

	//AddAttacker ... Add the attacker of the game
	AddAttacker(player.Player) error

	//AddDefender ... Add the defender of the game
	AddDefender(player.Player) error

	//SetLevel ... Set the level of the game
	SetLevel(int) error

	//Play ... Play the game
	Play(resultsChannel chan []result.Result, winnerPlayer chan player.Player)
}

//Creator ... A func which returns an instance of Game
type Creator func() Game
