package main

import (
	"fmt"
	"log"
	"os"
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
	f, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	ranges := strings.Split(string(f), ",")

	sum := 0
	for _, r := range ranges {
		from := 0
		to := 0

		_, err := fmt.Sscanf(r, "%d-%d", &from, &to)
		if err != nil {
			log.Fatal(err)
		}

		for i := from; i <= to; i++ {
			str := fmt.Sprintf("%d", i)
			if len(str)%2 == 1 {
				continue
			}

			if str[0:len(str)/2] != str[len(str)/2:] {
				continue
			}
			sum += i
		}
	}

	fmt.Println(sum)
}

func part2() {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	ranges := strings.Split(string(f), ",")

	sum := 0
	for _, r := range ranges {
		from := 0
		to := 0

		_, err := fmt.Sscanf(r, "%d-%d", &from, &to)
		if err != nil {
			log.Fatal(err)
		}

		for i := from; i <= to; i++ {
			str := fmt.Sprintf("%d", i)

			for j := 1; j <= len(str)/2; j++ {
				pattern := str[:j]

				m := strings.Count(str, pattern)
				if m*len(pattern) == len(str) {
					sum += i
					break
				}
			}
		}
	}

	fmt.Println(sum)
}
