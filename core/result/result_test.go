package result

import (
	"testing"

	"github.com/samtholiya/virtual-ping-pong/core/player"
)

func TestResultMatchLevel(t *testing.T) {
	res := NewResult()
	if err := res.SetMatchLevel(-1); err == nil {
		t.Error("Should throe error when match level is negative")
	}
	if err := res.SetMatchLevel(1); err != nil {
		t.Error("Should not throw error when 1 or more than 1")
	}
}

func TestResultRoundLevel(t *testing.T) {
	res := NewResult()
	if err := res.SetMatchLevel(-1); err == nil {
		t.Error("Should throe error when match level is negative")
	}
	if err := res.SetMatchLevel(1); err != nil {
		t.Error("Should not throw error when 1 or more than 1")
	}
}

func TestResultAttacker(t *testing.T) {
	res := NewResult()
	if err := res.SetAttacker(nil); err == nil {
		t.Error("Should throw error when Attacker is nil")
	}
	if err := res.SetAttacker(player.NewPlayer("T")); err != nil {
		t.Error("Should not throw error when player is provided")
	}
}

func TestResultDefender(t *testing.T) {
	res := NewResult()
	if err := res.SetDefender(nil); err == nil {
		t.Error("Should throw error when Attacker is nil")
	}
	if err := res.SetDefender(player.NewPlayer("T")); err != nil {
		t.Error("Should not throw error when player is provided")
	}
}

func TestResultWinner(t *testing.T) {
	res := NewResult()
	if err := res.SetWinner(nil); err == nil {
		t.Error("Should throw error when Attacker is nil")
	}
	if err := res.SetWinner(player.NewPlayer("T")); err != nil {
		t.Error("Should not throw error when player is provided")
	}
}
