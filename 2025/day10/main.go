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

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
)

const (
	inputFile = "input.txt"
)

type machine struct {
	lights  []byte
	buttons [][]int
	joltage []int
}

func main() {
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

	part1(machines)
	part2(machines)
}

func part1(machines []machine) {
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

func part2(machines []machine) {
	var sum int
	for _, m := range machines {

		// minimize number of pressed buttons
		c := []float64{1, 1, 1, 1, 1, 1}

		// equality: A * x = b
		A := mat.NewDense(len(m.joltage), len(m.lights), []float64{
			// FIXME
		})
		b := joltageToFloat(m.joltage)

		// solve now
		min, x, err := lp.Simplex(c, A, b, 0.0, nil)
		if err != nil {
			log.Fatal("simplex error:", err)
		}

		fmt.Println(min, x)
		sum += int(min)
	}
	fmt.Println(sum)
}

func joltageToFloat(joltage []int) []float64 {
	f := make([]float64, len(joltage))
	for i, j := range joltage {
		f[i] = float64(j)
	}
	return f
}
