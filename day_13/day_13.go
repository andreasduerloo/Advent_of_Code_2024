package day_13

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(13)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	machines := parse(inStr)

	var first int

	for _, m := range machines {
		if m.solvable() {
			// fmt.Println(m.solve())
			first += m.solve()
		}
	}

	var second int

	for _, m := range machines {
		m.prizex = m.prizex + 10000000000000
		m.prizey = m.prizey + 10000000000000

		if m.solvable() {
			second += m.solve()
		}

	}

	return first, second
}
