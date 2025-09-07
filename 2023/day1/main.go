package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numberMap = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func main() {
	f, err := os.Open("./2023/day1/input_1.txt")
	if err != nil {
		log.Fatalln("could not open input file:", err)
	}

	f2, err := os.Open("./2023/day1/input_2.txt")
	if err != nil {
		log.Fatalln("could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
		_ = f2.Close()
	}()

	s := bufio.NewScanner(f)
	s2 := bufio.NewScanner(f2)

	rows := make([]string, 0, 1000)
	for s.Scan() {
		rows = append(rows, s.Text())
	}
	if s.Err() != nil {
		log.Fatalln("could not read all 1st lines:", err)
	}

	rows2 := make([]string, 0, 1000)
	for s2.Scan() {
		rows2 = append(rows2, s2.Text())
	}
	if s.Err() != nil {
		log.Fatalln("could not read all 2nd lines:", err)
	}

	log.Println("Result (part 1):", part1(rows))
	log.Println("Result (part 2):", part2(rows2))
}

func part1(rows []string) int {
	sum := 0

	for _, l := range rows {
		first, last := 0, 0
		foundFirst := false
		for _, c := range l {
			if unicode.IsDigit(c) {
				if !foundFirst {
					first, _ = strconv.Atoi(string(c))
					foundFirst = true
				}

				last, _ = strconv.Atoi(string(c))
			}
		}
		sum += first*10 + last
	}

	return sum
}

func part2(rows []string) int {
	sum := 0

	for _, l := range rows {
		for text, digit := range numberMap {
			if strings.Contains(l, text) {
				l = strings.ReplaceAll(l, text, digit)
			}
		}
		first, last := 0, 0
		foundFirst := false
		for _, c := range l {
			if unicode.IsDigit(c) {
				if !foundFirst {
					first, _ = strconv.Atoi(string(c))
					foundFirst = true
				}

				last, _ = strconv.Atoi(string(c))
			}
		}
		sum += first*10 + last
	}

	return sum
}
