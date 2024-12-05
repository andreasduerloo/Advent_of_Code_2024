package day_05

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(5)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	updates := parse(inStr)

	first := addMiddles(updates)
	second := fixAndCount(updates)

	return first, second
}
