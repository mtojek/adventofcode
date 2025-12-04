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

	var rolls [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		roll := scanner.Bytes()

		r := make([]byte, len(roll))
		copy(r, roll)
		rolls = append(rolls, r)
	}

	var rollsWithMargin [][]byte

	zeros1 := make([]byte, len(rolls[0])+2)
	for i := 0; i < len(zeros1); i++ {
		zeros1[i] = '.'
	}
	rollsWithMargin = append(rollsWithMargin, zeros1)
	for _, roll := range rolls {
		row := []byte{'.'}
		row = append(row, roll...)
		row = append(row, '.')
		rollsWithMargin = append(rollsWithMargin, row)
	}
	zeros2 := make([]byte, len(rolls[0])+2)
	for i := 0; i < len(zeros2); i++ {
		zeros2[i] = '.'
	}
	rollsWithMargin = append(rollsWithMargin, zeros2)

	var sum int
	for y := 1; y < len(rollsWithMargin)-1; y++ {
		roll := rollsWithMargin[y]
		for x := 1; x < len(roll)-1; x++ {
			if roll[x] != '@' {
				continue
			}

			nw := 0
			if rollsWithMargin[y-1][x-1] == '@' {
				nw++
			}

			n := 0
			if rollsWithMargin[y-1][x] == '@' {
				n++
			}

			ne := 0
			if rollsWithMargin[y-1][x+1] == '@' {
				ne++
			}

			w := 0
			if rollsWithMargin[y][x-1] == '@' {
				w++
			}

			e := 0
			if rollsWithMargin[y][x+1] == '@' {
				e++
			}

			sw := 0
			if rollsWithMargin[y+1][x-1] == '@' {
				sw++
			}

			s := 0
			if rollsWithMargin[y+1][x] == '@' {
				s++
			}

			se := 0
			if rollsWithMargin[y+1][x+1] == '@' {
				se++
			}

			total := nw + n + ne + w + e + sw + s + se
			if total < 4 {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func part2() {
}
