package day3

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day3 struct {
}

func (d Day3) readJoltages() []string {
	content, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return strings.Split(string(content), "\n")
}

func (d Day3) SolveFirstHalf() int {
	joltages := d.readJoltages()

	ans := 0
	for _, joltage := range joltages {
		mxVal := 0
		for i := range len(joltage) {
			for j := i + 1; j < len(joltage); j++ {
				joltageVal, err := strconv.Atoi(joltage[i:i+1] + joltage[j:j+1])
				if err != nil {
					log.Fatalf("Error to convert string to int: %v", err)
				}

				mxVal = int(math.Max(float64(mxVal), float64(joltageVal)))
			}
		}
		ans += mxVal
	}

	return ans
}

var maxValSecondSolution = 0

func (d Day3) solveSecondHalf(index int, count int, joltageCollect int, joltageOrig string) int {
	if count == 12 {
		maxValSecondSolution = int(math.Max(float64(joltageCollect), float64(maxValSecondSolution)))
		return maxValSecondSolution
	}
	if index == len(joltageOrig) {
		return 0
	}
	if maxValSecondSolution > (joltageCollect+1)*int(math.Pow10(12-count)) {
		return maxValSecondSolution
	}

	returnValue := int(math.Max(float64(d.solveSecondHalf(index+1, count+1, joltageCollect*10+int(joltageOrig[index]-'0'), joltageOrig)), float64(d.solveSecondHalf(index+1, count, joltageCollect, joltageOrig))))
	maxValSecondSolution := int(math.Max(float64(returnValue), float64(maxValSecondSolution)))

	return maxValSecondSolution
}

func (d Day3) SolveSecondHalf() int {
	joltages := d.readJoltages()

	ans := 0
	for _, joltage := range joltages {
		maxValSecondSolution = 0
		mxVal := d.solveSecondHalf(0, 0, 0, joltage)
		ans += mxVal
	}

	return ans
}
