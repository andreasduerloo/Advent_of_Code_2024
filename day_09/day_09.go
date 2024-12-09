package day_09

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(9)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	massiveSlice := parse(inStr)
	defragmented := defragment(massiveSlice)
	first := checkSum(defragmented)

	filesMoved := moveFiles(massiveSlice)
	second := checkSum(filesMoved)

	return first, second
}
