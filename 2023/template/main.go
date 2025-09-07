package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./2023/day{{day}}/input_example.txt")
	if err != nil {
		log.Fatalln("could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)

	rows := make([]string, 0, 1000)
	for s.Scan() {
		rows = append(rows, s.Text())
	}
	if s.Err() != nil {
		log.Fatalln("could not read all lines:", err)
	}

	log.Println("Result (part 1):", part1())
	log.Println("Result (part 2):", part2())
}

func part1() int {
	panic("Not implemented!")
}

func part2() int {
	panic("Not implemented!")
}
