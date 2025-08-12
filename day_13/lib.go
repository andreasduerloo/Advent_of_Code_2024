package day_13

import (
	"advent/helpers"
	"strings"
)

type machine struct {
	ax     int
	ay     int
	bx     int
	by     int
	prizex int
	prizey int
}

func parse(s string) []machine {
	machines := make([]machine, 0)
	blocks := strings.Split(strings.TrimSpace(s), "\n\n")

	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		if len(lines) == 3 {
			aNums := helpers.ReGetInts(lines[0])
			bNums := helpers.ReGetInts(lines[1])
			prizeNums := helpers.ReGetInts(lines[2])

			machines = append(machines, machine{
				ax:     aNums[0],
				ay:     aNums[1],
				bx:     bNums[0],
				by:     bNums[1],
				prizex: prizeNums[0],
				prizey: prizeNums[1],
			})
		}
	}

	return machines
}

func (m machine) solve() int {
	var a, b int

	// Find B first
	b = (m.prizey*m.ax - m.ay*m.prizex) / (m.ax*m.by - m.bx*m.ay)

	// Then A, using B
	a = (m.prizex - m.bx*b) / m.ax

	//if a < 0 || a > 100 || b < 0 || b > 100 {
	if a < 0 || b < 0 {
		return 0
	}

	// Sanity check
	xCoord := a*m.ax + b*m.bx
	yCoord := a*m.ay + b*m.by

	if xCoord != m.prizex || yCoord != m.prizey {
		return 0
	}

	return 3*a + b
}

func (m machine) solvable() bool {
	if m.prizex%GCD(m.ax, m.bx) == 0 && m.prizey%GCD(m.ay, m.by) == 0 {
		return true
	}

	return false
}

func GCD(a, b int) int {
	var bigger, smaller int

	if a == b {
		return a
	}

	if a > b {
		bigger = a
		smaller = b
	} else {
		bigger = b
		smaller = a
	}

	if smaller == 0 {
		return bigger
	}

	return GCD(smaller, bigger%smaller)
}
