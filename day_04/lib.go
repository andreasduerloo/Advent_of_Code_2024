package day_04

import "strings"

func horizontal(s string) int {
	var out int

	for _, line := range strings.Split(s, "\n") {
		for i, c := range line {
			switch c {
			case 'X':
				if i <= len(line)-4 {
					// if line[i+1:i+3] == "MAS" { // Does string indexing work this way? Or is it rune-by-run? Either way it won't work vertically
					if line[i+1] == 'M' && line[i+2] == 'A' && line[i+3] == 'S' {
						out += 1
					}
				}
			case 'S':
				if i <= len(line)-4 {
					// if line[i+1:i+3] == "AMX" {
					if line[i+1] == 'A' && line[i+2] == 'M' && line[i+3] == 'X' {
						out += 1
					}
				}
			}
		}
	}
	return out
}

func vertical(s string) int {
	var out int
	split := strings.Split(strings.TrimSpace(s), "\n")

	for col := 0; col < len(split[0]); col++ {
		for line := 0; line < len(split); line++ {
			switch split[line][col] {
			case 'X':
				if line <= len(split)-4 {
					if split[line+1][col] == 'M' && split[line+2][col] == 'A' && split[line+3][col] == 'S' {
						out += 1
					}
				}
			case 'S':
				if line <= len(split)-4 {
					if split[line+1][col] == 'A' && split[line+2][col] == 'M' && split[line+3][col] == 'X' {
						out += 1
					}
				}
			}
		}
	}

	return out
}

func downRight(s string) int {
	var out int

	split := strings.Split(strings.TrimSpace(s), "\n")
	width := len(split[0])
	height := len(split)

	for row := 1; row < height; row++ { // Left edge
		localRow := row
		col := 0

		for localRow < height-1 {
			switch split[localRow][col] {
			case 'X':
				if localRow <= height-4 && col <= width-4 {
					if split[localRow+1][col+1] == 'M' && split[localRow+2][col+2] == 'A' && split[localRow+3][col+3] == 'S' {
						out += 1
					}
				}
			case 'S':
				if localRow <= height-4 && col <= width-4 {
					if split[localRow+1][col+1] == 'A' && split[localRow+2][col+2] == 'M' && split[localRow+3][col+3] == 'X' {
						out += 1
					}
				}
			}
			localRow++
			col++
		}
	}

	for col := 0; col < width; col++ { // Top edge
		row := 0
		localCol := col

		for localCol < width-1 {
			switch split[row][localCol] {
			case 'X':
				if row <= height-4 && localCol <= width-4 {
					if split[row+1][localCol+1] == 'M' && split[row+2][localCol+2] == 'A' && split[row+3][localCol+3] == 'S' {
						out += 1
					}
				}
			case 'S':
				if row <= height-4 && localCol <= width-4 {
					if split[row+1][localCol+1] == 'A' && split[row+2][localCol+2] == 'M' && split[row+3][localCol+3] == 'X' {
						out += 1
					}
				}
			}
			row++
			localCol++
		}
	}

	return out
}

func downLeft(s string) int {
	var out int

	split := strings.Split(strings.TrimSpace(s), "\n")
	width := len(split[0])
	height := len(split)

	for row := 0; row < height-1; row++ { // Left edge
		localRow := row
		col := 0

		for localRow > 0 {
			switch split[localRow][col] {
			case 'X':
				if localRow >= 3 && col <= width-4 {
					if split[localRow-1][col+1] == 'M' && split[localRow-2][col+2] == 'A' && split[localRow-3][col+3] == 'S' {
						out += 1
					}
				}
			case 'S':
				if localRow >= 3 && col <= width-4 {
					if split[localRow-1][col+1] == 'A' && split[localRow-2][col+2] == 'M' && split[localRow-3][col+3] == 'X' {
						out += 1
					}
				}
			}
			localRow--
			col++
		}
	}

	for col := 0; col < width; col++ { // Bottom edge TO DO
		row := width - 1
		localCol := col

		for localCol < width-1 {
			switch split[row][localCol] {
			case 'X':
				if row >= 3 && localCol <= width-4 {
					if split[row-1][localCol+1] == 'M' && split[row-2][localCol+2] == 'A' && split[row-3][localCol+3] == 'S' {
						out += 1
					}
				}
			case 'S':
				if row >= 3 && localCol <= width-4 {
					if split[row-1][localCol+1] == 'A' && split[row-2][localCol+2] == 'M' && split[row-3][localCol+3] == 'X' {
						out += 1
					}
				}
			}
			row--
			localCol++
		}
	}

	return out
}

func scanForAs(s string) int {
	var out int
	split := strings.Split(strings.TrimSpace(s), "\n")

	for row := 1; row < len(split)-1; row++ {
		for col := 1; col < len(split[0])-1; col++ {
			if split[row][col] == 'A' {
				if checkA(split, row, col) {
					out += 1
				}
			}
		}
	}
	return out
}

func checkA(s []string, row, col int) bool {
	m := make([][]int, 0)
	x := make([][]int, 0)

	switch s[row-1][col-1] {
	case 'M':
		m = append(m, []int{row - 1, col - 1})
	case 'S':
		x = append(x, []int{row - 1, col - 1})
	}
	switch s[row-1][col+1] {
	case 'M':
		m = append(m, []int{row - 1, col + 1})
	case 'S':
		x = append(x, []int{row - 1, col + 1})
	}
	switch s[row+1][col-1] {
	case 'M':
		m = append(m, []int{row + 1, col - 1})
	case 'S':
		x = append(x, []int{row + 1, col - 1})
	}
	switch s[row+1][col+1] {
	case 'M':
		m = append(m, []int{row + 1, col + 1})
	case 'S':
		x = append(x, []int{row + 1, col + 1})
	}

	if len(m) == len(x) && len(m) == 2 {
		if m[0][0] == m[1][0] || m[0][1] == m[1][1] { // For it to be an X, the M's have to be in the same column or the same row
			return true
		}
	}

	return false
}
