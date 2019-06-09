package util

import (
	"testing"
)

func TestUtilNumber(t *testing.T) {
	utils := GetUtility()
	if utils.GetRandomNumber() <= 0 || utils.GetRandomNumber() > 10 {
		t.Error("Random number should be 1 to 10 only")
	}
}

func TestUtilSlice(t *testing.T) {
	utils := GetUtility()
	if len(utils.GetRandomSlice(8)) != 8 {
		t.Error("Length of slice should be 8")
	}
}
