package day_08

import "strings"

type point struct {
	x int
	y int
}

func parse(s string) map[rune][]point {
	antennas := make(map[rune][]point)

	for y, line := range strings.Split(strings.TrimSpace(s), "\n") {
		for x, r := range line {
			if r != '.' { // This is an antenna
				if _, present := antennas[r]; !present { // Which we haven't seen before
					antennas[r] = []point{{x: x, y: y}}
				} else { // Which we have seen before
					antennaSlice := antennas[r]
					antennaSlice = append(antennaSlice, point{x: x, y: y})
					antennas[r] = antennaSlice
				}
			}
		}
	}

	return antennas
}

func dimensions(s string) (int, int) {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	height := len(lines)
	width := len(lines[0])

	return height, width
}

func antiNodes(am map[rune][]point, height, width int) map[point]int { // Returns a map where the positions of antinodes are the key, and the number of times that antinode appears is the value
	out := make(map[point]int)

	for _, v := range am {
		scanSlice(v, out, height, width)
	}

	return out
}

func scanSlice(points []point, nm map[point]int, height, width int) {
	for a := 0; a < len(points)-1; a++ {
		for b := a + 1; b < len(points); b++ {
			findNodes(points[a], points[b], nm, height, width)
		}
	}
}

func findNodes(a, b point, nm map[point]int, height, width int) {
	horizontalDistance := diff(a.x, b.x)
	verticalDistance := diff(a.y, b.y)

	node1 := point{0, 0}
	node2 := point{0, 0}

	// Find the co-ordinates of the antiNodes
	if a.x < b.x {
		if a.y < b.y {
			node1 = point{a.x - horizontalDistance, a.y - verticalDistance}
			node2 = point{b.x + horizontalDistance, b.y + verticalDistance}
		} else {
			node1 = point{a.x - horizontalDistance, a.y + verticalDistance}
			node2 = point{b.x + horizontalDistance, b.y - verticalDistance}
		}
	} else {
		if a.y < b.y {
			node1 = point{a.x + horizontalDistance, a.y - verticalDistance}
			node2 = point{b.x - horizontalDistance, b.y + verticalDistance}
		} else {
			node1 = point{a.x + horizontalDistance, a.y + verticalDistance}
			node2 = point{b.x - horizontalDistance, b.y - verticalDistance}
		}
	}

	if (node1.x >= 0 && node1.x < width) && (node1.y >= 0 && node1.y <= height) {
		nm[node1]++
	}

	if (node2.x >= 0 && node2.x < width) && (node2.y >= 0 && node2.y < height) {
		nm[node2]++
	}
}

func diff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func antiNodes2(am map[rune][]point, height, width int) map[point]int { // Returns a map where the positions of antinodes are the key, and the number of times that antinode appears is the value
	out := make(map[point]int)

	for _, v := range am {
		scanSlice2(v, out, height, width)
	}

	return out
}

func scanSlice2(points []point, nm map[point]int, height, width int) {
	for a := 0; a < len(points)-1; a++ {
		for b := a + 1; b < len(points); b++ {
			findNodes2(points[a], points[b], nm, height, width)
		}
	}
}

func findNodes2(a, b point, nm map[point]int, height, width int) {
	// This is a first-degree/linear equation
	rise := b.y - a.y
	run := b.x - a.x

	// Start drawing the line from one point
	// First adding the rise and run
	mult := 0
	for {
		x := a.x + mult*run
		y := a.y + mult*rise

		if (x >= 0 && x < width) && (y >= 0 && y < height) {
			nm[point{x, y}]++
			mult++
		} else {
			break
		}
	}

	// Then subtract the rise and run, ignore the antenna this time
	mult = 1
	for {
		x := a.x - mult*run
		y := a.y - mult*rise

		if (x >= 0 && x < width) && (y >= 0 && y < height) {
			nm[point{x, y}]++
			mult++
		} else {
			break
		}
	}
}
