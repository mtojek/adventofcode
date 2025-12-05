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

	var ingredientRanges []ingredientRange
	var fresh int

	var ingredientsMode bool
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rng := scanner.Bytes()
		strRange := string(rng)

		if strRange == "" {
			ingredientsMode = true
			continue
		}

		if ingredientsMode {
			var id int64
			fmt.Sscanf(strRange, "%d", &id)

			if isFresh(ingredientRanges, id) {
				fresh++
			}

		} else {
			var from, to int64
			fmt.Sscanf(strRange, "%d-%d", &from, &to)

			ingredientRanges = append(ingredientRanges, ingredientRange{
				from: from,
				to:   to,
			})
		}

	}

	fmt.Println(fresh)
}

type ingredientRange struct {
	from int64
	to   int64
}

func isFresh(ranges []ingredientRange, id int64) bool {
	for _, r := range ranges {
		if r.from <= id && id <= r.to {
			return true
		}
	}
	return false
}

func part2() {

}
