package player

//PlayerInfo ... player info his name, his score
type PlayerInfo interface {

	//GetName ... Name of the player
	GetName() string

	//GetScore ... Score of the plauer
	GetScore() int
}
