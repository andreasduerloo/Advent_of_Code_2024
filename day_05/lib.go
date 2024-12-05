package day_05

import (
	"advent/helpers"
	"slices"
	"strings"
)

type page struct {
	number int
	before []int
	after  []int
}

type update []page

func parse(s string) []update {
	inputParts := strings.Split(s, "\n\n")
	ruleInput := inputParts[0]
	updatesInput := inputParts[1]

	rules := parseRules(ruleInput)
	updates := parseUpdates(updatesInput, rules)

	return updates
}

func parseRules(s string) map[int]page {
	rules := make(map[int]page)

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for _, line := range lines {
		values := helpers.ReGetInts(line)

		// Add or update the first page mentioned
		if _, present := rules[values[0]]; !present {
			// This page is not in the map yet - add it
			rules[values[0]] = page{
				number: values[0],
				before: []int{values[1]},
				after:  make([]int, 0),
			}
		} else {
			// This page is known, we need to update the 'before' field
			entry := rules[values[0]]
			entry.before = append(entry.before, values[1])
			rules[values[0]] = entry
		}

		// Add or update the second page mentioned
		if _, present := rules[values[1]]; !present {
			// This page is not in the map yet - add it
			rules[values[1]] = page{
				number: values[1],
				before: make([]int, 0),
				after:  []int{values[0]},
			}
		} else {
			// This page is known, we need to update the 'after' field
			entry := rules[values[1]]
			entry.after = append(entry.after, values[0])
		}
	}

	return rules
}

func parseUpdates(s string, r map[int]page) []update {
	updates := make([]update, 0)

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for _, line := range lines {
		newUpdate := make(update, 0)
		numbers := helpers.ReGetInts(line)

		for _, n := range numbers {
			newUpdate = append(newUpdate, r[n])
		}
		updates = append(updates, newUpdate)
	}

	return updates
}

func validateUpdate(u update) int {
	for p := len(u) - 1; p > 0; p-- {
		for other := p - 1; other >= 0; other-- {
			if slices.Contains(u[other].after, u[p].number) || slices.Contains(u[p].before, u[other].number) {
				return 0
			}
		}
	}
	return u[len(u)/2].number
}

func addMiddles(u []update) int {
	var out int

	for _, ud := range u {
		out += validateUpdate(ud)
	}

	return out
}

func getIncorrect(u []update) []update {
	out := make([]update, 0)

	for _, ud := range u {
		if validateUpdate(ud) == 0 {
			out = append(out, ud)
		}
	}

	return out
}

func fixAndCount(u []update) int {
	var out int

	toFix := getIncorrect(u)

	for _, ud := range toFix {
		slices.SortFunc(ud, compare) // This is where the magic happens
		out += ud[len(ud)/2].number
	}

	return out
}

func compare(a, b page) int {
	// a before b -> return a negative number
	if slices.Contains(a.before, b.number) || slices.Contains(b.after, a.number) {
		return -1
	}

	// b before a -> return a positive number
	if slices.Contains(a.after, b.number) || slices.Contains(b.before, a.number) {
		return 1
	}

	// equal -> return 0
	return 0
}
