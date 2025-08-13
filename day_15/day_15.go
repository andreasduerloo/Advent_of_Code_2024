package day_15

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(15)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	warehouse, bot := parse(inStr)

	moreMoves := true

	for moreMoves {
		moreMoves = bot.nextMove(warehouse)
	}

	first := warehouse.score()

	return first, 0
}
