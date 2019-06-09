package game

import (
	"errors"

	"github.com/samtholiya/virtual-ping-pong/core/player"
	"github.com/samtholiya/virtual-ping-pong/core/result"
)

type game struct {
	attacker     player.Player
	defender     player.Player
	winningScore int
	result       []result.Result
	level        int
}

func (g *game) SetWinningScore(score int) error {
	if score < 1 {
		return errors.New("Winning score should be at least 1")
	}
	g.winningScore = score
	return nil
}

func (g *game) SetLevel(level int) error {
	if level < 0 {
		return errors.New("Level can not be negative")
	}
	g.level = level
	return nil
}

func (g *game) AddAttacker(player player.Player) error {
	if player == nil {
		return errors.New("Attacker can not be nil")
	}
	player.SetScore(0)
	g.attacker = player
	return nil
}

func (g *game) AddDefender(player player.Player) error {
	if player == nil {
		return errors.New("Defender can not be nil")
	}
	player.SetScore(0)
	g.defender = player
	return nil
}

//Play ... Play the game
func (g *game) Play(resultsChannel chan []result.Result, winnerPlayer chan player.Player) {
	results := make([]result.Result, 0)
	// fmt.Println(g.attacker.GetName() + " is Attacker")
	// fmt.Println(g.defender.GetName() + " is Defender")
	roundLevel := 1
	for g.attacker.GetScore() != g.winningScore && g.defender.GetScore() != g.winningScore {
		result := result.NewResult()
		result.SetRoundLevel(roundLevel)
		result.SetMatchLevel(g.level)
		g.round(result)
		roundLevel++
		results = append(results, result)
	}
	resultsChannel <- results
	if g.winningScore == g.attacker.GetScore() {
		winnerPlayer <- g.attacker
	} else {
		winnerPlayer <- g.defender
	}
}

func (g *game) round(result result.Result) {
	attack := g.attacker.Attack()
	defence := g.defender.Defend()
	defenderWins := false
	for _, element := range defence {
		if attack == element {
			defenderWins = true
		}
	}
	if defenderWins {
		g.defender.IncrementScore()
		result.SetAttacker(g.attacker)
		result.SetDefender(g.defender)
		result.SetWinner(g.defender)
		g.defender, g.attacker = g.attacker, g.defender
	} else {
		g.attacker.IncrementScore()
		result.SetAttacker(g.attacker)
		result.SetDefender(g.defender)
		result.SetWinner(g.attacker)
	}
}
