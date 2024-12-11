package day_11

import "strconv"

func blink(rocks []int) []int {
	out := make([]int, 0)

	for _, rock := range rocks {
		if rock == 0 {
			out = append(out, 1)
			continue
		}

		numStr := strconv.Itoa(rock)
		if len(numStr)%2 == 0 {
			leftStr := numStr[0 : len(numStr)/2]
			rightStr := numStr[(len(numStr) / 2):]

			left, _ := strconv.Atoi(leftStr)
			right, _ := strconv.Atoi(rightStr)

			out = append(out, left)
			out = append(out, right)
			continue
		}

		out = append(out, rock*2024)
	}
	return out
}

type state struct {
	rock  int
	depth int
}

func blinkAndRemember(rock int, mem map[state]int, depth, maxDepth int) int {
	if depth < maxDepth {
		if val, present := mem[state{rock, depth}]; present {
			return val
		}

		if rock == 0 {
			length := blinkAndRemember(1, mem, depth+1, maxDepth)
			mem[state{rock, depth}] = length
			return length
		}

		numStr := strconv.Itoa(rock)
		if len(numStr)%2 == 0 {
			leftStr := numStr[0 : len(numStr)/2]
			rightStr := numStr[(len(numStr) / 2):]

			left, _ := strconv.Atoi(leftStr)
			right, _ := strconv.Atoi(rightStr)

			lLength := blinkAndRemember(left, mem, depth+1, maxDepth)
			rLength := blinkAndRemember(right, mem, depth+1, maxDepth)

			mem[state{rock, depth}] = lLength + rLength
			return lLength + rLength
		}

		length := blinkAndRemember(rock*2024, mem, depth+1, maxDepth)
		mem[state{rock, depth}] = length
		return length

	} else {
		return 1
	}
}
