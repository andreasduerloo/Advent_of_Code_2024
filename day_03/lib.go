package day_03

import "advent/helpers"

func reduceRe(s []string) int {
	var out int
	do := true

	for _, hit := range s {
		switch hit {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if do {
				ints := helpers.ReGetInts(hit)
				out += (ints[0] * ints[1])
			}
		}
	}
	return out
}
