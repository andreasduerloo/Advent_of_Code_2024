package day_13

import (
	"advent/helpers"
	"fmt"
)

// TODO
// Swap out the slice of candidates with a min heap
// This saves us iterating through all candidates to get the lowest cost AND
// we won't need to iterate through all the candidates AGAIN to remove the lowest cost candidate

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(13)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	machines := parse(inStr)

	var first int
	for _, m := range machines {
		first += solveMachine(m)
	}

	return first, 0
}
