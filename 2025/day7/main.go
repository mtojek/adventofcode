package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFile = "input.txt"

func main() {
	field := part1()
	part2(field)
}

func part1() [][]byte {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var field [][]byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		f := make([]byte, len(line))
		copy(f, line)

		field = append(field, f)
	}

	var total int
	for y := 1; y < len(field); y++ {

		// copy beams from the level above
		for x := 0; x < len(field[0]); x++ {
			if (field[y-1][x] == '|' || field[y-1][x] == 'S') && field[y][x] != '^' {
				field[y][x] = '|'
			}

			if field[y][x] == '^' {
				field[y][x-1] = '|'
				field[y][x+1] = '|'

				if field[y-1][x] == '|' {
					total++
				}
			}
		}

		fmt.Print(string(field[y]))
		fmt.Println("  ", total)
	}
	return field
}

func part2(field [][]byte) {
	prev := make([]int, len(field[0]))
	for x := 0; x < len(field[0]); x++ {
		if field[0][x] == 'S' {
			prev[x] = 1
		}
	}

	curr := make([]int, len(field[0]))
	for y := 1; y < len(field); y++ {
		for x := 0; x < len(curr); x++ {
			if x-1 >= 0 && field[y][x-1] == '^' {
				curr[x] += prev[x-1]
			}
			if x+1 < len(field[0]) && field[y][x+1] == '^' {
				curr[x] += prev[x+1]
			}
			if field[y][x] == '|' {
				curr[x] += prev[x]
			}
		}

		copy(prev, curr)
		curr = make([]int, len(curr))
	}

	var paths int
	for x := 0; x < len(prev); x++ {
		paths += prev[x]
	}
	fmt.Println(paths)
}
