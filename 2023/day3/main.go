package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

type coordinate [2]int

func main() {
	f, err := os.Open("./2023/day3/input.txt")
	if err != nil {
		log.Fatalln("could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)

	rows := make([]string, 0, 1000)
	for s.Scan() {
		rows = append(rows, s.Text())
	}
	if s.Err() != nil {
		log.Fatalln("could not read all lines:", err)
	}

	log.Println("Result (part 1):", part1(rows))
	log.Println("Result (part 2):", part2(rows))
}

func part1(rows []string) int {
	symbols := mapSymbols(rows)
	numbers := getNumbers(rows)

	sum := 0
	for _, n := range numbers {
		neighbours := make([]coordinate, 0, 4)

		str := ""
		for i, pos := range n {
			str += string(rows[pos[0]][pos[1]])
			neighbours = append(neighbours, calcNeighbours(i, pos, len(n)-1)...)
		}

		if ok, _, _ := validate(neighbours, symbols); ok {
			num, _ := strconv.Atoi(str)
			sum += num
		}
	}
	return sum
}

func part2(rows []string) int {
	sum := 0
	symbols := mapSymbols(rows)
	numbers := getNumbers(rows)

	starNumbers := make(map[coordinate][]int, 10)

	for _, n := range numbers {
		neighbours := make([]coordinate, 0, 4)

		str := ""
		for i, pos := range n {
			str += string(rows[pos[0]][pos[1]])
			neighbours = append(neighbours, calcNeighbours(i, pos, len(n)-1)...)
		}

		if ok, star, starCoordinate := validate(neighbours, symbols); ok && star {
			arr, ok := starNumbers[starCoordinate]
			if !ok {
				arr = make([]int, 0, 2)
			}

			nInt, _ := strconv.Atoi(str)
			starNumbers[starCoordinate] = append(arr, nInt)
		}
	}

	for _, starNumber := range starNumbers {
		if len(starNumber) == 1 {
			continue
		}

		val := 1
		for _, n := range starNumber {
			val = val * n
		}
		sum += val
	}

	return sum
}

func getNumbers(rows []string) [][]coordinate {
	digits := make([]coordinate, 0, 100)
	for i, row := range rows {
		for j, c := range row {
			if unicode.IsDigit(c) {
				digits = append(digits, coordinate{i, j})
			}
		}
	}

	numbers := make([][]coordinate, 0, 100)
	cur := make([]coordinate, 0, 3)

	for i, d := range digits {
		cur = append(cur, d)

		if i+1 == len(digits) {
			numbers = append(numbers, cur)
			cur = []coordinate{}
			continue
		}

		next := digits[i+1]
		if d[0] != next[0] || d[1]+1 != next[1] {
			numbers = append(numbers, cur)
			cur = []coordinate{}
		}
	}
	return numbers
}

func validate(neighbours []coordinate, symbols map[coordinate]rune) (bool, bool, coordinate) {
	var (
		valid          = false
		isStar         = false
		starCoordinate = coordinate{}
	)
	for _, ne := range neighbours {
		if s, ok := symbols[ne]; ok {
			valid = true
			if s == '*' {
				isStar = true
				starCoordinate = ne
			}
			break
		}
	}
	return valid, isStar, starCoordinate
}

func calcNeighbours(i int, curPos coordinate, maxWidth int) []coordinate {
	neighbours := make([]coordinate, 0, 8)
	if i == 0 {
		// If first, add left side and diagonal
		neighbours = append(neighbours, coordinate{curPos[0], curPos[1] - 1})
		neighbours = append(neighbours, coordinate{curPos[0] - 1, curPos[1] - 1})
		neighbours = append(neighbours, coordinate{curPos[0] + 1, curPos[1] - 1})
	}
	if i == maxWidth {
		// If last, add right side and diagonal
		neighbours = append(neighbours, coordinate{curPos[0], curPos[1] + 1})
		neighbours = append(neighbours, coordinate{curPos[0] - 1, curPos[1] + 1})
		neighbours = append(neighbours, coordinate{curPos[0] + 1, curPos[1] + 1})
	}
	// Always add top/bottom
	neighbours = append(neighbours, coordinate{curPos[0] - 1, curPos[1]})
	neighbours = append(neighbours, coordinate{curPos[0] + 1, curPos[1]})

	return neighbours
}

func mapSymbols(rows []string) map[coordinate]rune {
	m := make(map[coordinate]rune, 1000)
	for i, row := range rows {
		for j, c := range row {
			if !unicode.IsDigit(c) && c != '.' {
				m[coordinate{i, j}] = c
			}
		}
	}
	return m
}
