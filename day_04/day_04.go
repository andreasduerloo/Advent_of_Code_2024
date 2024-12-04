package day_04

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(4)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	first := horizontal(inStr) + vertical(inStr) + downRight(inStr) + downLeft(inStr)
	second := scanForAs(inStr)

	return first, second
}
