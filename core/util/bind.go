package util

var util *utility

//GetUtility ... an instance of utility
func GetUtility() Utility {
	if util == nil {
		util = &utility{
			limit: 10,
		}
	}
	return util
}
