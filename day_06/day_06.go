package day_06

import (
	"advent/helpers"
	"fmt"
	"sync"
	"time"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(6)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	b, g := parseMap(inStr)
	startingPosition := g.location

	// First star
	for g.inBounds {
		g.step(b)
	}

	first := g.unique

	// Second star
	toTry := helpers.Uniq(g.path) // We only try points already on the path, and each point only once (even if it is on the path multiple times)

	var validPositions int
	found := make(chan struct{}, 5)

	go func() {
		for range found {
			validPositions++
		}
	}()

	// Split the input into four parts
	parts := [][]point{toTry[0 : len(toTry)/4], toTry[(len(toTry)/4)+1 : len(toTry)/2], toTry[(len(toTry)/2)+1 : 3*len(toTry)/4], toTry[3*(len(toTry)/4)+1:]}

	b, g = parseMap(inStr) // Load everything up from scratch

	// Make a WaitGroup
	var wg sync.WaitGroup

	start := time.Now()
	for _, part := range parts {
		wg.Add(1)

		go func() {
			defer wg.Done()

			lg := g.copy()
			lb := b.copy()

			for _, position := range part {
				if position == startingPosition {
					continue
				}

				lb.obstacles[position] = true

				for lg.inBounds && !lg.cycle {
					lg.step(lb)
				}

				if lg.cycle {
					found <- struct{}{}
				}

				lb.obstacles[position] = false
			}
		}()
	}

	wg.Wait()
	close(found)

	end := time.Now()
	fmt.Println(end.Sub(start))

	return first, validPositions
}
