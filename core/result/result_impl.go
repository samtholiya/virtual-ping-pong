package result

import (
	"errors"

	"github.com/samtholiya/virtual-ping-pong/core/player"
)

type result struct {
	attackerName  string
	defenderName  string
	winnerName    string
	attackerScore int
	defenderScore int
	matchLevel    int
	roundLevel    int
}

//GetAttacker ... Attacker of the game
func (r result) GetAttackerName() string {
	return r.attackerName
}

//GetDefender ... Defender of the game
func (r result) GetDefenderName() string {
	return r.defenderName
}

//GetWinner ... Winner of the game
func (r result) GetWinnerName() string {
	return r.winnerName
}

//GetMatchLevel ... Match Lavel in heirarchy
func (r result) GetMatchLevel() int {
	return r.matchLevel
}

//SetAttacker ... Set Attacker
func (r *result) SetAttacker(player player.PlayerInfo) error {
	if player == nil {
		return errors.New("Attacker can not be nil")
	}
	r.attackerName = player.GetName()
	r.attackerScore = player.GetScore()
	return nil
}

//SetDefender ... Set Defender
func (r *result) SetDefender(player player.PlayerInfo) error {
	if player == nil {
		return errors.New("Defender can not be nil")
	}
	r.defenderName = player.GetName()
	r.defenderScore = player.GetScore()
	return nil
}

//SetWinner ... Set Winner
func (r *result) SetWinner(player player.PlayerInfo) error {
	if player == nil {
		return errors.New("Winner can not be nil")
	}
	r.winnerName = player.GetName()
	return nil
}

//SetMatchLevel ... Match Lavel in heirarchy
func (r *result) SetMatchLevel(level int) error {
	if level < 0 {
		return errors.New("Level can not be negative")
	}
	r.matchLevel = level
	return nil
}

//SetRoundLevel ... Level of the round
func (r *result) SetRoundLevel(roundLevel int) error {
	if roundLevel < 1 {
		return errors.New("Round Level can not be less than 1")
	}
	r.roundLevel = roundLevel
	return nil
}

//GetRoundLevel ... Level of the round
func (r result) GetRoundLevel() int {
	return r.roundLevel
}
