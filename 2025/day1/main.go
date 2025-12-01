package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	inputFile = "input.txt"

	startPos   = 50
	circleSize = 100
)

func main() {
	part1()
	part2()
}

func part1() {
	c := startPos
	zeros := 0

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		move := scanner.Text()

		steps, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Fatal(err)
		}

		direction := 1
		if move[0] == 'L' {
			direction = -1
		}

		c = mod(c+direction*steps, circleSize)
		if c == 0 {
			zeros++
		}

		//fmt.Printf("%s, c = %d\n", move, c)
	}

	//fmt.Println("c = ", c)
	fmt.Println("zeros = ", zeros)
}

func part2() {
	c := startPos
	zeros := 0

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		move := scanner.Text()

		steps, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Fatal(err)
		}

		direction := 1
		if move[0] == 'L' {
			direction = -1
		}

		for i := 0; i < steps; i++ {
			c = mod(c+direction, circleSize)
			if c == 0 {
				zeros++
			}
		}

		//fmt.Printf("%s, c = %d\n", move, c)
	}

	//fmt.Println("c = ", c)
	fmt.Println("zeros = ", zeros)
}

func mod(a, m int) int {
	return ((a % m) + m) % m
}
