package day_15

import "strings"

type point struct {
	row int
	col int
}

type grid struct {
	data map[point]rune // Easier to work with than [][]rune
}

func (g grid) score() int {
	var out int

	for pt, val := range g.data {
		if val == 'O' {
			out += (100*pt.row + pt.col)
		}
	}

	return out
}

type robot struct {
	position           point
	instructions       []rune
	currentInstruction int
}

func (r *robot) nextMove(g grid) bool {
	switch r.instructions[r.currentInstruction] {
	case '^':
		switch g.data[point{r.position.row - 1, r.position.col}] {
		case '#':
			break
		case '.':
			g.data[r.position] = '.'
			g.data[point{r.position.row - 1, r.position.col}] = '@'
			r.position = point{r.position.row - 1, r.position.col}
		case 'O':
			// We have to keep looking in that direction until we either hit a '.' or a '#'
			firstCrate := point{r.position.row - 1, r.position.col}
			crateLine := []point{firstCrate}

			nextPosition := point{firstCrate.row - 1, firstCrate.col}
			for g.data[nextPosition] == 'O' {
				crateLine = append(crateLine, nextPosition)
				nextPosition = point{nextPosition.row - 1, nextPosition.col}
			}

			if g.data[nextPosition] == '#' {
				break
			}
			if g.data[nextPosition] == '.' {
				g.data[r.position] = '.'
				g.data[crateLine[0]] = '@'
				g.data[nextPosition] = 'O'
				r.position = crateLine[0]
			}
		}
	case 'v':
		switch g.data[point{r.position.row + 1, r.position.col}] {
		case '#':
			break
		case '.':
			g.data[r.position] = '.'
			g.data[point{r.position.row + 1, r.position.col}] = '@'
			r.position = point{r.position.row + 1, r.position.col}
		case 'O':
			// We have to keep looking in that direction until we either hit a '.' or a '#'
			firstCrate := point{r.position.row + 1, r.position.col}
			crateLine := []point{firstCrate}

			nextPosition := point{firstCrate.row + 1, firstCrate.col}
			for g.data[nextPosition] == 'O' {
				crateLine = append(crateLine, nextPosition)
				nextPosition = point{nextPosition.row + 1, nextPosition.col}
			}

			if g.data[nextPosition] == '#' {
				break
			}
			if g.data[nextPosition] == '.' {
				g.data[r.position] = '.'
				g.data[crateLine[0]] = '@'
				g.data[nextPosition] = 'O'
				r.position = crateLine[0]
			}
		}
	case '>':
		switch g.data[point{r.position.row, r.position.col + 1}] {
		case '#':
			break
		case '.':
			g.data[r.position] = '.'
			g.data[point{r.position.row, r.position.col + 1}] = '@'
			r.position = point{r.position.row, r.position.col + 1}
		case 'O':
			// We have to keep looking in that direction until we either hit a '.' or a '#'
			firstCrate := point{r.position.row, r.position.col + 1}
			crateLine := []point{firstCrate}

			nextPosition := point{firstCrate.row, firstCrate.col + 1}
			for g.data[nextPosition] == 'O' {
				crateLine = append(crateLine, nextPosition)
				nextPosition = point{nextPosition.row, nextPosition.col + 1}
			}

			if g.data[nextPosition] == '#' {
				break
			}
			if g.data[nextPosition] == '.' {
				g.data[r.position] = '.'
				g.data[crateLine[0]] = '@'
				g.data[nextPosition] = 'O'
				r.position = crateLine[0]
			}
		}
	case '<':
		switch g.data[point{r.position.row, r.position.col - 1}] {
		case '#':
			break
		case '.':
			g.data[r.position] = '.'
			g.data[point{r.position.row, r.position.col - 1}] = '@'
			r.position = point{r.position.row, r.position.col - 1}
		case 'O':
			// We have to keep looking in that direction until we either hit a '.' or a '#'
			firstCrate := point{r.position.row, r.position.col - 1}
			crateLine := []point{firstCrate}

			nextPosition := point{firstCrate.row, firstCrate.col - 1}
			for g.data[nextPosition] == 'O' {
				crateLine = append(crateLine, nextPosition)
				nextPosition = point{nextPosition.row, nextPosition.col - 1}
			}

			if g.data[nextPosition] == '#' {
				break
			}
			if g.data[nextPosition] == '.' {
				g.data[r.position] = '.'
				g.data[crateLine[0]] = '@'
				g.data[nextPosition] = 'O'
				r.position = crateLine[0]
			}
		}
	}

	if r.currentInstruction == len(r.instructions)-1 {
		return false
	} else {
		r.currentInstruction += 1
		return true
	}
}

func parse(s string) (grid, robot) {
	parts := strings.Split(s, "\n\n")
	gridLines := strings.Split(parts[0], "\n")
	instructions := []rune(strings.TrimSpace(parts[1]))

	gridMap := make(map[point]rune)
	var robotStart point

	for row, line := range gridLines {
		for col, r := range line {
			gridMap[point{row: row, col: col}] = r
			if r == '@' {
				robotStart = point{row: row, col: col}
			}
		}
	}

	return grid{
			data: gridMap,
		},

		robot{
			position:           robotStart,
			instructions:       instructions,
			currentInstruction: 0,
		}
}
