package main

import (
	"fmt"
	"strings"
)

type puzzleInfo struct {
	input      string
	lineLength int
	length     int
}

func day4(input string) {
	firstLineLength := strings.Index(input, "\n")

	newInput := strings.ReplaceAll(input, "\n", "")

	info := puzzleInfo{
		input:      newInput,
		lineLength: firstLineLength,
		length:     len(newInput),
	}

	total := 0

	for i := range info.input {
		if recursiveCheckLetter(info, 'X', i, 0) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 1) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 2) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 3) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 4) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 5) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 6) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 7) {
			total++
		}
	}

	fmt.Println(total)
}

// Direction:
// o - 1 - 2
// 3 - X - 4
// 5 - 6 - 7
func recursiveCheckLetter(info puzzleInfo, targetLetter byte, currentIndex int, direction int) bool {
	if info.input[currentIndex] != targetLetter {
		return false
	}

	nextLetter := getNextLetter(targetLetter)

	if nextLetter == ' ' {
		return false
	}

	if nextLetter == '.' {
		return true
	}

	nextIndex := getNextTargetPosition(info, currentIndex, direction)

	if nextIndex == -1 {
		return false
	}

	return recursiveCheckLetter(info, nextLetter, nextIndex, direction)
}

func getNextLetter(currentLetter byte) byte {
	switch currentLetter {
	case 'X':
		return 'M'
	case 'M':
		return 'A'
	case 'A':
		return 'S'
	case 'S':
		return '.'
	}

	return ' '
}

func getNextTargetPosition(info puzzleInfo, currentIndex int, direction int) int {
	newIndex := -1

	switch direction {
	case 0:
		newIndex = get0Dir(info, currentIndex)
	case 1:
		newIndex = get1Dir(info, currentIndex)
	case 2:
		newIndex = get2Dir(info, currentIndex)
	case 3:
		newIndex = get3Dir(info, currentIndex)
	case 4:
		newIndex = get4Dir(info, currentIndex)
	case 5:
		newIndex = get5Dir(info, currentIndex)
	case 6:
		newIndex = get6Dir(info, currentIndex)
	case 7:
		newIndex = get7Dir(info, currentIndex)
	}

	return newIndex
}

func InTopEdge(currentIndex int, lineLength int) bool {
	if currentIndex < lineLength {
		return true
	}

	return false
}

func InLeftEdge(currentIndex int, lineLength int) bool {
	if currentIndex%lineLength == 0 {
		return true
	}

	return false
}

func InRightEdge(currentIndex int, lineLength int) bool {
	if currentIndex%lineLength == lineLength-1 {
		return true
	}

	return false
}

func InBottomEdge(currentIndex int, lineLength int, totalLength int) bool {
	if currentIndex > totalLength-lineLength-1 {
		return true
	}

	return false
}

func get0Dir(info puzzleInfo, currentIndex int) int {
	if InLeftEdge(currentIndex, info.lineLength) {
		return -1
	}

	if InTopEdge(currentIndex, info.lineLength) {
		return -1
	}

	return currentIndex - info.lineLength - 1
}

func get1Dir(info puzzleInfo, currentIndex int) int {
	if InTopEdge(currentIndex, info.lineLength) {
		return -1
	}

	return currentIndex - info.lineLength
}

func get2Dir(info puzzleInfo, currentIndex int) int {
	if InRightEdge(currentIndex, info.lineLength) {
		return -1
	}

	if InTopEdge(currentIndex, info.lineLength) {
		return -1
	}

	return currentIndex - info.lineLength + 1
}

func get3Dir(info puzzleInfo, currentIndex int) int {
	if InLeftEdge(currentIndex, info.lineLength) {
		return -1
	}

	return currentIndex - 1
}

func get4Dir(info puzzleInfo, currentIndex int) int {
	if InRightEdge(currentIndex, info.lineLength) {
		return -1
	}

	return currentIndex + 1
}

func get5Dir(info puzzleInfo, currentIndex int) int {
	if InBottomEdge(currentIndex, info.lineLength, info.length) {
		return -1
	}

	if InLeftEdge(currentIndex, info.lineLength) {
		return -1
	}

	return currentIndex + info.lineLength - 1
}

func get6Dir(info puzzleInfo, currentIndex int) int {
	if InBottomEdge(currentIndex, info.lineLength, info.length) {
		return -1
	}

	return currentIndex + info.lineLength
}

func get7Dir(info puzzleInfo, currentIndex int) int {
	if InBottomEdge(currentIndex, info.lineLength, info.length) {
		return -1
	}

	if InRightEdge(currentIndex, info.lineLength) {
		return -1
	}

	return currentIndex + info.lineLength + 1
}
