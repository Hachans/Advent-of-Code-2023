package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInput("../input.txt")
	minFirstSol := firstSolution(lines)
	minSecondSol := secondSolution(lines)
	log.Println("min: ", minFirstSol, minSecondSol)

}

func firstSolution(lines []string) int {
	seeds := extractSeeds(lines[0])
	min := math.MaxInt32
	for _, seed := range seeds {
		val := seed
		for i := 3; i < len(lines); i++ {
			values := getIntValues(lines[i])
			if len(values) == 3 && val >= values[1] && val <= values[1]+values[2] {
				val = val - values[1] + values[0]
				for {
					if len(lines[i]) != 0 && i+1 < len(lines) {
						i += 1
					} else {
						break
					}
				}
			}
		}
		if val < min {
			min = val
		}
	}
	return min
}

func secondSolution(lines []string) int {
	seeds := extractSeeds(lines[0])
	min := math.MaxInt32
	log.Println(len(seeds))
	for sed := 0; sed < len(seeds); sed += 2 {
		for x := seeds[sed]; x <= seeds[sed]+seeds[sed+1]-1; x++ {
			val := x
			log.Println("SEED: ", x)
			for i := 3; i < len(lines); i++ {
				values := getIntValues(lines[i])
				if len(values) == 3 && val >= values[1] && val <= values[1]+values[2] {
					val = val - values[1] + values[0]
					for {
						if len(lines[i]) != 0 && i+1 < len(lines) {
							i += 1
						} else {
							break
						}
					}
				}
			}
			if val < min {
				min = val
			}
		}
	}
	return min
}

func readInput(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func extractSeeds(line string) []int {
	l := strings.Split(line, ": ")[1]
	k := strings.Split(l, " ")
	var seeds []int
	for _, x := range k {
		val, _ := strconv.Atoi(x)
		seeds = append(seeds, val)
	}
	return seeds
}

func getIntValues(line string) []int {
	l := strings.Split(line, " ")
	var intVals []int
	for _, s := range l {
		v, _ := strconv.Atoi(s)
		intVals = append(intVals, v)
	}
	return intVals
}
