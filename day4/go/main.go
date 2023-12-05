package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInput("../test.txt")
	total := 0
	total2 := 0
	for _, line := range lines {
		sum := getSum(line, 0)
		total += sum
	}
	for i, line := range lines {
		i += 1
		matches := getSum(line, 1)
		total2 += matches
		log.Println("CARD: ", i, "MATCHES: ", matches)
		for k := i + 1; k <= i+matches; k++ {
			submatches := getSum(lines[k], 1)
			log.Println("subCard: ", k, "subMatches : ", submatches)
			total2 += submatches
			for x := 0; x < submatches; x++ {
				sum := getSum(lines[k], 1)
				log.Println("subMatch: ", k+1)
				total2 += sum
			}
			log.Println("card: ", k)
		}

		log.Println("---")
	}
	log.Println(total, total2)
}

func getSum(line string, flag int) int {
	values := parseLine(line)
	winning := extractNums(values[0])
	mine := extractNums(values[1])
	sum := 0
	if flag == 1 {
		sum = compare(winning, mine)

	} else {
		sum = compareAndMultiply(winning, mine)

	}
	return sum
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

func parseLine(line string) []string {
	splits := strings.Split(line, ": ")[1]
	nums := strings.Split(splits, " | ")
	return nums
}

func extractNums(values string) []int {
	var nums []int
	numList := strings.Split(values, " ")
	for _, num := range numList {
		if len(num) == 0 {
			continue
		}
		nInt, _ := strconv.Atoi(string(num))
		nums = append(nums, nInt)
	}
	return nums
}

func compareAndMultiply(winning []int, mine []int) int {
	count := 0
	for _, w := range winning {
		for _, m := range mine {
			if m == w {
				count += 1
			}
		}
	}
	if count == 1 || count == 0 {
		return count
	}
	if count == 2 {
		return 1 * count
	}
	sum := 1
	for i := 1; i < count; i++ {
		sum *= 2
	}
	return sum
}

func compare(winning []int, mine []int) int {
	count := 0
	for _, w := range winning {
		for _, m := range mine {
			if m == w {
				count += 1
			}
		}
	}
	return count
}
