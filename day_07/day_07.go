package day_07

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(7)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	lines := parse(inStr)

	// First star
	first := helpers.Sum(helpers.Map(helpers.Filter(lines, fix), func(s []int) int { // TO DO: all of this could be done in one loop with a MapReduce function
		return s[0]
	}))

	// Second star
	second := helpers.Sum(helpers.Map(helpers.Filter(lines, fixConcat), func(s []int) int {
		return s[0]
	}))

	return first, second
}
