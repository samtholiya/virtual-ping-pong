package refree

import (
	"testing"

	"github.com/samtholiya/virtual-ping-pong/core/result"

	"github.com/samtholiya/virtual-ping-pong/core/core_helper_test"
	"github.com/samtholiya/virtual-ping-pong/core/player"

	"github.com/samtholiya/virtual-ping-pong/core/game"
)

func TestRefreeRounds(t *testing.T) {
	r := NewRefree()
	r.SetGameCreator(func() game.Game {
		g := game.NewGame()
		g.SetWinningScore(3)
		return g
	})
	r.SetPlayerCount(2)
	attacker := core_helper_test.PlayerDummy{
		Name:         "Team A",
		AttackNumber: 3,
	}
	defender := core_helper_test.PlayerDummy{
		Name:         "Team B",
		DefenceSlice: []int{1, 2, 5, 6},
	}
	playerChannel := make(chan player.Player, 2)
	attacker.RegisterChampionship(playerChannel)
	defender.RegisterChampionship(playerChannel)
	resultsChannel := make(chan []result.Result, 1)
	r.StartRounds(playerChannel, resultsChannel)
	results := <-resultsChannel
	if len(results) != 3 {
		t.Errorf("Should return 3 results")
	}
}
