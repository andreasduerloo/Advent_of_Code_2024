package day_12

import (
	"slices"
	"strings"
)

type square struct {
	value rune
	field int
}

type grid struct {
	squares map[point]square
	height  int
	width   int
}

type field struct {
	id      int
	size    int
	borders int
	squares []point
}

type point struct {
	x int
	y int
}

// Then loop through them and multiply

func parse(s string) grid {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	squares := make(map[point]square)

	for y, line := range lines {
		for x, r := range line {
			squares[point{x, y}] = square{
				value: r,
			}
		}
	}

	return grid{
		squares: squares,
		height:  len(lines),
		width:   len(lines[0]),
	}
}

func buildField(l point, id int, layout grid, fields map[int]field) {
	queue := neighbors(l, layout) // BFS
	var nb point

	startingSquare := layout.squares[l]
	startingSquare.field = id
	layout.squares[l] = startingSquare

	fields[id] = field{
		id:      id,
		size:    1,
		borders: 4,
		squares: []point{l},
	}

	for len(queue) > 0 {
		nb, queue = dequeue(queue)
		if layout.squares[nb].value == layout.squares[l].value {
			// Add the neighbors not yet in the field to the queue
			nbs := neighbors(nb, layout)
			for _, nnb := range nbs {
				if layout.squares[nnb].field == 0 && !slices.Contains(queue, nnb) {
					queue = append(queue, nnb)
				}
			}

			// Join the field with this square
			thisSquare := layout.squares[nb]
			thisSquare.field = id
			layout.squares[nb] = thisSquare

			// Figure out the new size and borders
			thisField := fields[id]
			thisField.size++
			thisField.squares = append(thisField.squares, nb)
			var alreadyIn int

			for _, nnb := range nbs {
				if layout.squares[nnb].field == id {
					alreadyIn++
				}
			}

			thisField.borders = thisField.borders - (alreadyIn) + (4 - alreadyIn)
			fields[id] = thisField
		}
	}
}

func neighbors(l point, layout grid) []point { // Returns neighbors in the order LEFT, UP, RIGHT, DOWN
	out := make([]point, 0)

	if l.x > 0 {
		out = append(out, point{l.x - 1, l.y})
	}

	if l.y > 0 {
		out = append(out, point{l.x, l.y - 1})
	}

	if l.x < layout.width-1 {
		out = append(out, point{l.x + 1, l.y})
	}

	if l.y < layout.height-1 {
		out = append(out, point{l.x, l.y + 1})
	}

	return out
}

func dequeue(s []point) (point, []point) {
	return s[0], s[1:]
}

func calcSides(shape []point) int {
	var out int

	// Find the minimum and maximum x and y values. This is a rectangle around our shape
	minx := shape[0].x
	maxx := shape[0].x
	miny := shape[0].y
	maxy := shape[0].y

	for _, p := range shape {
		if p.x < minx {
			minx = p.x
		}
		if p.x > maxx {
			maxx = p.x
		}
		if p.y < miny {
			miny = p.y
		}
		if p.y > maxy {
			maxy = p.y
		}
	}

	// Vertical scan
	for x := minx; x <= maxx; x++ {
		for y := miny - 1; y <= maxy+1; y++ {
			// fmt.Println("Looking at point", x, y)
			if !slices.Contains(shape, point{x, y}) && slices.Contains(shape, point{x, y + 1}) {
				// We're entering the shape from above.
				// This is a new side, unless point{x - 1, y + 1} IS in the shape, and point{x - 1, y} is NOT.
				if !(slices.Contains(shape, point{x - 1, y + 1}) && !slices.Contains(shape, point{x - 1, y})) {
					// fmt.Println("Adding one because we're entering from above")
					out++
					continue
				}
			}
			if slices.Contains(shape, point{x, y}) && !slices.Contains(shape, point{x, y + 1}) {
				// We're leaving the shape through the bottom.
				// This is a new side, unless point{x - 1, y} IS in the shape, and point{x - 1, y + 1} is NOT.
				if !(slices.Contains(shape, point{x - 1, y}) && !slices.Contains(shape, point{x - 1, y + 1})) {
					// fmt.Println("Adding one because we're leaving through the bottom")
					out++
					continue
				}
			}
		}
	}

	// Horizontal scan
	for y := miny; y <= maxy; y++ {
		for x := minx - 1; x <= maxx+1; x++ {
			// fmt.Println("Looking at point", x, y)
			if !slices.Contains(shape, point{x, y}) && slices.Contains(shape, point{x + 1, y}) {
				// We're entering the shape from the left.
				// This is a new side, unless point{x + 1, y - 1} IS in the shape, and point{x, y - 1} is NOT.
				if !(slices.Contains(shape, point{x + 1, y - 1}) && !slices.Contains(shape, point{x, y - 1})) {
					// fmt.Println("Adding one because we're entering from the left")
					out++
					continue
				}
			}
			if slices.Contains(shape, point{x, y}) && !slices.Contains(shape, point{x + 1, y}) {
				// We're leaving the shape to the right.
				// This is a new side, unless point{x, y - 1} IS in the shape, and point{x + 1, y - 1} is NOT.
				if !(slices.Contains(shape, point{x, y - 1}) && !slices.Contains(shape, point{x + 1, y - 1})) {
					// fmt.Println("Adding one as we're leaving to the right")
					out++
					continue
				}
			}
		}
	}

	return out
}
