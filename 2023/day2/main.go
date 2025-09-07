package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID            int
	MaxRedCount   int
	MaxGreenCount int
	MaxBlueCount  int
}

func (g *Game) Validate() bool {
	return g.MaxRedCount <= 12 && g.MaxGreenCount <= 13 && g.MaxBlueCount <= 14
}

func (g *Game) Power() int {
	return g.MaxRedCount * g.MaxGreenCount * g.MaxBlueCount
}

var pickRegexp = regexp.MustCompile("(?P<count>[0-9]+) (?P<color>[a-zA-Z]+)")

func main() {
	f, err := os.Open("./2023/day2/input.txt")
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
	sum := 0
	for _, row := range rows {
		gameInfo := strings.FieldsFunc(row, func(r rune) bool {
			switch r {
			case ':', ';':
				return true
			default:
				return false
			}
		})

		g := &Game{}
		for i, part := range gameInfo {
			if i == 0 {
				id, _ := strconv.Atoi(strings.SplitAfter(part, " ")[1])
				g.ID = id
				continue
			}

			diceCount := pickRegexp.FindAllStringSubmatch(part, -1)

			for _, match := range diceCount {
				num, _ := strconv.Atoi(match[1])
				color := match[2]

				if color == "red" && num > g.MaxRedCount {
					g.MaxRedCount = num
				} else if color == "green" && num > g.MaxGreenCount {
					g.MaxGreenCount = num
				} else if color == "blue" && num > g.MaxBlueCount {
					g.MaxBlueCount = num
				}
			}
		}

		if g.Validate() {
			sum += g.ID
		}
	}

	return sum
}

func part2(rows []string) int {
	sum := 0
	for _, row := range rows {
		gameInfo := strings.FieldsFunc(row, func(r rune) bool {
			switch r {
			case ':', ';':
				return true
			default:
				return false
			}
		})

		g := &Game{}
		for i, part := range gameInfo {
			if i == 0 {
				id, _ := strconv.Atoi(strings.SplitAfter(part, " ")[1])
				g.ID = id
				continue
			}

			diceCount := pickRegexp.FindAllStringSubmatch(part, -1)

			for _, match := range diceCount {
				num, _ := strconv.Atoi(match[1])
				color := match[2]

				if color == "red" && num > g.MaxRedCount {
					g.MaxRedCount = num
				} else if color == "green" && num > g.MaxGreenCount {
					g.MaxGreenCount = num
				} else if color == "blue" && num > g.MaxBlueCount {
					g.MaxBlueCount = num
				}
			}
		}

		sum += g.Power()
	}

	return sum
}
