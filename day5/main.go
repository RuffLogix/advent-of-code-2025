package day5

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day5 struct {
}

func (d Day5) readIngredients() ([][]int, []int) {
	content, err := os.ReadFile("day5/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	intervals := [][]int{}
	ingredientIDs := []int{}
	isIntervals := true
	for _, text := range strings.Split(string(content), "\n") {
		if text == "" {
			isIntervals = false
			continue
		}

		if isIntervals {
			text := strings.Split(text, "-")
			beginNumber, err := strconv.Atoi(text[0])
			if err != nil {
				log.Fatalf("Error to convert string to int: %v", err)
			}

			endNumber, err := strconv.Atoi(text[1])
			if err != nil {
				log.Fatalf("Error to convert string to int: %v", err)
			}

			intervals = append(intervals, []int{beginNumber, endNumber})
		} else {
			number, err := strconv.Atoi(text)
			if err != nil {
				log.Fatalf("Error to convert string to int: %v", err)
			}

			ingredientIDs = append(ingredientIDs, number)
		}
	}

	return intervals, ingredientIDs
}

func (d Day5) SolveFirstHalf() int {
	intervals, ingredientIDs := d.readIngredients()

	ans := 0
	for i := range ingredientIDs {
		for j := range intervals {
			if intervals[j][0] <= ingredientIDs[i] && ingredientIDs[i] <= intervals[j][1] {
				ans += 1
				break
			}
		}
	}

	return ans
}

func (d Day5) removeDuplicatedIntervals(intervals [][]int) [][]int {
	ans := [][]int{}
	for _, interval := range intervals {
		if len(ans) == 0 {
			ans = append(ans, interval)
		} else {
			if ans[len(ans)-1][1] < interval[0] {
				ans = append(ans, interval)
			} else if ans[len(ans)-1][1] >= interval[0] && ans[len(ans)-1][1] < interval[1] {
				ans[len(ans)-1][1] = interval[1]
			}
		}
	}
	return ans
}

func (d Day5) SolveSecondHalf() int {
	intervals, _ := d.readIngredients()

	sort.Slice(intervals, func(i int, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	intervals = d.removeDuplicatedIntervals(intervals)

	ans := 0
	for _, interval := range intervals {
		ans += interval[1] - interval[0] + 1
	}

	return ans
}
