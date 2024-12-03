package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)
	for s.Scan() {
		row := s.Text()
		parts := strings.Split(row, "   ")
		if len(parts) != 2 {
			log.Fatalln("Line was not two columns, was:", len(parts))
		}
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		left = append(left, a)
		right = append(right, b)
	}
	slices.Sort(left)
	slices.Sort(right)

	res1 := part1(left, right)
	res2 := part2(left, right)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func part1(left, right []int) int {
	res := 0
	for i, v := range left {
		res += int(math.Abs(float64(v - right[i])))
	}
	return res
}

func part2(left, right []int) int {
	counters := make(map[int]int)
	for _, v := range right {
		counters[v]++
	}

	res := 0
	for _, n := range left {
		res += n * counters[n]
	}
	return res
}
