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
	results := helpers.Filter(lines, fix)
	values := helpers.Map(results, func(s []int) int {
		return s[0]
	})
	first := helpers.Sum(values)

	// Second star
	results = helpers.Filter(lines, fixConcat)
	values = helpers.Map(results, func(s []int) int {
		return s[0]
	})
	second := helpers.Sum(values)

	return first, second
}
