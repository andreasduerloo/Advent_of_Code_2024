package day_24

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(24)
	if err != nil {
		fmt.Println("There was an issue getting the input")
		return 0, 0
	}

	wires := parse(inStr)

	first := zVals(wires)

	return first, 0
}
