package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
	maxSize   = 100000

	//inputFile = "input0.txt"
	//maxSize   = 13
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

	// Load tiles
	var tiles []tile

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}

		tiles = append(tiles, tile{
			x: x,
			y: y,
		})
	}

	// Find max field
	var max float64
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles); j++ {
			if i == j {
				continue
			}

			a := tiles[i]
			b := tiles[j]

			f := (math.Abs(float64(a.x-b.x)) + 1) * (math.Abs(float64(a.y-b.y)) + 1)
			if f > max {
				max = f
			}
		}
	}

	fmt.Println(int(max))
}

func part2() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Load tiles
	var tiles []tile

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}

		tiles = append(tiles, tile{
			x: x,
			y: y,
		})
	}

	var board [][]byte
	// Initialize board
	for i := 0; i < maxSize; i++ {
		row := make([]byte, maxSize)
		for j := 0; j < len(row); j++ {
			row[j] = '.'
		}
		board = append(board, row)
	}

	// Place red tiles
	for _, t := range tiles {
		board[t.y][t.x] = '#'
	}

	// Draw horizontal lines
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles); j++ {
			if i == j {
				continue
			}

			a := tiles[i]
			b := tiles[j]

			if a.y != b.y {
				continue
			}

			if a.x > b.x {
				a, b = b, a
			}

			for k := a.x + 1; k < b.x; k++ {
				if board[a.y][k] != '#' {
					board[a.y][k] = 'X'
				}
			}
		}
	}

	// Draw vertical lines
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles); j++ {
			if i == j {
				continue
			}

			a := tiles[i]
			b := tiles[j]

			if a.x != b.x {
				continue
			}

			if a.y > b.y {
				a, b = b, a
			}

			for k := a.y + 1; k < b.y; k++ {
				if board[k][a.x] != '#' {
					board[k][a.x] = 'X'
				}
			}
		}
	}

	// Color dots inside
	for i := 0; i < len(board); i++ {

		var first int
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				first = j
				break
			}
		}

		var last int
		for j := len(board[0]) - 1; j >= 0; j-- {
			if board[i][j] != '.' {
				last = j
				break
			}
		}

		for j := first; j < last; j++ {
			if board[i][j] == '.' {
				board[i][j] = 'X'
			}
		}
	}

	// Draw board
	/*for i := 0; i < len(board); i++ {
		fmt.Println(string(board[i]))
	}*/

	// Find max field
	fmt.Println("Find max field now")

	var max float64
	for i := 0; i < len(tiles); i++ {
		fmt.Printf("%20d %d\n", int64(max), i)

		for j := 0; j < len(tiles); j++ {
			if i == j {
				continue
			}

			a := tiles[i]
			b := tiles[j]

			// Check if rectangle is legal
			xFrom := a.x
			xTo := b.x
			yFrom := a.y
			yTo := b.y

			if xFrom > xTo {
				xFrom, xTo = xTo, xFrom
			}

			if yFrom > yTo {
				yFrom, yTo = yTo, yFrom
			}

			legal := true
			for ii := yFrom; ii < yTo+1; ii++ {
				if board[ii][xFrom] == '.' {
					legal = false
					break
				}

				if board[ii][xTo] == '.' {
					legal = false
					break
				}
			}

			if legal {
				for jj := xFrom; jj < xTo+1; jj++ {
					if board[yFrom][jj] == '.' {
						legal = false
						break
					}

					if board[yTo][jj] == '.' {
						legal = false
						break
					}
				}
			}

			if !legal {
				continue
			}

			// Calculate max
			f := (math.Abs(float64(a.x-b.x)) + 1) * (math.Abs(float64(a.y-b.y)) + 1)
			if f > max {
				fmt.Println(a, b)
				max = f
			}
		}
	}

	fmt.Printf("%20d\n", int64(max))
}

type tile struct {
	x int
	y int
}
