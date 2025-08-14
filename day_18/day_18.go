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

	// Second star, add blocks one by one
	var steps int
	try := 1024

	for steps >= 0 {
		resetVisits(mem)
		corruptLocation(mem, locations[try])

		steps = BFS(mem, point{x: 0, y: 0}, point{x: 70, y: 70})
		if steps >= 0 {
			try++
		}
	}

	second := locations[try]

	return first, second
}
