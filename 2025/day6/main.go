package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

var (
	reWhitespaces = regexp.MustCompile(`\s+`)
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

	var numbers [][]int
	var operations []byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.TrimSpace(reWhitespaces.ReplaceAllString(line, " "))

		if strs[0] == '+' || strs[0] == '*' {
			operations = bytes.ReplaceAll([]byte(strs), []byte(" "), []byte(""))
			break
		}

		var num []int
		for _, str := range strings.Split(strs, " ") {
			n, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}

			num = append(num, n)
		}
		numbers = append(numbers, num)
	}

	var total int

	for i, op := range operations {
		var sum int
		if op == '*' {
			sum = 1 // fix multiplication
		}

		for j := 0; j < len(numbers); j++ {
			if op == '+' {
				sum += numbers[j][i]
			} else {
				sum *= numbers[j][i]
			}
		}

		total += sum
	}

	fmt.Println(total)
}

func part2() {
}
