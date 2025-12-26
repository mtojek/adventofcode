package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "input.txt"

	presentWidth  = 3
	presentHeight = 3
)

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	validRegions := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "x") {
			continue // ignore until boards
		}

		s := strings.SplitN(line, ": ", 2)

		var regionWidth int
		var regionHeight int
		fmt.Sscanf(s[0], "%dx%d", &regionWidth, &regionHeight)

		boxes := make([]int, 6)
		fmt.Sscanf(s[1], "%d %d %d %d %d %d", &boxes[0], &boxes[1], &boxes[2], &boxes[3], &boxes[4], &boxes[5])

		// part1
		// simple case: check tiles
		var numBoxes int
		for _, b := range boxes {
			numBoxes += b
		}
		totalVolume := numBoxes * presentWidth * presentHeight
		regionArea := regionWidth * regionHeight
		if totalVolume <= regionArea {
			validRegions++
		}
	}

	fmt.Println(validRegions)
}
