package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
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
		j := strings.Index(line, "{")

		lights := line[1:i]
		strButtonGroups := line[i+2 : j-1]
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

	// Toggle all machines
	var sum int
	for i, m := range machines {
		fmt.Println("iter:", i)

		startState := bytes.Repeat([]byte{'.'}, len(m.lights))
		m := toggleLights(m.lights, m.buttons, startState, 0)

		sum += m
	}
	fmt.Println(sum)
}

func toggleLights(lights []byte, buttons [][]int, currentLights []byte, nextButton int) int {
	if bytes.Equal(lights, currentLights) {
		return 0
	}

	min := math.MaxInt
	for i := nextButton; i < len(buttons); i++ {
		nextLights := slices.Clone(currentLights)
		for _, lightIndex := range buttons[i] {
			if nextLights[lightIndex] == '.' {
				nextLights[lightIndex] = '#'
			} else {
				nextLights[lightIndex] = '.'
			}
		}

		m := toggleLights(lights, buttons, nextLights, i+1)
		if m != math.MaxInt {
			m++
			if m < min {
				min = m
			}
		}
	}
	return min
}

type machine struct {
	lights  []byte
	buttons [][]int
	joltage []int
}

func part2() {
}
