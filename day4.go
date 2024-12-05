package main

import "strings"

type puzzleInfo struct {
	input      string
	lineLength int
	length     int
}

func day4(input string) {
	firstLineLength := strings.Index(input, "\n") + 1
	count := strings.Count(input, "\n")

	info := puzzleInfo{
		input:      input,
		lineLength: firstLineLength,
		length:     len(input),
	}

	for i := 0; i < count; i++ {

	}
}

// Direction:
// o - 1 - 2
// 3 - X - 4
// 5 - 6 - 7
func recursiveCheckLetter(info puzzleInfo, targetLetter rune, currentIndex int, direction int) {
}

func getNextTargetLetter(info puzzleInfo, currentIndex int, direction int) (int, bool) {
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

	isValid := newIndex != -1

	return newIndex, isValid
}

func InTopEdge(currentIndex int, lineLength int) bool {
	if currentIndex < lineLength {
		return true
	}

	return false
}

func InLeftEdge(currentIndex int, lineLength int) bool {
	if currentIndex%lineLength != 0 {
		return true
	}

	return false
}

func InRightEdge(currentIndex int, lineLength int) bool {
	if currentIndex%lineLength != lineLength-1 {
		return true
	}

	return false
}

func InBottomEdge(currentIndex int, lineLength int, totalLength int) bool {
	if currentIndex > totalLength-lineLength {
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
