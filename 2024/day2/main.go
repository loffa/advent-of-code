package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./2024/day2/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)

	reports := make([][]int, 0, 1000)
	for s.Scan() {
		row := s.Text()
		reports = append(reports, toLevels(row))
	}

	res1 := part1(reports)
	res2 := part2(reports)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func part1(reports [][]int) int {
	numSafe := 0
	for _, row := range reports {
		if isSafe(row) {
			numSafe++
		}
	}
	return numSafe
}

func isSafe(row []int) bool {
	lastDiff := 0
	for i := 0; i < len(row)-1; i++ {
		a, b := row[i], row[i+1]
		if a == b {
			return false
		}
		diff := b - a
		if math.Abs(float64(diff)) > 3 || math.Abs(float64(diff)) < 1 {
			return false
		}
		if (lastDiff < 0 && diff > 0) || (lastDiff > 0 && diff < 0) {
			return false
		}
		lastDiff = diff
	}
	return true
}

func part2(reports [][]int) int {
	numSafe := 0
	for _, row := range reports {
		if isSafe(row) {
			numSafe++
			continue
		}
		for i := range row {
			cp := make([]int, len(row))
			copy(cp, row)
			if isSafe(append(cp[:i], cp[i+1:]...)) {
				numSafe++
				break
			}
		}
	}
	return numSafe
}

func toLevels(s string) []int {
	parts := strings.Split(s, " ")
	res := make([]int, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		res = append(res, n)
	}
	return res
}
