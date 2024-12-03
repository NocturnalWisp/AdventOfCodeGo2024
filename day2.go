package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getListFromInput(lines []string, list *[][]int) {
	*list = make([][]int, len(lines))

	for i, line := range lines {
		numsSplit := strings.Split(line, " ")

		levelList := make([]int, len(numsSplit))

		for j, numStr := range numsSplit {
			num, err := strconv.Atoi(numStr)

			if err != nil {
				fmt.Println("Invalid num parsed.")
			}
			levelList[j] = num
		}

		(*list)[i] = levelList
	}
}

func day2(input string) {
	lines := strings.Split(input, "\n")

	var levelList [][]int

	getListFromInput(lines, &levelList)

	succeeded := 0

	for _, level := range levelList {
		prevValue := level[0]
		prevDifference := 0

		found := checkLevel(level, &prevValue, &prevDifference)

		if found {
			succeeded++
		}
	}

	fmt.Println(succeeded)
}

func day2p2(input string) {
	lines := strings.Split(input, "\n")

	var levelList [][]int

	getListFromInput(lines, &levelList)

	succeeded := 0

	for _, level := range levelList {
		prevValue := level[0]
		prevDifference := 0

		found := checkLevel(level, &prevValue, &prevDifference)

		if found {
			succeeded++
		} else {
			for i := 0; i < len(level); i++ {
				if i == 0 {
					prevValue = level[1]
					prevDifference = 1
				} else {
					prevValue = level[0]
					prevDifference = 0
				}

				firstHalf := append([]int{}, level[:i]...)
				secondHalf := append([]int{}, level[i+1:]...)
				combine := append(firstHalf, secondHalf...)
				found = checkLevel(combine, &prevValue, &prevDifference)

				if found {
					succeeded++
					break
				} else {
				}
			}
		}
	}

	fmt.Println(succeeded)
}

func checkLevel(level []int, prevValue *int, prevDifference *int) bool {
	found := false

	for i := 1; i < len(level); i++ {
		found = false

		difference := level[i] - *prevValue

		if difference == 0 || difference > 3 || difference < -3 {
			break
		}

		if i != 1 {
			if sign(*prevDifference) != sign(difference) {
				break
			}
		}

		*prevValue = level[i]
		*prevDifference = difference

		found = true
	}
	return found
}

func sign(value int) bool {
	if value < 0 {
		return false
	} else {
		return true
	}
}
