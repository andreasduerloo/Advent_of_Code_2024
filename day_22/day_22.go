package day_22

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(22)
	if err != nil {
		fmt.Println("There was an issue getting the input")
		return 0, 0
	}

	secretNums := parse(inStr)

	var first int

	for _, s := range secretNums {
		for i := 0; i < 2000; i++ {
			s = next(s)
		}
		first += s
	}

	return first, 0
}
