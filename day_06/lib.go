package day_06

import (
	"strings"
)

type point struct {
	x int
	y int
}

type state struct {
	location  point
	direction int
}

type guard struct {
	location  point
	direction int
	unique    int
	path      []point
	states    map[state]struct{}
	inBounds  bool
	cycle     bool
}

type board struct {
	obstacles map[point]bool
	height    int
	width     int
}

func (g *guard) step(b board) {
	switch g.direction {
	case UP:
		if b.obstacles[point{g.location.x, g.location.y - 1}] {
			g.turn()
		} else {
			if g.location.y-1 >= 0 {
				g.location = point{g.location.x, g.location.y - 1}   // Move the guard
				g.path = append(g.path, g.location)                  // Add this point to the path
				if _, present := b.obstacles[g.location]; !present { // Have we been here?
					b.obstacles[g.location] = false // We have not, add it
					g.unique++
				}
			} else {
				g.inBounds = false
			}
		}
	case RIGHT:
		if b.obstacles[point{g.location.x + 1, g.location.y}] {
			g.turn()
		} else {
			if g.location.x+1 < b.width {
				g.location = point{g.location.x + 1, g.location.y}   // Move the guard
				g.path = append(g.path, g.location)                  // Add this point to the path
				if _, present := b.obstacles[g.location]; !present { // Have we been here?
					b.obstacles[g.location] = false // We have not, add it
					g.unique++
				}
			} else {
				g.inBounds = false
			}
		}
	case DOWN:
		if b.obstacles[point{g.location.x, g.location.y + 1}] {
			g.turn()
		} else {
			if g.location.y+1 < b.height {
				g.location = point{g.location.x, g.location.y + 1}   // Move the guard
				g.path = append(g.path, g.location)                  // Add this point to the path
				if _, present := b.obstacles[g.location]; !present { // Have we been here?
					b.obstacles[g.location] = false // We have not, add it
					g.unique++
				}
			} else {
				g.inBounds = false
			}
		}
	case LEFT:
		if b.obstacles[point{g.location.x - 1, g.location.y}] {
			g.turn()
		} else {
			if g.location.x-1 >= 0 {
				g.location = point{g.location.x - 1, g.location.y}   // Move the guard
				g.path = append(g.path, g.location)                  // Add this point to the path
				if _, present := b.obstacles[g.location]; !present { // Have we been here?
					b.obstacles[g.location] = false // We have not, add it
					g.unique++
				}
			} else {
				g.inBounds = false
			}
		}
	}
	if g.inBounds {
		// Have we had this state before?
		g.cycle = seenState(state{g.location, g.direction}, g.states)

		// Add the state
		g.states[state{g.location, g.direction}] = struct{}{}
	}
}

func (g *guard) turn() {
	g.direction = (g.direction + 1) % 4
}

func (g *guard) copy() guard {
	return guard{
		location:  g.location,
		direction: g.direction,
		unique:    0,
		path:      make([]point, 0),
		states:    make(map[state]struct{}),
		inBounds:  true,
		cycle:     false,
	}
}

func (g *guard) reset(start point) {
	g.location = start
	g.direction = 0
	g.states = map[state]struct{}{{g.location, 0}: {}}
	g.inBounds = true
	g.cycle = false
}

func (b *board) copy() board {
	obs := make(map[point]bool)

	for k, v := range b.obstacles {
		obs[k] = v
	}
	return board{
		obstacles: obs,
		height:    b.height,
		width:     b.width,
	}
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func parseMap(s string) (board, guard) {
	m := make(map[point]bool)
	g := guard{}

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for y, line := range lines {
		for x, r := range line {
			switch r {
			case '#':
				m[point{x, y}] = true
			case '^':
				g.location = point{x, y}
				g.direction = 0
				g.unique = 1
				g.path = []point{{x, y}}
				g.states = map[state]struct{}{{point{x, y}, 0}: {}}
				g.inBounds = true
				g.cycle = false

				m[point{x, y}] = false
			}
		}
	}

	return board{obstacles: m, height: len(lines), width: len(lines[0])}, g
}

func seenState(s state, allStates map[state]struct{}) bool {
	_, seen := allStates[s]
	return seen
}
