package setup

import "math/rand"

var r *rand.Rand

func init() {
	var source int64 = 95737
	r = rand.New(rand.NewSource(source))
}

// GetRandLimit returns a random int ranging from 0 to end.
//
// 0 is inclusive, the end value is not.
func GetRandLimit(end int) int {
	return r.Intn(end)
}

// GetRandLimitInt64 returns a random int64 ranging from 0 to end.
//
// 0 is inclusive, the end value is not.
func GetRandLimitInt64(end int) int64 {
	return int64(GetRandLimit(end))
}

// GetRandRange returns a random int ranging from start to end.
//
// The start value is inclusive, the end value is not.
func GetRandRange(start, end int) int {
	return r.Intn(end-start) + start
}

// GetRandRangeInt64 returns a random int64 ranging from start to end.
//
// The start value is inclusive, the end value is not.
func GetRandRangeInt64(start, end int) int64 {
	return int64(GetRandRange(start, end))
}

func GetRandInt64() int64 {
	return r.Int63()
}

// GetRandBoolWeighted returns a bool, the value of which will be chosen randomly based on the provide true factor.
func GetRandBoolWeighted(trueFactor float64) bool {

	limit := 100
	posLimit := int(trueFactor * float64(limit))

	randVal := GetRandLimit(limit)

	return randVal < posLimit
}
