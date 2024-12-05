package helpers

import (
	"errors"
	"os"
	"regexp"
	"strconv"
)

func GetInput(day int) (string, error) {
	if day > 0 && day <= 25 {
		dayString := strconv.Itoa(day)

		if len(dayString) == 1 {
			dayString = "0" + dayString
		}

		path := "./inputs/" + dayString + ".txt"

		input, err := os.ReadFile(path)
		if err != nil {
			return "", err
		} else {
			return string(input), nil
		}
	} else {
		return "", errors.New("not a valid day")
	}
}

func ReGetInts(s string) []int {
	re := regexp.MustCompile(`-?[0-9]+`)
	matches := re.FindAllString(s, -1)

	ints := make([]int, 0)

	for _, match := range matches {
		val, err := strconv.Atoi(match)
		if err != nil {
			continue
		}
		ints = append(ints, val)
	}

	return ints
}

func Filter[T any](s []T, f func(T) bool) []T {
	out := make([]T, 0)

	for _, elem := range s {
		if f(elem) {
			out = append(out, elem)
		}
	}

	return out
}

func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0)

	for _, elem := range s {
		out = append(out, f(elem))
	}

	return out
}

func MapReduce[T, U any](s []T, acc U, f func(T, U) U) U {
	out := acc

	for _, elem := range s {
		out = f(elem, out)
	}

	return out
}
