package util

//Utility ... Game tools
type Utility interface {

	//GetRandomNumber ... Get a random number
	GetRandomNumber() int

	//GetRandomSlice ... Get a random slice
	GetRandomSlice(int) []int
}
