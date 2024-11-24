package day_01

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	testStr, err := helpers.GetInput(1)
	if err != nil {
		fmt.Println("Oh noes")
	}
	fmt.Println(testStr)

	return 0, 0
}
