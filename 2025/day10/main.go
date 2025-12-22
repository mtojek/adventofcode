package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"maps"
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
		fmt.Println(i)

		m := toggleLights(m.lights, m.buttons, initLights(m.lights), map[int]bool{})
		sum += m
	}
	fmt.Println(sum)
}

func initLights(lights []byte) []byte {
	return bytes.Repeat([]byte{'.'}, len(lights))
}

func toggleLights(lights []byte, buttons [][]int, currentLights []byte, pressedButtons map[int]bool) int {
	min := math.MaxInt
	for i, button := range buttons {
		if _, pressed := pressedButtons[i]; pressed {
			continue
		}

		nextPressedButtons := maps.Clone(pressedButtons)
		nextPressedButtons[i] = true

		nextLights := slices.Clone(currentLights)
		for _, lightIndex := range button {
			val := nextLights[lightIndex]
			if val == '.' {
				val = '#'
			} else {
				val = '.'
			}
			nextLights[lightIndex] = val
		}

		if bytes.Equal(lights, nextLights) {
			return len(nextPressedButtons)
		}

		m := toggleLights(lights, buttons, nextLights, nextPressedButtons)
		if m < min {
			min = m
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
