package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "input.txt"

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

	part1(devices)
	part2(devices)
}

func part1(devices map[string][]string) {
	paths := map[string]int{}
	var numberOfPaths func(current, destination string) int
	numberOfPaths = func(current, destination string) int {
		if current == destination {
			return 1
		}

		if val, ok := paths[current]; ok {
			return val
		}

		sum := 0
		for _, o := range devices[current] {
			sum += numberOfPaths(o, destination)
		}
		paths[current] = sum
		return sum
	}

	total := numberOfPaths("you", "out")
	fmt.Println(total)
}

func part2(devices map[string][]string) {
	type key struct {
		deviceName string
		fft        bool
		dac        bool
	}
	paths := map[key]int{}

	var numberOfPaths func(current key, destination string) int
	numberOfPaths = func(current key, destination string) int {
		if current.deviceName == "fft" {
			current.fft = true
		}
		if current.deviceName == "dac" {
			current.dac = true
		}

		if current.deviceName == "out" {
			if current.fft && current.dac {
				return 1
			}
			return 0
		}

		if val, ok := paths[current]; ok {
			return val
		}

		sum := 0
		for _, next := range devices[current.deviceName] {
			sum += numberOfPaths(key{
				deviceName: next,
				fft:        current.fft,
				dac:        current.dac,
			}, destination)
		}
		paths[current] = sum
		return sum
	}

	total := numberOfPaths(key{
		deviceName: "svr",
	}, "out")
	fmt.Println(total)
}
