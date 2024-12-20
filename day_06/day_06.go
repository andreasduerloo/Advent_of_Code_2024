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

	// First star
	for g.inBounds {
		g.step(b)
	}

	first := g.unique

	// Second star
	toTry := helpers.Uniq(g.path) // We only try points already on the path, and each point only once (even if it is on the path multiple times)
	var second int

	/*
		b, g = parseMap(inStr) // Load evetything back up
		startingPosition := g.location

		start := time.Now()


			for _, position := range toTry {
				if position == startingPosition {
					continue
				}

				b.obstacles[position] = true

				for g.inBounds && !g.cycle {
					g.step(b)
				}

				if g.cycle {
					second++
				}

				delete(b.obstacles, position)
				g.reset(startingPosition)
			}

			fmt.Println(time.Since(start))
	*/

	// Second try to multithread
	found := make(chan struct{})

	go func() {
		for range found {
			second++
		}
	}()

	parts := [][]point{
		toTry[0 : len(toTry)/4],
		toTry[(len(toTry)/4)+1 : len(toTry)/2],
		toTry[(len(toTry)/2)+1 : 3*(len(toTry)/4)],
		toTry[3*(len(toTry)/4)+1:],
	}

	b, g = parseMap(inStr)
	startingPosition := g.location

	var wg sync.WaitGroup

	start := time.Now()
	for _, part := range parts {
		wg.Add(1)

		lb := b.copy()
		lg := g.copy()

		go func() {
			defer wg.Done()

			for _, position := range part {
				if position == startingPosition {
					continue
				}

				lg.reset(startingPosition)
				lb.obstacles[position] = true

				for lg.inBounds && !lg.cycle {
					lg.step(lb)
				}

				if lg.cycle {
					found <- struct{}{}
				}

				delete(lb.obstacles, position)
			}
		}()
	}

	wg.Wait()
	close(found)

	fmt.Println(time.Since(start))

	return first, second
}
