package player

//NewPlayer ... Create a new player
func NewPlayer(name string) Player {
	newPlayer := &player{}
	err := newPlayer.SetName(name)
	if err != nil {
		panic(err)
	}
	return newPlayer
}
