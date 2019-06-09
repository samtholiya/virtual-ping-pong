package util

import (
	"math/rand"
)

type utility struct {
	limit int
}

func (u utility) GetRandomNumber() int {
	return rand.Intn(u.limit) + 1
}

func (u utility) GetRandomSlice(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(u.limit) + 1
	}
	return arr
}
