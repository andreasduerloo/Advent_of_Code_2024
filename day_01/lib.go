package day_01

import (
	"advent/helpers"
	"strings"
)

func toSlices(inStr string) ([]int, []int) {
	var s1, s2 []int

	for _, line := range strings.Split(inStr, "\n") {
		if line != "" {
			ints := helpers.ReGetInts(line)
			s1 = append(s1, ints[0])
			s2 = append(s2, ints[1])
		}
	}

	return s1, s2
}

func accDistance(s1, s2 []int) int {
	var out int

	for i, val := range s1 {
		if val >= s2[i] {
			out += (val - s2[i])
		} else {
			out += (s2[i] - val)
		}
	}

	return out
}

func intCount(s2 []int) map[int]int {
	out := make(map[int]int)

	for _, val := range s2 {
		if out[val] == 0 {
			out[val] = 1
		} else {
			out[val] += 1
		}
	}

	return out
}

func similarity(s1 []int, c map[int]int) int {
	var out int

	for _, val := range s1 {
		out += val * c[val]
	}

	return out
}
