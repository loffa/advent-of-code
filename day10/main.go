package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos [2]int

func main() {
	f, err := os.Open("./day10/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	rows := make([][]int, 0, 100)
	for s.Scan() {
		r := s.Text()
		row := []byte(strings.TrimSpace(r))
		pos := make([]int, 0, len(row))
		for _, b := range row {
			n, _ := strconv.Atoi(string(b))
			pos = append(pos, n)
		}
		rows = append(rows, pos)
	}

	res1 := part1(rows)
	res2 := part2(rows)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)

}

func checkInside(grid [][]int, pos Pos) bool {
	return pos[0] >= 0 && pos[0] < len(grid) && pos[1] >= 0 && pos[1] < len(grid[pos[0]])
}

func path(grid [][]int, current Pos, visited map[Pos]struct{}) int {
	if grid[current[0]][current[1]] == 9 {
		visited[current] = struct{}{}
		return 1
	}
	next := []Pos{
		{current[0] - 1, current[1]},
		{current[0] + 1, current[1]},
		{current[0], current[1] - 1},
		{current[0], current[1] + 1},
	}
	currentVal := grid[current[0]][current[1]]
	score := 0
	for _, pos := range next {
		if checkInside(grid, pos) && grid[pos[0]][pos[1]] == currentVal+1 {
			score += path(grid, pos, visited)
		}
	}
	return score
}

func part1(grid [][]int) int {
	res := 0

	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				m := make(map[Pos]struct{}, 10)
				path(grid, Pos{i, j}, m)
				res += len(m)
			}
		}
	}

	return res
}

func part2(grid [][]int) int {
	res := 0

	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				m := make(map[Pos]struct{}, 10)
				res += path(grid, Pos{i, j}, m)
			}
		}
	}

	return res
}
