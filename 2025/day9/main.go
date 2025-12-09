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

type tile struct {
	x int
	y int
}

func part2() {
}
