package day_23

import (
	"slices"
	"strings"
)

type computer struct {
	name      string
	first     rune
	neighbors []string
}

func parse(s string) map[string]*computer {
	out := make(map[string]*computer)

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for _, line := range lines {
		comps := strings.Split(line, "-")

		if c1, exists := out[comps[0]]; exists {
			if c2, exists := out[comps[1]]; exists {
				c1.neighbors = append(c1.neighbors, c2.name)
				c2.neighbors = append(c2.neighbors, c1.name)
			} else {
				c2 := &computer{
					name:      comps[1],
					first:     rune(comps[1][0]),
					neighbors: []string{c1.name},
				}
				out[c2.name] = c2
				c1.neighbors = append(c1.neighbors, c2.name)
			}
		} else {
			if c2, exists := out[comps[1]]; exists {
				c1 := &computer{
					name:      comps[0],
					first:     rune(comps[0][0]),
					neighbors: []string{c2.name},
				}
				out[c1.name] = c1
				c2.neighbors = append(c2.neighbors, c1.name)
			} else {
				c1 := &computer{
					name:      comps[0],
					first:     rune(comps[0][0]),
					neighbors: []string{comps[1]},
				}
				c2 := &computer{
					name:      comps[1],
					first:     rune(comps[1][0]),
					neighbors: []string{comps[0]},
				}
				out[c1.name] = c1
				out[c2.name] = c2
			}
		}
	}

	return out
}

func scan(c *computer, m map[string]*computer) []string {
	out := make([]string, 0)

	if c.first == 't' {
		for i, nb1 := range c.neighbors[:len(c.neighbors)-1] {
			for _, nb2 := range c.neighbors[i+1:] {
				if slices.Contains(m[nb1].neighbors, nb2) {
					setSlice := []string{c.name, nb1, nb2}
					slices.Sort(setSlice)
					out = append(out, toString(setSlice))
				}
			}
		}
	}

	return out
}

func toString(ss []string) string {
	out := ""

	for _, elem := range ss {
		out += elem
	}

	return out
}

func unique(s []string) []string {
	out := make([]string, 0)
	seen := make(map[string]struct{})

	for _, elem := range s {
		if _, present := seen[elem]; present {
			continue
		} else {
			out = append(out, elem)
			seen[elem] = struct{}{}
		}
	}

	return out
}
