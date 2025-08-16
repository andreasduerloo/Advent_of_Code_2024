package day_23

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(23)
	if err != nil {
		fmt.Println("There was an issue getting the input")
		return 0, 0
	}

	computers := parse(inStr)

	sets := make([]string, 0)

	for _, c := range computers {
		sets = append(sets, scan(c, computers)...)
	}

	sets = unique(sets)
	first := len(sets)

	return first, 0
}
