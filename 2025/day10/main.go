package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input0.txt"
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

	// Load machine configuration
	var machines []machine

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		i := strings.Index(line, "]")
		lights := line[1:i]

		j := strings.Index(line, "{")
		strButtonGroups := line[i+2:]
		buttonGroups := strings.Split(strButtonGroups, " ")

		var buttons [][]int
		for _, buttonGroup := range buttonGroups {
			bg := buttonGroup[1 : len(buttonGroup)-1]
			sbg := strings.Split(bg, ",")

			var btns []int
			for _, b := range sbg {
				button, err := strconv.Atoi(b)
				if err != nil {
					log.Fatal(err)
				}
				btns = append(btns, button)
			}

			buttons = append(buttons, btns)
		}

		strJoltage := line[j+1 : len(line)-1]
		sJlt := strings.Split(strJoltage, ",")

		var joltage []int
		for _, j := range sJlt {
			jolt, err := strconv.Atoi(j)
			if err != nil {
				log.Fatal(err)
			}
			joltage = append(joltage, jolt)
		}

		machines = append(machines, machine{
			lights:  []byte(lights),
			buttons: buttons,
			joltage: joltage,
		})
	}

	// Display machines
	for _, m := range machines {
		fmt.Print(string(m.lights))
		fmt.Print(" ")
		fmt.Print(m.buttons)
		fmt.Print(" ")
		fmt.Println(m.joltage)
	}
	fmt.Println(machines)
}

func part2() {
}

type machine struct {
	lights  []byte
	buttons [][]int
	joltage []int
}
