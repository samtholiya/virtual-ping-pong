package game

import (
	"testing"

	"github.com/samtholiya/virtual-ping-pong/core/core_helper_test"
	"github.com/samtholiya/virtual-ping-pong/core/result"

	"github.com/samtholiya/virtual-ping-pong/core/player"
)

func TestGameAttackerWin(t *testing.T) {
	newGame := NewGame()

	aPlayer := &core_helper_test.PlayerDummy{
		Name:         "Test A",
		AttackNumber: 3,
	}
	dPlayer := &core_helper_test.PlayerDummy{
		Name:         "Test D",
		DefenceSlice: []int{1, 2, 4, 5, 6, 7},
	}
	newGame.AddAttacker(aPlayer)
	newGame.AddDefender(dPlayer)
	newGame.SetWinningScore(3)
	winChannel := make(chan player.Player, 1)
	resChannel := make(chan []result.Result, 1)
	go newGame.Play(resChannel, winChannel)
	<-resChannel
	winPlayer := <-winChannel
	if winPlayer.GetName() != "Test A" {
		t.Errorf("Attacker should win")
	}
	if winPlayer.GetScore() != 3 {
		t.Errorf("Score should be 3")
	}
}

func TestGameDefenderWin(t *testing.T) {
	newGame := NewGame()

	aPlayer := &core_helper_test.PlayerDummy{
		Name:         "Test A",
		AttackNumber: 3,
		DefenceSlice: []int{1, 2, 4, 5, 6, 7},
	}
	dPlayer := &core_helper_test.PlayerDummy{
		Name:         "Test D",
		DefenceSlice: []int{1, 2, 3, 5, 6, 7},
		AttackNumber: 2,
	}
	newGame.AddAttacker(aPlayer)
	newGame.AddDefender(dPlayer)
	newGame.SetWinningScore(3)
	winChannel := make(chan player.Player, 1)
	resChannel := make(chan []result.Result, 1)
	go newGame.Play(resChannel, winChannel)
	<-resChannel
	winPlayer := <-winChannel
	if winPlayer.GetName() != "Test D" {
		t.Errorf("Defender should win")
	}
	if winPlayer.GetScore() != 3 {
		t.Errorf("Score should be 3")
	}
}
