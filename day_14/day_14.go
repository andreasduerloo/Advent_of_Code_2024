package day_14

import (
	"advent/helpers"
	"fmt"
	"time"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(14)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	robots := parse(inStr)

	quad := quadrants{
		values: [4]int{0, 0, 0, 0},
	}

	for _, r := range robots {
		for i := 0; i < 100; i++ {
			r.move()
		}

		quad.add(r.quadrant())
	}

	first := quad.calculate()

	defer fmt.Println(first)

	// Second star

	robots = parse(inStr)

	var i int
	for {
		for _, r := range robots {
			r.move()
		}
		i++

		if i == 52 || (i-52)%101 == 0 {
			drawFrame(robots, i)
		}

		time.Sleep(50 * time.Millisecond)
	}

	return first, 0
}
