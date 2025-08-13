package day_17

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(17)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	c := parse(inStr)

	for !c.halt {
		c.execute()
	}

	return 0, 0
}
