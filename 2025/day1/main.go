package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	c := 50
	zeros := 0

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		move := scanner.Text()

		steps, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Fatal(err)
		}

		op := 1
		if move[0] == 'L' {
			op = -1
		}

		for i := 0; i < steps; i++ {
			c = c + op

			if c < 0 {
				c += 100
			}
			c = c % 100

			if c == 0 {
				zeros++
			}
		}

		fmt.Printf("%s c = %d\n", move, c)
	}

	fmt.Println("c = ", c)
	fmt.Println("zeros = ", zeros)
}
