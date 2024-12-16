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

	// For the second star we need to keep track of the optimal routes, i.e. every node also stores what was/were the previous node(s) in the shortest route to that node1.
	second := walkBack2(m)
	// fmt.Println(m.layout[m.end].previous)
	// fmt.Println(m.layout[point{13, 2}].cost, m.layout[point{12, 1}].cost)

	return first, second
}
