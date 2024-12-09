package day_09

import (
	"strconv"
)

// The dumb way, build the entire thing
func parse(s string) []int {
	var fileID int
	out := make([]int, 0)

	for i, r := range s {
		value, _ := strconv.Atoi(string(r))
		if i%2 == 0 { // File
			for j := 0; j < value; j++ {
				out = append(out, fileID)
			}
			fileID++
		} else { // Free space
			for j := 0; j < value; j++ {
				out = append(out, -1)
			}
		}
	}
	return out
}

func defragment(s []int) []int {
	out := make([]int, 0)

	frontIndex := 0
	backIndex := len(s) - 1

	for frontIndex <= backIndex {
		if s[frontIndex] != -1 {
			out = append(out, s[frontIndex])
			frontIndex++
		} else {
			if s[backIndex] != -1 {
				out = append(out, s[backIndex])
				backIndex--
				frontIndex++
			} else {
				backIndex--
			}
		}
	}

	return out
}

func checkSum(s []int) int {
	var out int

	for i, val := range s {
		if val >= 0 {
			out += i * val
		}
	}
	return out
}

func moveFiles(s []int) []int {
	localSlice := s
	lastIndex := len(s) - 1

	for fileID := s[len(s)-1]; fileID > 0; fileID-- {
		// Is there open space in front of that last file?
		for localSlice[lastIndex] == -1 || localSlice[lastIndex] > fileID {
			lastIndex--
		}

		var count int
		for index := lastIndex; index >= 0; index-- {
			if localSlice[index] == fileID {
				count++ // We know how long the file is
			}
		}

		lastIndex = lastIndex - count

		// Find the first empty space that can hold this file
		found, startIndex := findSpace(localSlice, lastIndex, count)
		if found {
			for i := 0; i < count; i++ {
				localSlice[startIndex+i] = fileID
				localSlice[lastIndex+i+1] = -1
			}
		} else {
			continue
		}
	}

	return localSlice
}

func findSpace(s []int, maxIndex int, len int) (bool, int) {
	var emptyCount int
	startIndex := -1

	for i := 0; i <= maxIndex; i++ {
		if s[i] == -1 {
			if startIndex == -1 {
				startIndex = i
			}
			emptyCount++
			if emptyCount == len {
				return true, startIndex
			}
		} else {
			startIndex = -1
			emptyCount = 0
		}
	}

	return false, 0
}

// Better to calculate the checksum on the fly

/*
func checkSum(s string) int {
	// Even indexes are files, odd indexes are free space
	frontIndex := 0
	frontID := 0
	blockPosition := 0

	backIndex := len(s) - 1
	backID := len(s) / 2
	remaining := 0

	var out int

	for frontIndex < backIndex {
		// Step 1: is the front index odd or even?
		if frontIndex%2 == 0 { // Even, this is a file and we need to add to the checksum
			value, _ := strconv.Atoi(string(s[frontIndex]))
			for i := 0; i < value; i++ {
				out += (blockPosition + i) * frontID
			}
			blockPosition += value
			frontID++
		} else { // Odd - this is free space
			// Do we have any remaining of the last file we found?
			if remaining != 0 {
				//
			} else { // We don't, get the last file
				//
			}
			// Get the last value
		}
	}
}
*/
