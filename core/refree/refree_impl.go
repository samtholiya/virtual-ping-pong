package refree

import (
	"errors"
	"log"
	"math"

	"github.com/samtholiya/virtual-ping-pong/core/game"
	"github.com/samtholiya/virtual-ping-pong/core/player"
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

type refree struct {
	playersLen       int
	players          []player.Player
	game             game.Game
	roundsCount      int
	results          []result.Result
	currentPlayerLen int
	gameCreator      game.Creator
}

func (r *refree) SetPlayerCount(count int) error {
	logValue := math.Log2(float64(count))
	if logValue != float64(int64(logValue)) {
		return errors.New("Player count is invalid")
	}
	r.roundsCount = int(logValue)
	r.playersLen = count
	r.currentPlayerLen = count
	return nil
}

func (r *refree) StartRounds(playerChannel chan player.Player, resultChannel chan []result.Result) {
	r.results = make([]result.Result, 0)
	r.waitForPlayers(playerChannel)
	for i := 0; i < r.roundsCount; i++ {
		r.executeRounds(i + 1)
	}
	resultChannel <- r.results
}

func (r *refree) executeRounds(level int) {
	resultChannel := make(chan []result.Result, 0)
	winnerPlayerChannel := make(chan player.Player, 0)
	for i := 0; i < r.currentPlayerLen; i += 2 {
		game := r.gameCreator()
		attacker := r.players[i]
		defender := r.players[i+1]
		game.AddAttacker(attacker)
		game.AddDefender(defender)
		game.SetLevel(level)
		go game.Play(resultChannel, winnerPlayerChannel)
	}
	winnerPlayers := make([]player.Player, 0)

	for i := 0; i < r.currentPlayerLen/2; i++ {
		results := <-resultChannel
		r.results = append(r.results, results...)
		winnerPlayers = append(winnerPlayers, <-winnerPlayerChannel)
	}
	r.players = winnerPlayers
	r.currentPlayerLen = len(winnerPlayers)
}

func (r *refree) waitForPlayers(playerChannel chan player.Player) {
	r.players = make([]player.Player, 0)
	for i := 0; i < r.playersLen; i++ {
		player := <-playerChannel
		log.Printf("[INFO] Player :%v registered\n", player.GetName())
		r.players = append(r.players, player)
	}
}

func (r *refree) SetGameCreator(gameCreator game.Creator) error {
	if gameCreator == nil {
		return errors.New("Error can not be nil")
	}
	r.gameCreator = gameCreator
	return nil
}

//GetResults ... Get the results
func (r *refree) GetResults() []result.Result {
	return nil
}

//GetGameCreator ... Returns Game creator
func (r refree) GetGameCreator() game.Creator {
	return r.gameCreator
}
