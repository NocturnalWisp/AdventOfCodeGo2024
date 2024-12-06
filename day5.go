package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	before int
	after  int
}

func parseInput(input string) (pairs []pair, pages [][]int) {
	pairsStr, pagesStr, _ := strings.Cut(input, "\n\n")

	pairsListStr := strings.Split(pairsStr, "\n")

	pairs = make([]pair, len(pairsListStr))

	for i, pairStr := range pairsListStr {
		firstOfPairStr, secondOfPairStr, _ := strings.Cut(pairStr, "|")

		firstOfPair, _ := strconv.Atoi(firstOfPairStr)
		secondOfPair, _ := strconv.Atoi(secondOfPairStr)

		pairs[i] = pair{
			before: firstOfPair,
			after:  secondOfPair,
		}
	}

	pagesListStr := strings.Split(pagesStr, "\n")

	pages = make([][]int, len(pagesListStr))

	for i, pageStr := range pagesListStr {
		pageNumberListStr := strings.Split(pageStr, ",")

		pages[i] = make([]int, len(pageNumberListStr))

		for j, numberStr := range pageNumberListStr {
			number, _ := strconv.Atoi(numberStr)

			pages[i][j] = number
		}
	}

	return
}

type byRules struct {
	page  []int
	rules []pair
}

func (b byRules) Len() int {
	return len(b.page)
}

func (b byRules) Swap(i, j int) {
	b.page[i], b.page[j] = b.page[j], b.page[i]
}

func (b byRules) Less(i, j int) bool {
	first, second := b.page[i], b.page[j]

	for _, rule := range b.rules {
		if rule.before == first && rule.after == second {
			return true
		} else if rule.before == second && rule.after == first {
			return false
		}
	}

	return false
}

func day5(input string) {
	pairs, pages := parseInput(input)

	total := 0

	for _, page := range pages {
		sort.Sort(byRules{page, pairs})

		length := float64(len(page))
		halfLength := length / 2
		total += page[int(math.Floor(halfLength))]

		fmt.Println(page)
	}

	fmt.Println(total)
}
