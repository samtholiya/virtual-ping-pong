package main

import (
	"fmt"
	"strconv"

	"github.com/samtholiya/virtual-ping-pong/core/championship"
	"github.com/samtholiya/virtual-ping-pong/core/game"
	"github.com/samtholiya/virtual-ping-pong/core/player"
	"github.com/samtholiya/virtual-ping-pong/core/refree"
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

func getGameConfig() game.Game {
	game := game.NewGame()
	game.SetWinningScore(5)
	return game
}

func main() {
	resultsChannel := make(chan []result.Result)

	championship := championship.NewChampionship()

	refree := refree.NewRefree()
	refree.SetPlayerCount(8)

	championship.AddRefree(refree)
	championship.SetGameCreator(getGameConfig)
	registerationChannel := championship.GetRegisterationChannel()
	championship.StartTournament(resultsChannel)

	for i := 0; i < 8; i++ {
		newPlayer := player.NewPlayer(strconv.Itoa(i))
		newPlayer.SetDefenceSize(i + 1)
		go newPlayer.RegisterChampionship(registerationChannel)
	}
	results := <-resultsChannel
	for _, res := range results {
		fmt.Printf("Winner for level %v round %v\nAttacker: %v Defender %v was %v\n", res.GetMatchLevel(),
			res.GetRoundLevel(), res.GetAttackerName(),
			res.GetDefenderName(), res.GetWinnerName())
	}
}
