package day_19

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(19)

	if err != nil {
		fmt.Println("There was an issue getting the input")
		return 0, 0
	}

	towels, patterns := parse(inStr)

	var first, second int

	for _, patt := range patterns {
		if possible(towels, patt) {
			first++
		}
	}

	for _, patt := range patterns {
		second += countPossible(towels, patt)
	}

	return first, second
}
