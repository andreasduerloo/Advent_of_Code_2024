package day_07

import (
	"advent/helpers"
	"fmt"
	"time"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(7)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	lines := parse(inStr)

	start := time.Now()

	// First star
	/*
		first := helpers.Sum(helpers.Map(helpers.Filter(lines, fix), func(s []int) int {
			return s[0]
		}))
	*/

	first := helpers.MapReduce(lines, 0, func(s []int, acc int) int {
		if fix(s) {
			acc += s[0]
		}
		return acc
	})

	// Second star
	/*
		second := helpers.Sum(helpers.Map(helpers.Filter(lines, fixConcat), func(s []int) int {
			return s[0]
		}))
	*/

	second := helpers.MapReduce(lines, 0, func(s []int, acc int) int {
		if fixConcat(s) {
			acc += s[0]
		}
		return acc
	})

	fmt.Println(time.Since(start))
	return first, second
}
