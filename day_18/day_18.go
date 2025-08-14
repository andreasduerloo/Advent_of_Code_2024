package day_18

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(18)

	if err != nil {
		fmt.Println("There was an issue getting the input")
		return 0, 0
	}

	mem := initMem()
	locations := parse(inStr)

	corruptLocations(mem, locations, 1024)

	first := BFS(mem, point{x: 0, y: 0}, point{x: 70, y: 70})

	return first, 0
}
