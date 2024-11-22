package main

import (
	"advent/day_01"
	"advent/day_02"
	"advent/day_03"
	"advent/day_04"
	"advent/day_05"
	"advent/day_06"
	"advent/day_07"
	"advent/day_08"
	"advent/day_09"
	"advent/day_10"
	"advent/day_11"
	"advent/day_12"
	"advent/day_13"
	"advent/day_14"
	"advent/day_15"
	"advent/day_16"
	"advent/day_17"
	"advent/day_18"
	"advent/day_19"
	"advent/day_20"
	"advent/day_21"
	"advent/day_22"
	"advent/day_23"
	"advent/day_24"
	"advent/day_25"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No argument was passed - exiting.")
		return
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The argument is not an integer - exiting.")
	}

	solved := []func() (interface{}, interface{}){
		day_01.Solve,
		day_02.Solve,
		day_03.Solve,
		day_04.Solve,
		day_05.Solve,
		day_06.Solve,
		day_07.Solve,
		day_08.Solve,
		day_09.Solve,
		day_10.Solve,
		day_11.Solve,
		day_12.Solve,
		day_13.Solve,
		day_14.Solve,
		day_15.Solve,
		day_16.Solve,
		day_17.Solve,
		day_18.Solve,
		day_19.Solve,
		day_20.Solve,
		day_21.Solve,
		day_22.Solve,
		day_23.Solve,
		day_24.Solve,
		day_25.Solve,
	}

	if day > 0 && day <= 25 {
		fmt.Println("Solutions for day", day)
		first, second := solved[day-1]()
		fmt.Println(first, second)
	} else {
		fmt.Println("That's not a valid day.")
	}
}
