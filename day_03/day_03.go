package day_03

import (
	"advent/helpers"
	"fmt"
	"regexp"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(3)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	// First star
	mulRe := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := mulRe.FindAllString(inStr, -1)

	first := reduceRe(matches)

	// Second star
	condRe := regexp.MustCompile(`do\(\)|don't\(\)|mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches2 := condRe.FindAllString(inStr, -1)

	second := reduceRe(matches2)

	return first, second
}
