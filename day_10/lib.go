package day_10

import (
	"slices"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type topoMap struct {
	grid   map[point]int
	height int
	width  int
}

func buildGrid(s string) topoMap {
	grid := make(map[point]int)

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for y, line := range lines {
		for x, r := range line {
			grid[point{x, y}], _ = strconv.Atoi(string(r))
		}
	}

	return topoMap{
		grid:   grid,
		height: len(lines),
		width:  len(lines[0]),
	}
}

func scanGrid(tm topoMap, second bool) int {
	var out int

	for y := 0; y < tm.height; y++ {
		for x := 0; x < tm.width; x++ {
			if tm.grid[point{x, y}] == 0 {
				out += findRoutes(point{x, y}, tm, second)
			}
		}
	}

	return out
}

func findRoutes(trailHead point, tm topoMap, second bool) int { // Depth-first search
	var out int
	currentPoint := trailHead
	stack := []point{currentPoint}
	visited := make([]point, 0)

	for len(stack) != 0 {
		currentPoint = stack[0]
		stack = stack[1:]

		if tm.grid[currentPoint] == 9 {
			if !second {
				if !slices.Contains(visited, currentPoint) {
					out++
					visited = append(visited, currentPoint)
				}
			} else {
				out++
			}
		}

		for _, nb := range neighbors(currentPoint, tm) {
			if tm.grid[nb] == tm.grid[currentPoint]+1 {
				stack = prepend(stack, nb)
			}
		}
	}

	return out
}

func neighbors(p point, tm topoMap) []point {
	out := make([]point, 0)

	if p.x > 0 {
		out = append(out, point{p.x - 1, p.y})
	}

	if p.x < tm.width-1 {
		out = append(out, point{p.x + 1, p.y})
	}

	if p.y > 0 {
		out = append(out, point{p.x, p.y - 1})
	}

	if p.y < tm.height-1 {
		out = append(out, point{p.x, p.y + 1})
	}

	return out
}

func prepend(s []point, val point) []point {
	out := make([]point, len(s)+1)
	out[0] = val
	copy(out[1:], s)

	return out
}
