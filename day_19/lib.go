package day_19

import (
	"strings"
)

func parse(s string) ([]string, []string) {
	blocks := strings.Split(s, "\n\n")

	towels := strings.Split(blocks[0], ", ")
	patterns := strings.Split(strings.TrimSpace(blocks[1]), "\n")

	return towels, patterns
}

func possible(t []string, p string) bool {
	var out bool

	for _, tow := range t {
		if out {
			break
		}
		if len(p) > len(tow) {
			if p[0:len(tow)] == tow {
				out = out || possible(t, p[len(tow):])
			}
		} else if len(p) == len(tow) {
			if p == tow {
				out = true
			}
		}
	}
	return out
}

// func countPossible(t []string, p string) int { // Oof. Memoization?
// 	var out int

// 	for _, tow := range t {
// 		if len(p) > len(tow) {
// 			if p[0:len(tow)] == tow {
// 				out += countPossible(t, p[len(tow):])
// 			}
// 		} else if len(p) == len(tow) {
// 			if p == tow {
// 				out = 1
// 			}
// 		}
// 	}
// 	return out
// }

func memoPossible() func([]string, string) int {
	mem := make(map[string]int)

	var countPoss func([]string, string) int

	countPoss = func(t []string, p string) int {
		var out int

		for _, tow := range t {
			if poss, present := mem[p]; present {
				return poss
			}
			if len(p) > len(tow) {
				if p[0:len(tow)] == tow {
					subcount := countPoss(t, p[len(tow):])
					out += subcount
					mem[p[len(tow):]] = subcount
				}
			} else if len(p) == len(tow) {
				if p == tow {
					out += 1
				}
			}
		}
		return out
	}

	return countPoss
}
