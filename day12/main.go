package main

import (
	"bufio"
	"log"
	"os"
)

type Pos [2]int

func main() {
	f, err := os.Open("./day12/input_example.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	plots := make([][]byte, 0, 100)
	for s.Scan() {
		row := s.Text()
		plots = append(plots, []byte(row))
	}

	res1 := part1(plots)
	res2 := part2()
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)

}

func checkInside(points [][]byte, pos Pos) bool {
	return pos[0] >= 0 && pos[0] < len(points) && pos[1] >= 0 && pos[1] < len(points[pos[0]])
}

func explore(plots [][]byte, curr Pos, visited map[Pos]bool) (int, int) {
	next := []Pos{
		{curr[0] - 1, curr[1]},
		{curr[0] + 1, curr[1]},
		{curr[0], curr[1] - 1},
		{curr[0], curr[1] + 1},
	}

	numbNeighbours, area := 0, 0
	for _, p := range next {
		if !checkInside(plots, p) {
			continue
		}
		if plots[p[0]][p[1]] != plots[curr[0]][curr[1]] {
			continue
		}
		numbNeighbours++
		if !visited[p] {
			newArea, newNumNeighbours := explore(plots, p, visited)
			area += newArea
			numbNeighbours += newNumNeighbours
		}
	}

	return area, 4 - numbNeighbours
}

func part1(plots [][]byte) int {
	res := 0

	visited := make(map[Pos]bool)
	for row := range plots {
		for col := range row {
			curr := Pos{row, col}
			if _, ok := visited[curr]; !ok {
				area, perimeter := explore(plots, curr, visited)
				res += (area + 1) * perimeter
			}
		}
	}

	return res
}

func part2() int {
	res := 0

	return res
}
