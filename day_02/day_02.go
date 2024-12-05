package day_02

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(2)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	reports := parse(inStr)

	first := len(helpers.Filter(reports, safe))
	second := len(helpers.Filter(reports, dampenSafe))

	return first, second
}
