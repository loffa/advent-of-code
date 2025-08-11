package main

import (
	"bufio"
	"log"
	"os"
)

type (
	Direction struct {
		X, Y int
	}
	Pos     [2]int
	Visited struct {
		pos Pos
		dir Direction
	}
)

func main() {
	f, err := os.Open("./2024/day6/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	startPosition := [2]int{0, 0}
	points := make([][]byte, 0, 130)
	for s.Scan() {
		r := s.Text()
		row := []byte(r)
		for i, b := range row {
			if b == '^' {
				startPosition = Pos{len(points), i}
				break
			}
		}
		points = append(points, row)
	}

	res1 := part1(startPosition, points)
	res2 := part2(startPosition, points)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func part1(startPos Pos, points [][]byte) int {
	res := 0
	currentPos := startPos
	currentDir := Direction{-1, 0}

	for checkInside(points, currentPos) {
		newTile := false
		currentPos, currentDir, newTile = move(currentPos, currentDir, points)
		if newTile {
			res++
		}
	}

	return res
}

func move(currentPos Pos, currentDir Direction, points [][]byte) (Pos, Direction, bool) {
	newPos := Pos{currentPos[0] + currentDir.X, currentPos[1] + currentDir.Y}
	if checkInside(points, newPos) && points[newPos[0]][newPos[1]] == '#' {
		tmp := currentDir.Y
		currentDir.Y = currentDir.X * -1
		currentDir.X = tmp
		newPos = currentPos
	}
	if points[currentPos[0]][currentPos[1]] != 'X' {
		points[currentPos[0]][currentPos[1]] = 'X'
		return newPos, currentDir, true
	}
	return newPos, currentDir, false
}

func checkInside(points [][]byte, pos Pos) bool {
	return pos[0] >= 0 && pos[0] < len(points) && pos[1] >= 0 && pos[1] < len(points[pos[0]])
}

func part2(startPos Pos, points [][]byte) int {
	res := 0
	currentPos := startPos
	currentDir := Direction{-1, 0}
	visitMap := make(map[Visited]struct{})

	for checkInside(points, currentPos) {
		currentPos, currentDir, _ = move(currentPos, currentDir, points)
		visitMap[Visited{
			pos: currentPos,
			dir: Direction{0, 0},
		}] = struct{}{}
	}

	for v, _ := range visitMap {
		if v.pos[0] >= len(points) || v.pos[1] >= len(points[v.pos[0]]) {
			continue
		}
		points[v.pos[0]][v.pos[1]] = '#'
		newMap := make(map[Visited]struct{})
		currentPos = startPos
		currentDir = Direction{-1, 0}
		for checkInside(points, currentPos) {
			currentPos, currentDir, _ = move(currentPos, currentDir, points)
			if _, ok := newMap[Visited{pos: currentPos, dir: currentDir}]; ok {
				res++
				points[v.pos[0]][v.pos[1]] = '0'
				break
			} else {
				newMap[Visited{pos: currentPos, dir: currentDir}] = struct{}{}
			}
		}
		if points[v.pos[0]][v.pos[1]] != '0' {
			points[v.pos[0]][v.pos[1]] = '.'
		}
	}

	return res
}
