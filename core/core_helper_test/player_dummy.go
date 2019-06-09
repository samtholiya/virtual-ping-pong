package core_helper_test

import (
	"errors"

	"github.com/samtholiya/virtual-ping-pong/core/player"
)

//PlayerDummy ... Dummy player obj
type PlayerDummy struct {
	Name         string
	Score        int
	defenceSize  int
	AttackNumber int
	DefenceSlice []int
}

func (p PlayerDummy) GetName() string {
	return p.Name
}

func (p *PlayerDummy) SetDefenceSlice(arr []int) {
	p.DefenceSlice = arr
}

func (p *PlayerDummy) SetAttackNumber(num int) {
	p.AttackNumber = num
}

func (p *PlayerDummy) SetName(Name string) error {
	if Name == "" {
		return errors.New("Name can not be empty")
	}
	p.Name = Name
	return nil
}

func (p PlayerDummy) GetScore() int {
	return p.Score
}

func (p *PlayerDummy) SetScore(Score int) error {
	if Score < 0 {
		return errors.New("Score can not be negative")
	}
	p.Score = Score
	return nil
}

func (p *PlayerDummy) IncrementScore() {
	p.Score++
}

func (p *PlayerDummy) SetDefenceSize(defenceSize int) error {
	if defenceSize < 1 {
		return errors.New("Defence Size can not be negative")
	}
	p.defenceSize = defenceSize
	return nil
}

func (p PlayerDummy) GetDefenceSize() int {
	return p.defenceSize
}

func (p PlayerDummy) Attack() int {
	return p.AttackNumber
}

func (p PlayerDummy) Defend() []int {
	return p.DefenceSlice
}

func (p *PlayerDummy) RegisterChampionship(playerRegisterationChannel chan player.Player) {
	playerRegisterationChannel <- p
}

func (p PlayerDummy) String() string {
	return p.Name
}
