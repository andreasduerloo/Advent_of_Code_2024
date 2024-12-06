package day_06

import (
	"advent/helpers"
	"fmt"
	"time"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(6)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	b, g := parseMap(inStr)

	// First star
	for g.inBounds {
		g.step(b)
	}

	first := g.unique

	// Second star
	toTry := helpers.Uniq(g.path) // We only try points already on the path, and each point only once (even if it is on the path multiple times)
	var second int

	b, g = parseMap(inStr) // Load evetything back up
	startingPosition := g.location

	start := time.Now()

	for _, position := range toTry {
		if position == startingPosition {
			continue
		}

		// Reset the guard
		// lg.reset(startingPosition)
		g.reset(startingPosition)

		b.obstacles[position] = true

		for g.inBounds && !g.cycle {
			g.step(b)
		}

		if g.cycle {
			second++
		}

		delete(b.obstacles, position)
	}

	fmt.Println(time.Since(start))
	return first, second
}

/*
var validPositions int
	found := make(chan int, 5)

	go func() {
		for result := range found {
			validPositions += result
		}
	}()

	// Split the input into four parts
	parts := [][]point{toTry[0 : len(toTry)/4], toTry[(len(toTry)/4)+1 : len(toTry)/2], toTry[(len(toTry)/2)+1 : 3*len(toTry)/4], toTry[3*(len(toTry)/4)+1:]}

	fmt.Println(len(toTry), len(parts[0])+len(parts[1])+len(parts[2])+len(parts[3]))

	b, g = parseMap(inStr) // Load everything up from scratch

	// Make a WaitGroup
	var wg sync.WaitGroup

	start := time.Now()
	for _, part := range parts {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// lg := g.copy()
			lb := b.copy()

			var localResult int

			for _, position := range part {
				if position == startingPosition {
					continue
				}

				// Reset the guard
				// lg.reset(startingPosition)
				lg := g.copy()

				lb.obstacles[position] = true

				for lg.inBounds && !lg.cycle {
					lg.step(lb)
				}

				if lg.cycle {
					localResult++
				}

				delete(lb.obstacles, position)
			}
			fmt.Println(localResult)
			found <- localResult
		}()
	}

	wg.Wait()
	close(found)

	end := time.Now()
	fmt.Println(end.Sub(start))
*/
