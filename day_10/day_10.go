package day_10

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(10)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	topoMap := buildGrid(inStr)

	first := scanGrid(topoMap, false)
	second := scanGrid(topoMap, true)

	return first, second
}
