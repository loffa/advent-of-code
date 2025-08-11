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
	f, err := os.Open("./2024/day11/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	stones := make([]int, 0, 100)
	s.Scan()
	parts := strings.Fields(s.Text())
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		stones = append(stones, n)
	}

	res1 := part1(stones)
	res2 := part2(stones)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)

}

func blink(stones []int) []int {
	newStones := make([]int, 0, len(stones))
	for _, s := range stones {
		if s == 0 {
			newStones = append(newStones, 1)
		} else if num := strconv.Itoa(s); len(num)%2 == 0 {
			split1, split2 := num[:len(num)/2], num[len(num)/2:]
			s, _ = strconv.Atoi(split1)
			s2, _ := strconv.Atoi(split2)
			newStones = append(newStones, s, s2)
		} else {
			newStones = append(newStones, s*2024)
		}
	}
	return newStones
}

func part1(stones []int) int {
	for range 25 {
		stones = blink(stones)
	}

	return len(stones)
}

func part2(stones []int) int {
	stoneResults := make(map[int]int)
	for _, s := range stones {
		stoneResults[s] = 1
	}

	for range 75 {
		newResults := make(map[int]int)
		for s, count := range stoneResults {
			localStones := blink([]int{s})
			for _, ls := range localStones {
				newResults[ls] += count
			}
		}
		stoneResults = newResults
	}

	res := 0
	for _, count := range stoneResults {
		res += count
	}

	return res
}
