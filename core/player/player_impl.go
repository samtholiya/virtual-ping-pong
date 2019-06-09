package player

import (
	"errors"

	"github.com/samtholiya/virtual-ping-pong/core/util"
)

type player struct {
	name        string
	score       int
	defenceSize int
}

func (p player) GetName() string {
	return p.name
}

func (p *player) SetName(name string) error {
	if name == "" {
		return errors.New("Name can not be empty")
	}
	p.name = name
	return nil
}

func (p player) GetScore() int {
	return p.score
}

func (p *player) SetScore(score int) error {
	if score < 0 {
		return errors.New("Score can not be negative")
	}
	p.score = score
	return nil
}

func (p *player) IncrementScore() {
	p.score++
}

func (p *player) SetDefenceSize(defenceSize int) error {
	if defenceSize < 1 {
		return errors.New("Defence Size can not be negative")
	}
	p.defenceSize = defenceSize
	return nil
}

func (p player) GetDefenceSize() int {
	return p.defenceSize
}

func (p player) Attack() int {
	util := util.GetUtility()
	return util.GetRandomNumber()
}

func (p player) Defend() []int {
	util := util.GetUtility()
	return util.GetRandomSlice(p.GetDefenceSize())
}

func (p *player) RegisterChampionship(playerRegisterationChannel chan Player) {
	playerRegisterationChannel <- p
}

func (p player) String() string {
	return p.name
}
