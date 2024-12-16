package day_16

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(16)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	m := parse(inStr)

	first := dijkstra(m)

	return first, 0
}
