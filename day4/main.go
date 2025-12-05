package day4

import (
	"log"
	"os"
	"strings"
)

type Day4 struct {
}

func (d Day4) readPaperRolls() []string {
	content, err := os.ReadFile("day4/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return strings.Split(string(content), "\n")
}

var dirX = []int{1, 1, -1, -1, 0, -1, 0, 1}
var dirY = []int{1, -1, 1, -1, 1, 0, -1, 0}

func (d Day4) validCell(i int, j int, paperRolls *[]string) bool {
	if i < 0 || j < 0 || i >= len(*paperRolls) || j >= len((*paperRolls)[i]) || (*paperRolls)[i][j] != '@' {
		return false
	}
	return true
}

func (d Day4) countAdjacentCells(i int, j int, paperRolls *[]string) int {
	cnt := 0
	for index := range dirX {
		if d.validCell(i+dirX[index], j+dirY[index], paperRolls) {
			cnt += 1
		}
	}
	return cnt
}

func (d Day4) SolveFirstHalf() int {
	paperRolls := d.readPaperRolls()

	ans := 0
	for i := range paperRolls {
		for j := range paperRolls[i] {
			if paperRolls[i][j] == '@' && d.countAdjacentCells(i, j, &paperRolls) < 4 {
				ans += 1
			}
		}
	}

	return ans
}

func (d Day4) SolveSecondHalf() int {
	paperRolls := d.readPaperRolls()

	ans := 0
	for {
		isChanged := false
		for i := range paperRolls {
			for j := range paperRolls[i] {
				if paperRolls[i][j] == '@' && d.countAdjacentCells(i, j, &paperRolls) < 4 {
					row := []rune(paperRolls[i])
					row[j] = 'x'
					paperRolls[i] = string(row)

					isChanged = true
					ans += 1
				}
			}
		}
		if !isChanged {
			break
		}
	}

	return ans
}
