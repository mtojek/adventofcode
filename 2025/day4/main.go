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

	var inputRolls [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		roll := scanner.Bytes()

		r := make([]byte, len(roll))
		copy(r, roll)
		inputRolls = append(inputRolls, r)
	}

	rolls := withMargin(inputRolls)

	var sum int
	for y := 1; y < len(rolls)-1; y++ {
		roll := rolls[y]
		for x := 1; x < len(roll)-1; x++ {
			if roll[x] != '@' {
				continue
			}

			nw := 0
			if rolls[y-1][x-1] == '@' {
				nw++
			}

			n := 0
			if rolls[y-1][x] == '@' {
				n++
			}

			ne := 0
			if rolls[y-1][x+1] == '@' {
				ne++
			}

			w := 0
			if rolls[y][x-1] == '@' {
				w++
			}

			e := 0
			if rolls[y][x+1] == '@' {
				e++
			}

			sw := 0
			if rolls[y+1][x-1] == '@' {
				sw++
			}

			s := 0
			if rolls[y+1][x] == '@' {
				s++
			}

			se := 0
			if rolls[y+1][x+1] == '@' {
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
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var inputRolls [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		roll := scanner.Bytes()

		r := make([]byte, len(roll))
		copy(r, roll)
		inputRolls = append(inputRolls, r)
	}

	rolls := withMargin(inputRolls)

	var sum int
	for {
		freed := accessRolls(rolls)
		if len(freed) > 0 {
			sum += len(freed)
		} else {
			break
		}

		rolls = freeRolls(rolls, freed)
	}
	fmt.Println(sum)
}

func withMargin(inputRolls [][]byte) [][]byte {
	var rolls [][]byte
	zeros1 := make([]byte, len(inputRolls[0])+2)
	for i := 0; i < len(zeros1); i++ {
		zeros1[i] = '.'
	}
	rolls = append(rolls, zeros1)
	for _, roll := range inputRolls {
		row := []byte{'.'}
		row = append(row, roll...)
		row = append(row, '.')
		rolls = append(rolls, row)
	}
	zeros2 := make([]byte, len(inputRolls[0])+2)
	for i := 0; i < len(zeros2); i++ {
		zeros2[i] = '.'
	}
	rolls = append(rolls, zeros2)
	return rolls
}

type rollCoord struct {
	x int
	y int
}

func accessRolls(rolls [][]byte) []rollCoord {
	var toFree []rollCoord
	for y := 1; y < len(rolls)-1; y++ {
		roll := rolls[y]
		for x := 1; x < len(roll)-1; x++ {
			if roll[x] != '@' {
				continue
			}

			nw := 0
			if rolls[y-1][x-1] == '@' {
				nw++
			}

			n := 0
			if rolls[y-1][x] == '@' {
				n++
			}

			ne := 0
			if rolls[y-1][x+1] == '@' {
				ne++
			}

			w := 0
			if rolls[y][x-1] == '@' {
				w++
			}

			e := 0
			if rolls[y][x+1] == '@' {
				e++
			}

			sw := 0
			if rolls[y+1][x-1] == '@' {
				sw++
			}

			s := 0
			if rolls[y+1][x] == '@' {
				s++
			}

			se := 0
			if rolls[y+1][x+1] == '@' {
				se++
			}

			total := nw + n + ne + w + e + sw + s + se
			if total < 4 {
				toFree = append(toFree, rollCoord{
					x: x,
					y: y,
				})
			}
		}
	}
	return toFree
}

func freeRolls(rolls [][]byte, coords []rollCoord) [][]byte {
	for _, coord := range coords {
		rolls[coord.y][coord.x] = '.'
	}
	return rolls
}
