package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	re1 = regexp.MustCompile(`mul\(\d+,\d+\)`)
	re2 = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
)

func main() {
	f, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	rows := make([]string, 0, 10)
	for s.Scan() {
		row := s.Text()
		rows = append(rows, row)
	}

	res1 := part1(rows)
	res2 := part2(rows)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func part1(rows []string) int {
	sum := 0
	for _, row := range rows {
		matches := re1.FindAllString(row, -1)

		for _, m := range matches {
			n, _ := strings.CutPrefix(m, "mul(")
			n, _ = strings.CutSuffix(n, ")")
			nums := strings.Split(n, ",")
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])
			sum += a * b
		}
	}
	return sum
}

func part2(rows []string) int {
	sum := 0
	enabled := true
	for _, row := range rows {
		matches := re2.FindAllString(row, -1)

		for _, m := range matches {
			if m == "do()" {
				enabled = true
			} else if m == "don't()" {
				enabled = false
			} else if enabled {
				n, _ := strings.CutPrefix(m, "mul(")
				n, _ = strings.CutSuffix(n, ")")
				nums := strings.Split(n, ",")
				a, _ := strconv.Atoi(nums[0])
				b, _ := strconv.Atoi(nums[1])

				sum += a * b
			}
		}
	}
	return sum
}
