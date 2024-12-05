package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	a, b int
}

type Rules []Rule

func (rs Rules) Map() map[int][]int {
	res := make(map[int][]int, len(rs))
	for _, r := range rs {
		if _, ok := res[r.a]; ok {
			res[r.a] = append(res[r.a], r.b)
		} else {
			res[r.a] = []int{r.b}
		}
	}
	return res
}

func main() {
	f, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	rules := make(Rules, 0, 100)
	updates := make([][]int, 0, 100)
	for s.Scan() {
		row := s.Text()
		if len(row) == 0 {
			continue
		} else if strings.Contains(row, "|") {
			parts := strings.Split(row, "|")
			if len(parts) != 2 {
				continue
			}
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			rules = append(rules, Rule{a: a, b: b})
		} else if strings.Contains(row, ",") {
			parts := strings.Split(row, ",")
			update := make([]int, 0, len(parts))
			for _, v := range parts {
				x, _ := strconv.Atoi(v)
				update = append(update, x)
			}
			updates = append(updates, update)
		}
	}

	res1 := part1(updates, rules)
	res2 := part2(updates, rules)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func validate(up []int, rules []Rule) bool {
	for _, rule := range rules {
		aPos, bPos := -1, -1
		for i, val := range up {
			if val == rule.a {
				aPos = i
			} else if val == rule.b {
				bPos = i
			}
		}
		if aPos != -1 && bPos != -1 && aPos > bPos {
			return false
		}
	}
	return true
}

func part1(updates [][]int, rules Rules) int {
	res := 0
	for _, up := range updates {
		if validate(up, rules) {
			mid := up[len(up)/2]
			res += mid
		}
	}
	return res
}

func part2(updates [][]int, rules Rules) int {
	res := 0
	ruleMap := rules.Map()
	for _, up := range updates {
		if validate(up, rules) {
			continue
		}

		slices.SortFunc(up, func(a, b int) int {
			if a == b {
				return 0
			}
			for _, n := range ruleMap[a] {
				if b == n {
					return -1
				}
			}
			return 1
		})
		mid := up[len(up)/2]
		res += mid
	}
	return res
}
