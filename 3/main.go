package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const InputFile string = "./input.txt"

func readInput(file *os.File) []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return lines
}

var pattern *regexp.Regexp = regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)

func matchToMulResult(match []string) int {
	a, err := strconv.Atoi(match[2])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(match[3])
	if err != nil {
		log.Fatal(err)
	}

	return a * b
}

func sumOfMuls(lines []string) int {
	sum := 0

	do := true
	for _, line := range lines {
		matches := pattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if strings.HasPrefix(match[1], "do(") {
				do = true
			} else if strings.HasPrefix(match[1], "don't") {
				do = false
			} else if do && strings.HasPrefix(match[1], "mul") {
				sum += matchToMulResult(match)
			}
		}
	}

	return sum
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := readInput(file)

	sum := sumOfMuls(lines)

	log.Println(sum)
}
