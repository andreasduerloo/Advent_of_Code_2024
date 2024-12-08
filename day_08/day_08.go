package day_08

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(8)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	antennas := parse(inStr)
	height, width := dimensions(inStr)

	first := len(antiNodes(antennas, height, width))
	second := len(antiNodes2(antennas, height, width))

	return first, second
}
