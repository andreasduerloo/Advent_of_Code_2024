package day_12

import (
	"advent/helpers"
	"fmt"
)

func Solve() (interface{}, interface{}) {
	inStr, err := helpers.GetInput(12)
	if err != nil {
		fmt.Println("There was an issue getting the input")
	}

	currentID := 1

	layout := parse(inStr)
	fields := make(map[int]field)

	for y := 0; y < layout.height; y++ {
		for x := 0; x < layout.width; x++ {
			if layout.squares[point{x, y}].field == 0 { // This square is not yet in a field - build the field
				buildField(point{x, y}, currentID, layout, fields)
				currentID++
			}
		}
	}

	var first int
	for _, v := range fields {
		first += (v.size * v.borders)
	}

	var second int
	for _, v := range fields {
		// fmt.Println(v.id, v.squares, calcSides(v.squares))
		second += (v.size * calcSides(v.squares))
	}

	return first, second
}
