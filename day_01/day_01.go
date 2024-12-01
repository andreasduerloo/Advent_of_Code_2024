package day_01

import (
	"advent/helpers"
	"fmt"
	"slices"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(1)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	s1, s2 := toSlices(inStr)

	// First star
	slices.Sort(s1)
	slices.Sort(s2)

	first := accDistance(s1, s2)

	// Second star
	counts := intCount(s2)
	second := similarity(s1, counts)

	return first, second
}
