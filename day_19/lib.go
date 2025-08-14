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

func countPossible(t []string, p string) int { // Oof. Memoization?
	var out int

	for _, tow := range t {
		if len(p) > len(tow) {
			if p[0:len(tow)] == tow {
				out += countPossible(t, p[len(tow):])
			}
		} else if len(p) == len(tow) {
			if p == tow {
				out = 1
			}
		}
	}
	return out
}
