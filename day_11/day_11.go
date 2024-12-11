package day_11

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(11)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	rocks := helpers.ReGetInts(inStr)

	// The naive way first, simulate the whole thing
	for i := 0; i < 25; i++ {
		rocks = blink(rocks)
	}

	first := len(rocks)

	// Second star: recursion + memoization
	rocks = helpers.ReGetInts(inStr)

	var second int
	mem := make(map[state]int)

	for _, rock := range rocks {
		second += blinkAndRemember(rock, mem, 0, 75)
	}

	return first, second
}
