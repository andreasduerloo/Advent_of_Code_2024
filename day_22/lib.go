package day_22

import (
	"advent/helpers"
	"strings"
)

func parse(s string) []int {
	out := make([]int, 0)

	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		ints := helpers.ReGetInts(line)
		out = append(out, ints[0])
	}

	return out
}

func mix(p, s int) int {
	return p ^ s
}

func prune(s int) int {
	return s % 16_777_216
}

func next(s int) int {
	s = prune(mix(s*64, s))
	s = prune(mix(s/32, s))
	s = prune(mix(s*2048, s))

	return s
}
