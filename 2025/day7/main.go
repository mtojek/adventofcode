package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input.txt"
)

func main() {
	part1()
	part2()
}

func part1() {
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
}

func part2() {
}
