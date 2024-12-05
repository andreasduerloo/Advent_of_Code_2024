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

// TO DO: pass updates (that contain pages) instead of a map ints

type update []int

func parse(s string) (map[int]page, []update) {
	// Separate the rules from the updates
	inputParts := strings.Split(s, "\n\n")
	ruleInput := inputParts[0]
	updatesInput := inputParts[1]

	rules := parseRules(ruleInput)
	updates := parseUpdates(updatesInput)

	return rules, updates
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

func parseUpdates(s string) []update {
	updates := make([]update, 0)

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for _, line := range lines {
		updates = append(updates, helpers.ReGetInts(line))
	}

	return updates
}

func validateUpdate(u update, r map[int]page) int {
	// An update is not valid if a rule is broken
	for p := len(u) - 1; p > 0; p-- {
		for other := p - 1; other >= 0; other-- {
			if slices.Contains(r[u[other]].after, u[p]) || slices.Contains(r[u[p]].before, u[other]) {
				return 0
			}
		}
	}

	// If we made it this far, the update is fine
	return u[len(u)/2]
}

func addMiddles(u []update, r map[int]page) int {
	var out int

	for _, ud := range u {
		out += validateUpdate(ud, r)
	}

	return out
}

func getIncorrect(u []update, r map[int]page) []update {
	out := make([]update, 0)

	for _, ud := range u {
		if validateUpdate(ud, r) == 0 {
			out = append(out, ud)
		}
	}

	return out
}

func fixAndCount(u []update, r map[int]page) int {
	var out int

	toFix := getIncorrect(u, r)

	for _, ud := range toFix {
		out += fix(ud, r)
	}

	return out
}

func fix(u update, r map[int]page) int {
	// Make a slice of updates rather than ints
	pages := make([]page, 0)

	for _, p := range u {
		pages = append(pages, r[p])
	}

	slices.SortFunc(pages, compare)

	return pages[len(pages)/2].number
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
