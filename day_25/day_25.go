package day_25

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(25)
	if err != nil {
		fmt.Println("There was an issue getting the input")
		return 0, 0
	}

	locks, keys := parse(inStr)

	var first int

	for _, l := range locks {
		for _, k := range keys {
			if fits(l, k) {
				first++
			}
		}
	}

	return first, 0
}
