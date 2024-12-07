package day_07

import (
	"advent/helpers"
	"fmt"
	"strconv"
	"strings"
)

func fix(values []int) bool {
	return tryPlus(values, false) || tryTimes(values, false)
}

func fixConcat(values []int) bool {
	return tryPlus(values, true) || tryTimes(values, true) || tryConcat(values)
}

func tryPlus(values []int, concat bool) bool {
	if len(values) >= 4 {
		if values[1]+values[2] > values[0] {
			return false
		}
		newValues := []int{values[0], values[1] + values[2]}
		newValues = append(newValues, values[3:]...)

		if concat {
			return fixConcat(newValues)
		} else {
			return fix(newValues)
		}
	}

	if len(values) == 3 {
		return values[1]+values[2] == values[0]
	}
	fmt.Println("We made it too far in the adding function")
	return false // Should never happen
}

func tryTimes(values []int, concat bool) bool {
	if len(values) >= 4 {
		if values[1]*values[2] > values[0] {
			return false
		}
		newValues := []int{values[0], values[1] * values[2]}
		newValues = append(newValues, values[3:]...)

		if concat {
			return fixConcat(newValues)
		} else {
			return fix(newValues)
		}
	}

	if len(values) == 3 {
		return values[1]*values[2] == values[0]
	}
	fmt.Println("We made it too far in the multiplying function")
	return false // Should never happen
}

func parse(s string) [][]int {
	out := make([][]int, 0)

	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		out = append(out, helpers.ReGetInts(line))
	}

	return out
}

func tryConcat(values []int) bool {
	if len(values) >= 4 {
		if concatInts(values[1], values[2]) > values[0] {
			return false
		}
		newValues := []int{values[0], concatInts(values[1], values[2])}
		newValues = append(newValues, values[3:]...)
		return fixConcat(newValues)
	}

	if len(values) == 3 {
		return concatInts(values[1], values[2]) == values[0]
	}
	fmt.Println("We made it too far in the concatenation function")
	return false // Should never happen
}

func concatInts(a, b int) int {
	as := strconv.Itoa(a)
	bs := strconv.Itoa(b)

	cs := as + bs

	out, err := strconv.Atoi(cs)
	if err != nil {
		fmt.Println("Something went wrong concatenating two numbers")
		return 0
	}

	return out
}
