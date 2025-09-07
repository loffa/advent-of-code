package main

import (
	"bufio"
	"log"
	"os"
)

type Pos [2]int

func main() {
	f, err := os.Open("./2024/day12/input_example.txt")
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

func part1(plots [][]byte) int {
	res := 0

	return res
}

func part2() int {
	res := 0

	return res
}
