package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day1(input string) {
	lines := strings.Split(input, "\n")

	var list1 []int
	var list2 []int

	getListsFromInput(lines, &list1, &list2)

	sort.Ints(list1)
	sort.Ints(list2)

	total := getTotalListDifference(lines, list2, list1)

	fmt.Println(total)
}

func getTotalListDifference(lines []string, list2 []int, list1 []int) int {
	total := 0

	for i := 0; i < len(lines); i++ {
		total += abs(list2[i] - list1[i])
	}
	return total
}

func getListsFromInput(lines []string, list1 *[]int, list2 *[]int) {
	*list1 = make([]int, len(lines))
	*list2 = make([]int, len(lines))

	for i, line := range lines {
		numsSplit := strings.Split(line, "   ")

		num1, err1 := strconv.Atoi(numsSplit[0])
		num2, err2 := strconv.Atoi(numsSplit[1])

		if err1 != nil {
			fmt.Println("Error in num1!")
		}
		if err2 != nil {
			fmt.Println("Error in num2!")
		}

		(*list1)[i] = num1
		(*list2)[i] = num2
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
