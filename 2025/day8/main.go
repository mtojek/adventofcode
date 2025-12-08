package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"

	maxConnections = 1000
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

	// Load junction boxes
	var boxes []junctionBox

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}

		z, err := strconv.Atoi(coords[2])
		if err != nil {
			log.Fatal(err)
		}

		boxes = append(boxes, junctionBox{
			x: x,
			y: y,
			z: z,
		})
	}

	// Calculate distances
	var shortest []junctionBoxDistance
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes); j++ {
			if i == j {
				continue // can't connect to itself
			}

			a := boxes[i]
			b := boxes[j]

			order := []junctionBox{a, b}
			slices.SortFunc(order, func(a, b junctionBox) int {
				if a.x < b.x {
					return -1
				}
				if a.x > b.x {
					return 1
				}

				if a.y < b.y {
					return -1
				}
				if a.y > b.y {
					return 1
				}

				if a.z < b.z {
					return -1
				}
				if a.z > b.z {
					return 1
				}
				return 0
			})

			distance := math.Sqrt(math.Pow(math.Abs(float64(b.x-a.x)), 2) +
				math.Pow(math.Abs(float64(b.y-a.y)), 2) +
				math.Pow(math.Abs(float64(b.z-a.z)), 2))

			jbd := junctionBoxDistance{
				from:     order[0],
				to:       order[1],
				distance: distance,
			}

			//if !slices.Contains(shortest, jbd) {
			shortest = append(shortest, jbd)
			//}
		}
	}

	// Order distances ASC
	slices.SortFunc(shortest, func(d1, d2 junctionBoxDistance) int {
		if d1.distance < d2.distance {
			return -1
		}
		if d1.distance > d2.distance {
			return 1
		}
		return 0
	})

	// Join circuits
	circuits := [][]string{}
	for _, box := range boxes {
		circuits = append(circuits, []string{box.String()})
	}

	moves := 0
	for i := 0; i < len(shortest); i += 2 {
		d := shortest[i]

		for i, c := range circuits {
			if slices.Contains(c, d.from.String()) && slices.Contains(c, d.to.String()) {
				break
			} else if slices.Contains(c, d.from.String()) && !slices.Contains(c, d.to.String()) {
				t := circuits[i]
				circuits[i] = append(t, d.to.String())
			} else if !slices.Contains(c, d.from.String()) && slices.Contains(c, d.to.String()) {
				t := circuits[i]
				circuits[i] = append(t, d.from.String())
			}
		}

		circuits = merge(circuits)

		moves++
		if moves == maxConnections {
			break
		}
	}

	// Find top 3
	var sizes []int
	for _, c := range circuits {
		sizes = append(sizes, len(c))
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})

	fmt.Println(sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3])
}

func merge(circuits [][]string) [][]string {
	merge := true
	for merge {
		merge = false

		for i, c1 := range circuits {
			for j, c2 := range circuits {
				if i == j {
					continue
				}

				if !intersect(c1, c2) {
					continue
				}

				c3 := unique(append(c1, c2...))
				circuits[i] = c3
				circuits[j] = []string{}
				merge = true
			}
		}

		var fixedCircuits [][]string
		for _, c := range circuits {
			if len(c) == 0 {
				continue
			}
			fixedCircuits = append(fixedCircuits, c)
		}
		circuits = fixedCircuits
	}
	return circuits
}

func intersect(a, b []string) bool {
	for _, elem := range a {
		if slices.Contains(b, elem) {
			return true
		}
	}
	return false
}

func unique(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	result := make([]string, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

type junctionBox struct {
	x int
	y int
	z int
}

func (jb junctionBox) String() string {
	return fmt.Sprintf("%d,%d,%d", jb.x, jb.y, jb.z)
}

type junctionBoxDistance struct {
	from junctionBox
	to   junctionBox

	distance float64
}

func part2() {
}
