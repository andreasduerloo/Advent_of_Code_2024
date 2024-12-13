package day_13

import (
	"advent/helpers"
	"strings"
)

type point struct {
	x int
	y int
}

type clawMachine struct {
	prize point
	a     func(point) point
	b     func(point) point
}

func parse(s string) []clawMachine {
	out := make([]clawMachine, 0)

	blocks := strings.Split(s, "\n\n")

	for _, b := range blocks {
		lines := strings.Split(strings.TrimSpace(b), "\n")

		intA := helpers.ReGetInts(lines[0])
		intB := helpers.ReGetInts(lines[1])
		intPrize := helpers.ReGetInts(lines[2])

		prize := point{intPrize[0], intPrize[1]}

		a := func(p point) point {
			return point{p.x + intA[0], p.y + intA[1]}
		}

		b := func(p point) point {
			return point{p.x + intB[0], p.y + intB[1]}
		}

		out = append(out, clawMachine{
			prize: prize,
			a:     a,
			b:     b,
		})
	}

	return out
}

// Dijkstra can solve this
func solveMachine(m clawMachine) int {
	currentPoint := point{0, 0}
	cost := map[point]int{currentPoint: 0}
	candidates := []point{m.a(currentPoint), m.b(currentPoint)}

	cost[m.a(currentPoint)] = cost[currentPoint] + 3
	cost[m.b(currentPoint)] = cost[currentPoint] + 1

	for len(candidates) != 0 {
		// Get the point with the lowest cost
		currentPoint = getLowest(candidates, cost)
		candidates = helpers.Filter(candidates, func(p point) bool { return p != currentPoint }) // This is super slow

		// Check the points reachable from here
		nba := m.a(currentPoint) // Three tokens
		nbb := m.b(currentPoint) // One token

		// For each neighbor, check:
		// - whether it is in bounds
		// - whether we have been here (if not, it becomes a candidate)
		// - whether we have reached it in a cheaper way now

		if nba.x <= m.prize.x && nba.y <= m.prize.y { // This point is in bounds
			if dist, visited := cost[nba]; visited { // And we've been here
				if cost[currentPoint]+3 < dist { // But we've found a cheaper route
					cost[nba] = cost[currentPoint] + 3
				}
			} else { // We haven't seen that point
				cost[nba] = cost[currentPoint] + 3
				candidates = append(candidates, nba)
			}
		}

		if nbb.x <= m.prize.x && nbb.y <= m.prize.y { // This point is in bounds
			if dist, visited := cost[nbb]; visited { // And we've been here
				if cost[currentPoint]+1 < dist { // But we've found a cheaper route
					cost[nbb] = cost[currentPoint] + 1
				}
			} else { // We haven't seen that point
				cost[nbb] = cost[currentPoint] + 1
				candidates = append(candidates, nbb)
			}
		}
	}

	// We're out of the loop, so we have found all reachable points that are in bounds
	if value, present := cost[m.prize]; present {
		return value
	} else {
		return 0
	}
}

func getLowest(candidates []point, cost map[point]int) point { // We'll have to replace this with a heap
	var min int
	var out point

	for _, p := range candidates {
		if min == 0 {
			min = cost[p]
			out = p
		} else {
			if cost[p] < min {
				min = cost[p]
				out = p
			}
		}
	}

	return out
}
