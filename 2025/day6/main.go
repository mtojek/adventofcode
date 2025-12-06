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
		var score int
		if op == '*' {
			score = 1 // fix multiplication
		}

		for j := 0; j < len(numbers); j++ {
			if op == '+' {
				score += numbers[j][i]
			} else {
				score *= numbers[j][i]
			}
		}

		total += score
	}

	fmt.Println(total)
}

func part2() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var numbers [][]byte
	var operations []byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()

		if line[0] == '+' || line[0] == '*' {
			operations = bytes.ReplaceAll(line, []byte(" "), []byte(""))
			break
		}

		t := make([]byte, len(line))
		copy(t, line)
		numbers = append(numbers, t)
	}
	rotated := rotateNumbers(numbers)

	var total int
	var j int
	for i := len(operations) - 1; i >= 0; i-- {
		op := operations[i]

		var score int
		if op == '*' {
			score = 1
		}

		for j < len(rotated) {
			if len(bytes.TrimSpace(rotated[j])) == 0 {
				j++
				break
			}

			n, err := strconv.Atoi(string(bytes.TrimSpace(rotated[j])))
			if err != nil {
				log.Fatal(err)
			}

			if op == '+' {
				score += n
			} else {
				score *= n
			}
			j++
		}

		//fmt.Println(string(op), score)
		total += score
	}

	fmt.Println(total)
}

// rotate matrix 90-degree counter clockwise
func rotateNumbers(numbers [][]byte) [][]byte {
	firstLine := numbers[0]

	var rotated [][]byte
	for i := len(firstLine) - 1; i >= 0; i-- {
		var r []byte
		for j := 0; j < len(numbers); j++ {
			r = append(r, numbers[j][i])
		}
		rotated = append(rotated, r)
	}
	return rotated
}
