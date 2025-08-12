package day_14

import (
	"advent/helpers"
	"fmt"
	"strings"
)

type point struct {
	row int
	col int
}

type quadrants struct {
	values [4]int
}

type robot struct {
	position point
	xvel     int
	yvel     int
}

func parse(s string) []*robot {
	out := make([]*robot, 0)

	lines := strings.Split(strings.TrimSpace(s), "\n")
	for _, line := range lines {
		values := helpers.ReGetInts(line)
		out = append(out, &robot{
			position: point{row: values[0], col: values[1]},
			xvel:     values[2],
			yvel:     values[3],
		})
	}

	return out
}

func (r robot) quadrant() int {
	if r.position.row < 50 {
		if r.position.col < 51 {
			return 0
		}
		if r.position.col > 51 {
			return 1
		}
	}

	if r.position.row > 50 {
		if r.position.col < 51 {
			return 2
		}
		if r.position.col > 51 {
			return 3
		}
	}

	return -1 // These guys are in the middle
}

func (q *quadrants) add(i int) {
	if i < 0 {
		return
	}

	q.values[i]++
}

func (q quadrants) calculate() int {
	out := 1

	for _, quad := range q.values {
		out *= quad
	}

	return out
}

func (r *robot) move() {
	// Horizontally
	newXPosition := r.position.row + r.xvel
	if newXPosition < 0 {
		r.position.row = 101 + newXPosition
	} else if newXPosition > 100 {
		r.position.row = newXPosition % 101
	} else {
		r.position.row = newXPosition
	}

	// Vertically
	newYPosition := r.position.col + r.yvel
	if newYPosition < 0 {
		r.position.col = 103 + newYPosition
	} else if newYPosition > 102 {
		r.position.col = newYPosition % 103
	} else {
		r.position.col = newYPosition
	}
}

func drawFrame(robots []*robot, i int) {
	frame := blankFrame(101, 103, " ")

	for _, r := range robots {
		frame[r.position.row][r.position.col] = "#"
	}

	fmt.Println("Frame nb:", i)
	for _, line := range frame {
		fmt.Println(line)
	}
}

func blankFrame(rows, cols int, def string) [][]string {
	out := make([][]string, rows)

	for i := range out {
		out[i] = make([]string, cols)
		for j := range out[i] {
			out[i][j] = def
		}
	}

	return out
}
