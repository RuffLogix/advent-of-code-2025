package day6

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Day6 struct {
}

func (d Day6) readOperations() []string {
	content, err := os.ReadFile("day6/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return strings.Split(string(content), "\n")
}

func (d Day6) SolveFirstHalf() int {
	contents := d.readOperations()

	numbers := make([][]int, len(contents)-2)
	operations := []string{}
	for i, content := range contents {
		texts := strings.Split(content, " ")

		for _, text := range texts {
			if len(text) == 0 {
				continue
			}

			if i == len(contents)-2 {
				operations = append(operations, text)
			} else {
				number, err := strconv.Atoi(text)
				if err != nil {
					log.Fatalf("Error to convert string to int: %v", err)
				}

				numbers[i] = append(numbers[i], number)
			}
		}
	}

	ans := 0
	for i := range operations {
		switch operations[i] {
		case "*":
			res := 1
			for j := range numbers {
				res *= numbers[j][i]
			}
			ans += res
		case "+":
			res := 0
			for j := range numbers {
				res += numbers[j][i]
			}
			ans += res
		}
	}

	return ans
}

func (d Day6) extendContents(contents []string) []string {
	ans := []string{}
	for _, content := range contents {
		content += strings.Repeat(" ", 5000-len(content))
		ans = append(ans, content)
	}
	return ans
}

func (d Day6) SolveSecondHalf() int {
	contents := d.readOperations()
	contents = d.extendContents(contents)

	ans := 0
	preOperation := 0
	for i, operation := range contents[len(contents)-2] {
		if operation != ' ' && i != 0 {
			var res int
			if contents[len(contents)-2][preOperation] == '*' {
				res = 1
			} else {
				res = 0
			}
			for pos := preOperation; pos < i-1; pos++ {
				num := 0
				for line := range len(contents) - 2 {
					if contents[line][pos] == ' ' {
						continue
					}
					num *= 10
					num += int(contents[line][pos]-'0') % 10
				}
				if contents[len(contents)-2][preOperation] == '*' {
					res *= num
				} else {
					res += num
				}
			}
			ans += res
			preOperation = i
		}
	}

	var res int
	if contents[len(contents)-2][preOperation] == '*' {
		res = 1
	} else {
		res = 0
	}
	for pos := preOperation; pos < len(contents[len(contents)-2]); pos++ {
		num := 0
		for line := range len(contents) - 2 {
			if contents[line][pos] == ' ' {
				continue
			}
			num *= 10
			num += int(contents[line][pos]-'0') % 10
		}
		if contents[len(contents)-2][preOperation] == '*' {
			res *= num
		} else {
			res += num
		}
	}
	ans += res

	return ans
}
