package championship

import (
	"strconv"
	"testing"

	"github.com/samtholiya/virtual-ping-pong/core/core_helper_test"

	"github.com/samtholiya/virtual-ping-pong/core/game"
	"github.com/samtholiya/virtual-ping-pong/core/refree"
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

func getGameConfig() game.Game {
	game := game.NewGame()
	game.SetWinningScore(5)
	return game
}

func TestChampionshipTournament(t *testing.T) {
	resultsChannel := make(chan []result.Result)

	championship := NewChampionship()

	refree := refree.NewRefree()
	refree.SetPlayerCount(8)

	championship.AddRefree(refree)
	championship.SetGameCreator(getGameConfig)
	registerationChannel := championship.GetRegisterationChannel()
	championship.StartTournament(resultsChannel)

	for i := 0; i < 8; i++ {
		newPlayer := core_helper_test.PlayerDummy{
			Name:         strconv.Itoa(i + 1),
			DefenceSlice: []int{1, 2, 4, 5, 6},
			AttackNumber: 3,
		}
		newPlayer.SetDefenceSize(i)
		go newPlayer.RegisterChampionship(registerationChannel)
	}
	results := <-resultsChannel
	if len(results) != 35 {
		t.Error("Total results are incorrect")
	}
}
