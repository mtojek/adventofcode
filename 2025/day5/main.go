package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func part2() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var ingredientRanges []ingredientRange

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rng := scanner.Bytes()
		strRange := string(rng)

		if strRange == "" {
			break
		}

		var from, to int64
		fmt.Sscanf(strRange, "%d-%d", &from, &to)

		ingredientRanges = append(ingredientRanges, ingredientRange{
			from: from,
			to:   to,
		})
	}

	// sort by from
	sort.Slice(ingredientRanges, func(i, j int) bool {
		return ingredientRanges[i].from < ingredientRanges[j].from
	})

	// count fresh ingredients
	last := ingredientRanges[0]
	fresh := last.to - last.from + 1
	for i := range ingredientRanges {
		if i == 0 {
			continue // skip the first one since it's already processed
		}

		curr := ingredientRanges[i]
		if last.to > curr.to { // last range is larger than current
			continue // do nothing
		} else if last.to >= curr.from { // ranges overlap partially
			curr.from = last.to + 1
		}

		fresh += curr.to - curr.from + 1
		last = curr
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
