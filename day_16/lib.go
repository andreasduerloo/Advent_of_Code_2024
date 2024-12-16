package day_16

import (
	"advent/helpers"
	"fmt"
	"slices"
	"strings"
)

type point struct {
	x int
	y int
}

const (
	EAST = iota
	SOUTH
	WEST
	NORTH
)

type square struct {
	value    rune
	cost     int
	visited  bool
	facing   int
	previous []point
}

type maze struct {
	layout map[point]square
	height int
	width  int
	start  point
	end    point
}

func parse(s string) maze {
	out := make(map[point]square)
	lines := strings.Split(strings.TrimSpace(s), "\n")
	var start, end point

	for y, line := range lines {
		for x, r := range line {
			out[point{x, y}] = square{value: r}
			if r == 'S' {
				start = point{x, y}
			}
			if r == 'E' {
				end = point{x, y}
			}
		}
	}

	return maze{
		layout: out,
		height: len(lines),
		width:  len(lines[0]),
		start:  start,
		end:    end,
	}
}

func dijkstra(m maze) int {
	currentPoint := m.start
	localc := m.layout[currentPoint]
	localc.visited = true
	m.layout[currentPoint] = localc
	candidates := neighbors(currentPoint, m)

	for _, nb := range candidates {
		localnb := m.layout[nb]
		localnb.previous = []point{currentPoint}
		switch m.layout[currentPoint].facing {
		case EAST:
			if nb.x == currentPoint.x+1 {
				localnb.cost = m.layout[currentPoint].cost + 1
				localnb.facing = EAST
			} else if nb.y == currentPoint.y-1 {
				localnb.cost = m.layout[currentPoint].cost + 1001
				localnb.facing = NORTH
			} else if nb.y == currentPoint.y+1 {
				localnb.cost = m.layout[currentPoint].cost + 1001
				localnb.facing = SOUTH
			}
		case SOUTH:
			if nb.y == currentPoint.y+1 {
				localnb.cost = m.layout[currentPoint].cost + 1
				localnb.facing = SOUTH
			} else if nb.x == currentPoint.x-1 {
				localnb.cost = m.layout[currentPoint].cost + 1001
				localnb.facing = WEST
			} else if nb.x == currentPoint.x+1 {
				localnb.cost = m.layout[currentPoint].cost + 1001
				localnb.facing = EAST
			}
		case WEST:
			if nb.x == currentPoint.x-1 {
				localnb.cost = m.layout[currentPoint].cost + 1
				localnb.facing = WEST
			} else if nb.y == currentPoint.y-1 {
				localnb.cost = m.layout[currentPoint].cost + 1001
				localnb.facing = NORTH
			} else if nb.y == currentPoint.y+1 {
				localnb.cost = m.layout[currentPoint].cost + 1001
				localnb.facing = SOUTH
			}
		case NORTH:
			if nb.y == currentPoint.y-1 {
				localnb.cost = m.layout[currentPoint].cost + 1
				localnb.facing = NORTH
			} else if nb.x == currentPoint.x-1 {
				localnb.cost = m.layout[currentPoint].cost + 1001
				localnb.facing = WEST
			} else if nb.x == currentPoint.x+1 {
				localnb.cost = m.layout[currentPoint].cost + 1001
				localnb.facing = EAST
			}
		}
		m.layout[nb] = localnb
	}

	for len(candidates) != 0 {
		currentPoint = getLowestCost(candidates, m)
		localc := m.layout[currentPoint]
		localc.visited = true
		m.layout[currentPoint] = localc
		candidates = helpers.Filter(candidates, func(p point) bool { return p != currentPoint }) // This is super slow

		nbs := neighbors(currentPoint, m)
		// For each neighbor, check:
		// - whether we have been here (if not, it becomes a candidate)
		// - whether we have reached it in a cheaper way now
		for _, nb := range nbs {
			localnb := m.layout[nb]

			switch m.layout[currentPoint].facing {
			case EAST:
				if nb.x == currentPoint.x+1 {
					localnb.cost = m.layout[currentPoint].cost + 1
					localnb.facing = EAST
				} else if nb.y == currentPoint.y-1 {
					localnb.cost = m.layout[currentPoint].cost + 1001
					localnb.facing = NORTH
				} else if nb.y == currentPoint.y+1 {
					localnb.cost = m.layout[currentPoint].cost + 1001
					localnb.facing = SOUTH
				}
			case SOUTH:
				if nb.y == currentPoint.y+1 {
					localnb.cost = m.layout[currentPoint].cost + 1
					localnb.facing = SOUTH
				} else if nb.x == currentPoint.x-1 {
					localnb.cost = m.layout[currentPoint].cost + 1001
					localnb.facing = WEST
				} else if nb.x == currentPoint.x+1 {
					localnb.cost = m.layout[currentPoint].cost + 1001
					localnb.facing = EAST
				}
			case WEST:
				if nb.x == currentPoint.x-1 {
					localnb.cost = m.layout[currentPoint].cost + 1
					localnb.facing = WEST
				} else if nb.y == currentPoint.y-1 {
					localnb.cost = m.layout[currentPoint].cost + 1001
					localnb.facing = NORTH
				} else if nb.y == currentPoint.y+1 {
					localnb.cost = m.layout[currentPoint].cost + 1001
					localnb.facing = SOUTH
				}
			case NORTH:
				if nb.y == currentPoint.y-1 {
					localnb.cost = m.layout[currentPoint].cost + 1
					localnb.facing = NORTH
				} else if nb.x == currentPoint.x-1 {
					localnb.cost = m.layout[currentPoint].cost + 1001
					localnb.facing = WEST
				} else if nb.x == currentPoint.x+1 {
					localnb.cost = m.layout[currentPoint].cost + 1001
					localnb.facing = EAST
				}
			}

			if m.layout[nb].visited {
				if m.layout[nb].cost > localnb.cost { // We found a shorter way
					localnb.previous = []point{currentPoint}
					m.layout[nb] = localnb
				} else if m.layout[nb].cost == localnb.cost {
					if m.layout[nb].value == 'E' {
						fmt.Println("Here, somehow", m.layout[nb].cost, localnb.cost)
					}
					localnb.previous = append(localnb.previous, currentPoint)
					m.layout[nb] = localnb
				}
			} else {
				localnb.previous = []point{currentPoint}
				m.layout[nb] = localnb
				candidates = append(candidates, nb)
			}
		}
	}

	return m.layout[m.end].cost
}

func neighbors(p point, m maze) []point {
	out := make([]point, 0)

	if m.layout[point{p.x, p.y + 1}].value != '#' {
		out = append(out, point{p.x, p.y + 1})
	}
	if m.layout[point{p.x, p.y - 1}].value != '#' {
		out = append(out, point{p.x, p.y - 1})
	}
	if m.layout[point{p.x + 1, p.y}].value != '#' {
		out = append(out, point{p.x + 1, p.y})
	}
	if m.layout[point{p.x - 1, p.y}].value != '#' {
		out = append(out, point{p.x - 1, p.y})
	}

	return out
}

func getLowestCost(candidates []point, m maze) point { // We'll have to replace this with a heap
	var min int
	var out point

	for _, p := range candidates {
		if min == 0 {
			min = m.layout[p].cost
			out = p
		} else {
			if m.layout[p].cost < min {
				min = m.layout[p].cost
				out = p
			}
		}
	}

	return out
}

func walkBack(m maze) int {
	currentPoint := m.end
	visited := []point{currentPoint}
	toCheck := m.layout[currentPoint].previous

	for len(toCheck) != 0 {
		currentPoint = toCheck[0]
		if !slices.Contains(visited, currentPoint) {
			visited = append(visited, currentPoint)
		}
		toCheck = helpers.Filter(toCheck, func(p point) bool { return p != currentPoint })

		// Are any of the previous points not yet visited?
		for _, p := range m.layout[currentPoint].previous {
			if !slices.Contains(visited, p) && !slices.Contains(toCheck, p) { // We haven't visited that point yet
				toCheck = append(toCheck, p)
			}
		}

	}

	return len(visited)
}

func walkBack2(m maze) int {
	currentPoint := m.end

	// Use a map as a set for efficient duplicate prevention
	visited := make(map[point]bool)
	toCheck := []point{currentPoint}

	for len(toCheck) > 0 {
		// Dequeue the first point
		currentPoint = toCheck[0]
		toCheck = toCheck[1:] // Remove the first element

		// Skip if we've already visited this point
		if visited[currentPoint] {
			continue
		}

		// Mark the current point as visited
		visited[currentPoint] = true

		// Add all previous nodes to the queue if they haven't been visited
		for _, p := range m.layout[currentPoint].previous {
			if !visited[p] && !slices.Contains(toCheck, p) {
				toCheck = append(toCheck, p)
			}
		}
	}

	return len(visited)
}
