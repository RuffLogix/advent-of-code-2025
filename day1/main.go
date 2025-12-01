package day1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Day1 struct {
}

func (d Day1) readActions() []string {
	content, err := os.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return strings.Split(string(content), "\n")
}

func (d Day1) SolveFirstHalf() int {
	actions := d.readActions()

	currentPosition := 50
	countZero := 0
	for _, action := range actions {
		if len(action) > 0 {
			direction := action[0]
			amount, err := strconv.Atoi(action[1:])
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}

			switch direction {
			case 'L':
				currentPosition -= amount
			case 'R':
				currentPosition += amount
			}

			currentPosition += 100
			currentPosition %= 100

			if currentPosition == 0 {
				countZero += 1
			}
		}
	}

	return countZero
}

func (d Day1) SolveSecondHalf() int {
	actions := d.readActions()

	currentPosition := 50
	countZero := 0
	for _, action := range actions {
		if len(action) > 0 {
			direction := action[0]
			amount, err := strconv.Atoi(action[1:])
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}

			switch direction {
			case 'L':
				countZero += amount / 100
				amount %= 100
				if currentPosition-amount <= 0 && currentPosition != 0 {
					countZero += 1
				}
				currentPosition -= amount
			case 'R':
				countZero += amount / 100
				amount %= 100
				if currentPosition+amount > 99 && currentPosition != 0 {
					countZero += 1
				}
				currentPosition += amount
			}

			currentPosition += 100
			currentPosition %= 100
		}
	}

	return countZero
}
