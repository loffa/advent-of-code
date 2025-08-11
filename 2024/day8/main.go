package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Pos [2]int

func main() {
	f, err := os.Open("./2024/day8/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	rows := make([][]byte, 0, 10)
	for s.Scan() {
		r := s.Text()
		row := []byte(strings.TrimSpace(r))
		rows = append(rows, row)
	}

	antennas := make(map[byte][]Pos)
	for x, row := range rows {
		for y, col := range row {
			if col == '.' {
				continue
			}
			if _, ok := antennas[col]; !ok {
				antennas[col] = make([]Pos, 0, 10)
			}
			antennas[col] = append(antennas[col], Pos{x, y})
		}
	}

	res1 := part1(rows, antennas)
	res2 := part2(rows, antennas)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func checkInside(grid [][]byte, pos Pos) bool {
	return pos[0] >= 0 && pos[0] < len(grid) && pos[1] >= 0 && pos[1] < len(grid[pos[0]])
}

func part1(grid [][]byte, antennas map[byte][]Pos) int {
	nodes := make(map[Pos]struct{})

	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := 0; j < len(positions); j++ {
				if i == j {
					continue
				}

				dx := positions[j][0] - positions[i][0]
				dy := positions[j][1] - positions[i][1]

				node := Pos{positions[i][0] + dx*2, positions[i][1] + dy*2}
				if checkInside(grid, node) {
					nodes[node] = struct{}{}
				}
			}
		}
	}

	return len(nodes)
}

func part2(grid [][]byte, antennas map[byte][]Pos) int {
	nodes := make(map[Pos]struct{})

	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := 0; j < len(positions); j++ {
				if i == j {
					continue
				}

				dx := positions[j][0] - positions[i][0]
				dy := positions[j][1] - positions[i][1]

				for k := -len(grid); k <= len(grid); k++ {
					node := Pos{positions[i][0] + dx*k, positions[i][1] + dy*k}
					if checkInside(grid, node) {
						nodes[node] = struct{}{}
					}
				}
			}
		}
	}

	return len(nodes)
}
