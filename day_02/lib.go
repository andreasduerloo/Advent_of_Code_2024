package day_02

import (
	"advent/helpers"
	"strings"
)

type report []int

func parse(s string) []report {
	var out []report

	for _, line := range strings.Split(s, "\n") {
		if line != "" {
			out = append(out, helpers.ReGetInts(line))
		}
	}

	return out
}

func diff(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

func safe(r report) bool {
	// A report is safe if it is strictly decreasing or strictly increasing
	// A report is safe if adjacent numbers have a difference between 1 and 3 (both included)

	increasing := true
	decreasing := true

	for i, n := range r {
		if i == 0 {
			continue
		} else {
			d := diff(n, r[i-1])
			if d >= 1 && d <= 3 {
				increasing = increasing && (n > r[i-1])
				decreasing = decreasing && (n < r[i-1])
			} else {
				return false
			}
		}
	}
	return increasing || decreasing
}

func trimReport(r report, ignore int) report {
	var out report

	for i, n := range r {
		if i != ignore {
			out = append(out, n)
		}
	}

	return out
}

func dampenSafe(r report) bool {
	// If there's no need to dampen, don't
	if safe(r) {
		return true
	}

	// The report isn't safe - ignore values one by one
	for i := 0; i < len(r); i++ {
		if safe(trimReport(r, i)) {
			return true
		}
	}

	return false
}
