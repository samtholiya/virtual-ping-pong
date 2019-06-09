package result

import "github.com/samtholiya/virtual-ping-pong/core/player"

//Result ... result of the match
type Result interface {
	//GetAttacker ... Attacker of the game
	GetAttackerName() string

	//GetDefender ... Defender of the game
	GetDefenderName() string

	//GetWinner ... Winner of the game
	GetWinnerName() string

	//GetMatchLevel ... Match Lavel in heirarchy
	GetMatchLevel() int

	//GetRoundLevel ... Level of the round for which the result is
	GetRoundLevel() int

	//SetAttacker ... Set Attacker
	SetAttacker(player.PlayerInfo) error

	//SetDefender ... Set Defender
	SetDefender(player.PlayerInfo) error

	//SetWinner ... Set Winner
	SetWinner(player.PlayerInfo) error

	//SetMatchLevel ... Match Lavel in heirarchy
	SetMatchLevel(int) error

	//SetRoundLevel ... Level of the round
	SetRoundLevel(int) error
}
