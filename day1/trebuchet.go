package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// https://adventofcode.com/2023/day/1#part2

func main() {
	lines := readInput("input.txt")
	myDict := initDict()
	total := 0

	for _, line := range lines {
		first := extractFirst(line, myDict)
		last := extractLast(line, myDict)
		sum, _ := strconv.Atoi(first + last)
		total += sum
	}

	log.Println(total)
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

func initDict() map[string]string {
	myDict := make(map[string]string)

	myDict["one"] = "1"
	myDict["two"] = "2"
	myDict["three"] = "3"
	myDict["four"] = "4"
	myDict["five"] = "5"
	myDict["six"] = "6"
	myDict["seven"] = "7"
	myDict["eight"] = "8"
	myDict["nine"] = "9"

	return myDict
}

func extractFirst(line string, myDict map[string]string) string {
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			return string(line[i])
		} else {
			for k, v := range myDict {
				if strings.Contains(line[:i+1], k) {
					return v
				}
			}
		}
	}
	return ""
}

func extractLast(line string, myDict map[string]string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return string(line[i])
		} else {
			for k, v := range myDict {
				if strings.Contains(line[i:], k) {
					return v
				}
			}
		}
	}
	return ""
}
