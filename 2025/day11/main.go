package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "input1.txt"
)

var (
	paths         = map[string]int{}
	numberOfPaths func(deviceName string) int
)

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	devices := map[string][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		s := strings.SplitN(line, ": ", 2)
		deviceName := s[0]
		outputs := strings.Split(s[1], " ")

		devices[deviceName] = outputs
	}

	numberOfPaths = func(deviceName string) int {
		if deviceName == "out" {
			return 1
		}

		if val, ok := paths[deviceName]; ok {
			return val
		}

		sum := 0
		for _, o := range devices[deviceName] {
			sum += numberOfPaths(o)
		}
		paths[deviceName] = sum
		return sum
	}

	part1()
	part2()
}

func part1() {
	total := numberOfPaths("you")
	fmt.Println(total)
}

func part2() {
	total := numberOfPaths("svr")
	fmt.Println(total)
}
