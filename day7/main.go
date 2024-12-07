package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Calibration struct {
	Sum   int
	Parts []int
}

func main() {
	f, err := os.Open("./day7/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	calibrations := make([]*Calibration, 0, 1000)
	for s.Scan() {
		r := s.Text()
		parts := strings.Split(r, ":")
		if len(parts) != 2 {
			log.Fatalln("Line did not contain sum and parts")
		}
		sum, _ := strconv.Atoi(parts[0])

		p2 := strings.Split(strings.TrimSpace(parts[1]), " ")
		c := &Calibration{
			Sum:   sum,
			Parts: make([]int, 0, len(p2)),
		}
		for _, v := range p2 {
			n, _ := strconv.Atoi(v)
			c.Parts = append(c.Parts, n)
		}
		calibrations = append(calibrations, c)
	}

	res1 := part1(calibrations)
	res2 := part2(calibrations)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func validate(target int, parts []int, v2 bool) bool {
	// Exit early cases, no parts left or sum is already too big
	if len(parts) == 1 || parts[0] > target {
		return parts[0] == target
	}
	concat, _ := strconv.Atoi(strconv.Itoa(parts[0]) + strconv.Itoa(parts[1]))
	return validate(target, append([]int{parts[0] + parts[1]}, parts[2:]...), v2) ||
		validate(target, append([]int{parts[0] * parts[1]}, parts[2:]...), v2) ||
		(v2 && validate(target, append([]int{concat}, parts[2:]...), v2))
}

func part1(calibrations []*Calibration) int {
	res := 0

	for _, c := range calibrations {
		if validate(c.Sum, c.Parts, false) {
			res += c.Sum
		}
	}

	return res
}

func part2(calibrations []*Calibration) int {
	res := 0

	for _, c := range calibrations {
		if validate(c.Sum, c.Parts, true) {
			res += c.Sum
		}
	}

	return res
}
