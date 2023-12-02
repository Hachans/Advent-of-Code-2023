package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

func main() {
	lines := readInput("../input.txt")
	sum := 0
	for _, line := range lines {
		if len(os.Args) <= 1 {
			// NOTE: First part solution
			id, values := parseLine(line)
			valid := checkIfValid(values)
			if valid {
				idInt, _ := strconv.Atoi(id)
				sum += idInt
			}
		} else {
			// NOTE: Second part solution
			_, values := parseLine(line)
			multiplied := findMinAmount(values)
			sum += multiplied
		}
	}
	log.Println("Total sum: ", sum)
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

func parseLine(line string) (string, []string) {
	//extract id
	idString := strings.Split(line, ":")[0]
	id := strings.Split(idString, " ")[1]

	//extract cubes
	var cubes []string
	cubeString := strings.Split(line, ":")[1]
	rounds := strings.Split(cubeString, ";")
	for _, entry := range rounds {
		parts := strings.Fields(entry)
		for i, part := range parts {
			parts[i] = strings.ReplaceAll(part, ",", "")
		}
		cubes = append(cubes, parts...)
	}
	return id, cubes
}

func checkIfValid(cubes []string) bool {
	flag := true
	for i := 0; i < len(cubes); i += 2 {
		switch color := cubes[i+1]; color {
		case RED:
			count, _ := strconv.Atoi(cubes[i])
			if count > 12 {
				flag = false
			}
		case GREEN:
			count, _ := strconv.Atoi(cubes[i])
			if count > 13 {
				flag = false
			}
		case BLUE:
			count, _ := strconv.Atoi(cubes[i])
			if count > 14 {
				flag = false
			}
		}
	}
	return flag
}

func findMinAmount(cubes []string) int {
	red := 0
	green := 0
	blue := 0
	for i := 0; i < len(cubes); i += 2 {
		switch color := cubes[i+1]; color {
		case RED:
			count, _ := strconv.Atoi(cubes[i])
			if count > red {
				red = count
			}
		case GREEN:
			count, _ := strconv.Atoi(cubes[i])
			if count > green {
				green = count
			}
		case BLUE:
			count, _ := strconv.Atoi(cubes[i])
			if count > blue {
				blue = count
			}
		}
	}
	sum := red * green * blue
	return sum
}
