package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const free = -1

type Block struct {
	ID   int
	Size int
}

type FreeArea struct {
	Blocks []*Block
}

func main() {
	f, err := os.Open("./day9/input.txt")
	if err != nil {
		log.Fatalln("Could not open input file:", err)
	}

	defer func() {
		_ = f.Close()
	}()

	s := bufio.NewScanner(f)
	fs := make([]Block, 0, 1000)
	s.Scan()
	row := s.Text()
	fileID := 0
	for i, b := range row {
		n, _ := strconv.Atoi(string(b))
		if i%2 == 0 {
			for range n {
				fs = append(fs, Block{ID: fileID, Size: n})
			}
			fileID++
		} else {
			for range n {
				fs = append(fs, Block{ID: free, Size: n})
			}
		}
	}

	fsCopy := make([]Block, len(fs))
	copy(fsCopy, fs)
	res1 := part1(fsCopy)
	copy(fsCopy, fs)
	res2 := part2(fsCopy)
	log.Println("Result (part 1):", res1)
	log.Println("Result (part 2):", res2)
}

func checksum(fs []Block) int {
	res := 0
	for i, n := range fs {
		if n.ID == free {
			continue
		}
		res += n.ID * i
	}
	return res
}

func part1(fs []Block) int {
	lastInsert := 0
	end := len(fs) - 1
	for lastInsert < end {
		if fs[lastInsert].ID == free {
			for fs[end].ID == free {
				end--
			}
			fs[lastInsert] = fs[end]
			fs[end].ID = free
			end--
		}
		lastInsert++
	}

	return checksum(fs)
}

func part2(fs []Block) int {
	for i := len(fs) - 1; i >= 0; i-- {
		if fs[i].ID == free {
			continue
		}
		freeSpaceStart := findFreeSpace(fs, fs[i].Size)
		if freeSpaceStart == -1 || freeSpaceStart > i {
			continue
		}
		// Set file ID for all new blocks
		freeSpaceSize := fs[freeSpaceStart].Size
		for j := 0; j < freeSpaceSize; j++ {
			if j < fs[i].Size {
				fs[freeSpaceStart+j].ID = fs[i].ID
				fs[freeSpaceStart+j].Size = fs[i].Size
			} else {
				newSize := fs[freeSpaceStart+j].Size - fs[i].Size
				fs[freeSpaceStart+j].Size = newSize
			}
		}
		// Set free on old blocks
		for j := i; j > i-fs[i].Size; j-- {
			fs[j].ID = free
		}
	}
	return checksum(fs)
}

func findFreeSpace(fs []Block, size int) int {
	for i := 0; i < len(fs); i++ {
		if fs[i].ID == free && fs[i].Size >= size {
			return i
		}
	}
	return -1
}
