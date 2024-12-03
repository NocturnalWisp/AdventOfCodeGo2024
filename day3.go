package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func checkValidInnerNumber(target string) bool {
	for _, char := range target {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

// Result: 164730528
func day3(input string) {
	total := 0

	betweenMuls := strings.Split(input, "mul(")

	if !strings.HasPrefix(input, "mul(") {
		betweenMuls = betweenMuls[1:]
	}

	for _, nextMul := range betweenMuls {
		before, _, found := strings.Cut(nextMul, ")")

		if !found {
			continue
		}

		innerMulSplit := strings.Split(before, ",")
		innerMulCount := len(innerMulSplit)

		if innerMulCount != 2 {
			continue
		}

		num1Str := innerMulSplit[0]
		num2Str := innerMulSplit[1]

		if !checkValidInnerNumber(num1Str) {
			continue
		}
		if !checkValidInnerNumber(num2Str) {
			continue
		}

		num1, _ := strconv.Atoi(num1Str)
		num2, _ := strconv.Atoi(num2Str)

		total += num1 * num2
	}

	fmt.Println(total)
}

// Output: 70478672
func day3p2(input string) {
	total := 0

	betweenMuls := strings.Split(input, "mul(")

	if !strings.HasPrefix(input, "mul(") {
		betweenMuls = betweenMuls[1:]
	}

	enabled := true

	for _, nextMul := range betweenMuls {
		wasEnabled := enabled

		before, after, found := strings.Cut(nextMul, ")")

		lastDo := strings.LastIndex(after, "do()")
		lastDont := strings.LastIndex(after, "don't()")

		if lastDo != -1 || lastDont != -1 {
			if lastDo > lastDont {
				enabled = true
			} else {
				enabled = false
			}
		}

		if !found {
			continue
		}

		if !wasEnabled {
			continue
		}

		innerMulSplit := strings.Split(before, ",")
		innerMulCount := len(innerMulSplit)

		if innerMulCount != 2 {
			continue
		}

		num1Str := innerMulSplit[0]
		num2Str := innerMulSplit[1]

		if !checkValidInnerNumber(num1Str) {
			continue
		}
		if !checkValidInnerNumber(num2Str) {
			continue
		}

		num1, _ := strconv.Atoi(num1Str)
		num2, _ := strconv.Atoi(num2Str)

		total += num1 * num2
	}

	fmt.Println(total)
}
