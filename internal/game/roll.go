package game

import "math/rand"

// doRoll rolls a single die and returns the result.
func doRoll() Berry {
	roll := rand.Float64()
	switch {
	case roll < 0.3:
		return Jumbleberry
	case roll < 0.6:
		return Sugarberry
	case roll < 0.8:
		return Pickleberry
	case roll < 0.9:
		return Moonberry
	default:
		return Pest
	}
}

// doRolls rolls n dice and returns the results in a slice.
func DoRolls(n int)[]Berry {
	result := make([]Berry, n)

	for i := range n {
		result[i] = doRoll()
	}

	return result
}