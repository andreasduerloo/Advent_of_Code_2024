package day_25

import "strings"

type lock [5]int
type key [5]int

func parse(s string) ([]lock, []key) {
	blocks := strings.Split(strings.TrimSpace(s), "\n\n")

	locks := make([]lock, 0)
	keys := make([]key, 0)

	for _, schem := range blocks {
		if schem[0] == '#' { // This is a lock
			locks = append(locks, newLock(schem))
		}
		if schem[0] == '.' { // This is a key
			keys = append(keys, newKey(schem))
		}
	}

	return locks, keys
}

func newLock(s string) lock {
	lines := strings.Split(s, "\n")
	out := lock{0, 0, 0, 0, 0}

	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				out[col] = row
			}
		}
	}

	return out
}

func newKey(s string) key {
	lines := strings.Split(s, "\n")
	out := key{0, 0, 0, 0, 0}

	lines = reverse(lines)

	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				out[col] = row
			}
		}
	}

	return out
}

func reverse(ss []string) []string {
	out := make([]string, 0)

	for i := len(ss) - 1; i >= 0; i-- {
		out = append(out, ss[i])
	}

	return out
}

func fits(l lock, k key) bool {
	out := true

	for i := 0; i < 5; i++ {
		out = out && l[i]+k[i] <= 5
	}

	return out
}
