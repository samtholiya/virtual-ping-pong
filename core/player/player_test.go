package player

import (
	"fmt"
	"testing"
)

func TestPlayerName(t *testing.T) {

	playerName := "Test Name"
	p := NewPlayer("Name")
	p.SetName(playerName)
	if p.GetName() != playerName {
		t.Errorf("Error setting player name")
	}
	if fmt.Sprintf("%v", p) != playerName {
		t.Errorf("Player name should be in string conversion")
	}
}

func TestPlayerScore(t *testing.T) {
	p := NewPlayer("Name")
	p.SetScore(8)
	if p.GetScore() != 8 {
		t.Errorf("Error in setting Player Score")
	}
	p.IncrementScore()
	if p.GetScore() != 9 {
		t.Errorf("Error incrementing Player Sore")
	}
}

func TestPlayerDefenceSize(t *testing.T) {
	p := NewPlayer("Name")
	p.SetDefenceSize(8)
	if len(p.Defend()) != 8 {
		t.Errorf("Error in defence array creation size")
	}
}

func TestPlayerRegisteration(t *testing.T) {
	p := NewPlayer("Name")
	playerChannel := make(chan Player, 0)
	go p.RegisterChampionship(playerChannel)
	rec := <-playerChannel
	if rec != p {
		t.Errorf("Registeration not working correctly")
	}
}
