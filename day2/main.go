package day1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Day2 struct {
}

func (d Day2) readIntervals() []string {
	content, err := os.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return strings.Split(string(content), ",")
}

func (d Day2) SolveFirstHalf() int {
	intervals := d.readIntervals()

	ans := 0
	for _, interval := range intervals {
		interval := strings.Split(interval, "-")
		a, err := strconv.Atoi(strings.Split(interval[0], "\n")[0])
		if err != nil {
			log.Fatalf("Error to convert string to int: %v", err)
		}

		b, err := strconv.Atoi(strings.Split(interval[1], "\n")[0])
		if err != nil {
			log.Fatalf("Error to convert string to int: %v", err)
		}

		for i := a; i <= b; i++ {
			strNumber := strconv.Itoa(i)
			if len(strNumber)%2 == 0 && strNumber[:len(strNumber)/2] == strNumber[len(strNumber)/2:] {
				ans += i
			}
		}
	}

	return ans
}

func (d Day2) SolveSecondHalf() int {
	intervals := d.readIntervals()

	ans := 0
	for _, interval := range intervals {
		interval := strings.Split(interval, "-")
		a, err := strconv.Atoi(strings.Split(interval[0], "\n")[0])
		if err != nil {
			log.Fatalf("Error to convert string to int: %v", err)
		}

		b, err := strconv.Atoi(strings.Split(interval[1], "\n")[0])
		if err != nil {
			log.Fatalf("Error to convert string to int: %v", err)
		}

		for i := a; i <= b; i++ {
			strNumber := strconv.Itoa(i)

			for j := range len(strNumber) - 1 {
				if len(strNumber)%(j+1) == 0 && strings.Repeat(strNumber[:j+1], len(strNumber)/(j+1)) == strNumber {
					ans += i
					break
				}
			}
		}
	}

	return ans
}
