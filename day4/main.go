package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	rows := make([][]int32, 0, 1000)
	for s.Scan() {
		row := s.Text()
		rows = append(rows, []int32(row))
	}

	res1 := part1(rows)
	res2 := part2(rows)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func part1(rows [][]int32) interface{} {
	res := 0
	for x := range rows {
		for y, c := range rows[x] {
			if c == 'X' {
				res += findInDirection(rows, x, y, 1, 0, "MAS")
				res += findInDirection(rows, x, y, -1, 0, "MAS")
				res += findInDirection(rows, x, y, 0, 1, "MAS")
				res += findInDirection(rows, x, y, 0, -1, "MAS")
				res += findInDirection(rows, x, y, -1, -1, "MAS")
				res += findInDirection(rows, x, y, 1, -1, "MAS")
				res += findInDirection(rows, x, y, -1, 1, "MAS")
				res += findInDirection(rows, x, y, 1, 1, "MAS")
			}
		}
	}
	return res
}

func findInDirection(rows [][]int32, startX, startY, dirX, dirY int, word string) int {
	x := startX + dirX
	y := startY + dirY
	for _, l := range word {
		if x >= len(rows) || y >= len(rows[0]) || x < 0 || y < 0 {
			return 0
		}
		if rows[x][y] != l {
			return 0
		}
		x += dirX
		y += dirY
	}
	return 1
}

func part2(rows [][]int32) interface{} {
	res := 0
	for x := range rows {
		for y, l := range rows[x] {
			if l == 'A' {
				b := 0

				// Top left
				tl := findInDirection(rows, x, y, -1, -1, "M")
				if tl == 1 {
					b += findInDirection(rows, x, y, 1, 1, "S")
				}

				// Top right
				tr := findInDirection(rows, x, y, -1, 1, "M")
				if tr == 1 {
					b += findInDirection(rows, x, y, 1, -1, "S")
				}

				// Lower right
				lr := findInDirection(rows, x, y, 1, 1, "M")
				if lr == 1 {
					b += findInDirection(rows, x, y, -1, -1, "S")
				}

				// Lower left
				ll := findInDirection(rows, x, y, 1, -1, "M")
				if ll == 1 {
					b += findInDirection(rows, x, y, -1, 1, "S")
				}

				if b == 2 {
					res++
				}
			}
		}
	}
	return res
}
