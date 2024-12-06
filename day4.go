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
		if recursiveCheckLetter(info, 'X', i, 0, getNextLetterP1) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 1, getNextLetterP1) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 2, getNextLetterP1) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 3, getNextLetterP1) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 4, getNextLetterP1) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 5, getNextLetterP1) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 6, getNextLetterP1) {
			total++
		}
		if recursiveCheckLetter(info, 'X', i, 7, getNextLetterP1) {
			total++
		}
	}

	fmt.Println(total)
}

func getNextLetterP1(currentLetter byte) byte {
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

func day4p2(input string) {
	firstLineLength := strings.Index(input, "\n")

	newInput := strings.ReplaceAll(input, "\n", "")

	info := puzzleInfo{
		input:      newInput,
		lineLength: firstLineLength,
		length:     len(newInput),
	}

	total := 0

	for i, char := range newInput {
		if char != 'A' {
			continue
		}

		amountCorrect := 0

		topLeftIndex := getNextTargetPosition(info, i, 0)
		if topLeftIndex == -1 {
			continue
		}
		bottomRightIndex := getNextTargetPosition(info, i, 7)
		if bottomRightIndex == -1 {
			continue
		}

		topRightIndex := getNextTargetPosition(info, i, 2)
		if topRightIndex == -1 {
			continue
		}
		bottomLeftIndex := getNextTargetPosition(info, i, 5)
		if bottomLeftIndex == -1 {
			continue
		}

		if recursiveCheckLetter(info, 'M', topLeftIndex, 7, getNextLetterP2) {
			amountCorrect++
		} else if recursiveCheckLetter(info, 'M', bottomRightIndex, 0, getNextLetterP2) {
			amountCorrect++
		}

		if recursiveCheckLetter(info, 'M', topRightIndex, 5, getNextLetterP2) {
			amountCorrect++
		} else if recursiveCheckLetter(info, 'M', bottomLeftIndex, 2, getNextLetterP2) {
			amountCorrect++
		}

		if amountCorrect == 2 {
			total++
		}
	}

	fmt.Println(total)
}

func getNextLetterP2(currentLetter byte) byte {
	switch currentLetter {
	case 'M':
		return 'A'
	case 'A':
		return 'S'
	case 'S':
		return '.'
	}

	return ' '
}

// Direction:
// o - 1 - 2
// 3 - X - 4
// 5 - 6 - 7
func recursiveCheckLetter(info puzzleInfo, targetLetter byte, currentIndex int, direction int, getLetterFunc func(byte) byte) bool {
	if info.input[currentIndex] != targetLetter {
		return false
	}

	nextLetter := getLetterFunc(targetLetter)

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

	return recursiveCheckLetter(info, nextLetter, nextIndex, direction, getLetterFunc)
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
