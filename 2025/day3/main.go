package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input.txt"

	maxBatteriesUsed = 12
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

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bank := scanner.Text()
		var batteries []int
		for _, c := range bank {
			batteries = append(batteries, int(c)-0x30)
		}

		var max int
		for i := 0; i < len(batteries)-1; i++ {
			for j := i + 1; j < len(batteries); j++ {
				t := batteries[i]*10 + batteries[j]
				if t > max {
					max = t
				}
			}
		}

		sum += max
	}

	fmt.Println(sum)
}

func part2() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bank := scanner.Text()
		var batteries []int
		for _, c := range bank {
			batteries = append(batteries, int(c)-0x30)
		}
		max := next(batteries, 0, 0)
		sum += max
		fmt.Println(max)
	}

	fmt.Println(sum)
}

func next(batteries []int, used int, t int) int {
	if used == maxBatteriesUsed {
		return t
	}

	if len(batteries) == 0 {
		return 0 // we don't have more batteries to use
	}

	for c := 9; c >= 1; c-- {
		for i, b := range batteries {
			if b == c {
				max := next(batteries[i+1:], used+1, 10*t+b)
				if max > 0 {
					return max
				}
			}
		}
	}

	return 0 // all batteries in the range checked
}
