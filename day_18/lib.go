package day_18

import (
	"advent/helpers"
	"strings"
)

type point struct {
	x int
	y int
}

type memByte struct {
	corrupt bool
	visited bool
	steps   int
}

type memorySpace map[point]*memByte

func initMem() memorySpace {
	out := make(map[point]*memByte)

	for y := 0; y <= 70; y++ {
		for x := 0; x <= 70; x++ {
			out[point{x, y}] = &memByte{corrupt: false, visited: false, steps: 0}
		}
	}

	return out
}

func parse(s string) []point {
	out := make([]point, 0)

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for _, line := range lines {
		coords := helpers.ReGetInts(line)
		out = append(out, point{coords[0], coords[1]})
	}

	return out
}

func corruptLocations(m memorySpace, locs []point, until int) {
	for i, loc := range locs {
		if i < until {
			m[loc].corrupt = true
		} else {
			break
		}
	}
}

func resetVisits(m memorySpace) {
	for _, val := range m {
		if val.visited {
			val.visited = false
		}
	}
}

func corruptLocation(m memorySpace, loc point) {
	m[loc].corrupt = true
}

type queue []point

func enqueue(p point, q queue) queue {
	q = append(q, p)
	return q
}

func dequeue(q queue) (point, queue) {
	// Some boundary checking would make sense here
	return q[0], q[1:]
}

func BFS(m memorySpace, start, end point) int {
	q := make(queue, 0)

	// Start at the, well, start
	currentLocation := start

	// Add the unvisited and uncorrupted neighbors to the queue
	neighbors := []point{
		point{currentLocation.x + 1, currentLocation.y},
		point{currentLocation.x - 1, currentLocation.y},
		point{currentLocation.x, currentLocation.y + 1},
		point{currentLocation.x, currentLocation.y - 1},
	}

	for _, nb := range neighbors {
		if loc, present := m[nb]; present {
			if !loc.corrupt && !loc.visited {
				loc.visited = true
				loc.steps = m[currentLocation].steps + 1
				q = enqueue(nb, q)
			}
		}
	}

	// Now for the loop
	for len(q) != 0 {
		currentLocation, q = dequeue(q)

		if currentLocation.x == 70 && currentLocation.y == 70 {
			return m[currentLocation].steps
		}

		// Add the unvisited and uncorrupted neighbors to the queue
		neighbors := []point{
			point{currentLocation.x + 1, currentLocation.y},
			point{currentLocation.x - 1, currentLocation.y},
			point{currentLocation.x, currentLocation.y + 1},
			point{currentLocation.x, currentLocation.y - 1},
		}

		for _, nb := range neighbors {
			if loc, present := m[nb]; present {
				if !loc.corrupt && !loc.visited {
					loc.visited = true
					loc.steps = m[currentLocation].steps + 1
					q = enqueue(nb, q)
				}
			}
		}
	}

	return -1
}
