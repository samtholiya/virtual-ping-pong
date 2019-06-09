package player

//Player ... Actions A Player can perform
type Player interface {
	PlayerInfo

	//SetName ... Set Name of the player
	SetName(string) error

	//GetName ... Get the Name of the player
	SetScore(int) error

	//IncrementScore ... Increment the current score by one
	IncrementScore()

	//Attack ... attack the oponent by generating a returning a number
	Attack() int

	//Defend ... defend the oponent by generating a defence slice
	Defend() []int

	//SetDefenceSize ... Set the defence slice length capacity of the player
	SetDefenceSize(int) error

	//GetDefenceSize ... Get the defence slice length of the player
	GetDefenceSize() int

	//RegisterChampionship ... Register for championship
	RegisterChampionship(chan Player)
}
