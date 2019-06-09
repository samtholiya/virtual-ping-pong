package championship

import (
	"errors"

	"github.com/samtholiya/virtual-ping-pong/core/game"
	"github.com/samtholiya/virtual-ping-pong/core/player"
	"github.com/samtholiya/virtual-ping-pong/core/refree"
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

type championship struct {
	refree        refree.Refree
	playerChannel chan player.Player
	resultChannel chan []result.Result
	gameCreator   game.Creator
}

func (c *championship) AddRefree(refree refree.Refree) error {
	if refree == nil {
		return errors.New("Refree can not be nil")
	}
	c.refree = refree
	return nil
}

func (c *championship) SetGameCreator(gameCreator game.Creator) error {
	if gameCreator == nil {
		return errors.New("game config can not be nil")
	}
	c.gameCreator = gameCreator
	return nil
}

func (c *championship) GetRegisterationChannel() chan player.Player {
	if c.playerChannel == nil {
		c.playerChannel = make(chan player.Player)
	}
	return c.playerChannel
}

func (c *championship) StartTournament(result chan []result.Result) error {
	c.resultChannel = result
	if c.refree == nil {
		return errors.New("You dont have a refree for this championship")
	}
	if c.gameCreator == nil && c.refree.GetGameCreator() == nil {
		return errors.New("No game creator was set")
	}
	if c.refree.GetGameCreator() == nil {
		c.refree.SetGameCreator(c.gameCreator)
	}
	go c.refree.StartRounds(c.playerChannel, c.resultChannel)
	return nil
}

func (c championship) GetResultsChannel() chan []result.Result {
	return c.resultChannel
}
