package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
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

var pattern *regexp.Regexp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func sumOfMuls(line string) int {
	sum := 0

	matches := pattern.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}

		sum += a * b
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

	sum := 0
	for _, line := range lines {
		sum += sumOfMuls(line)
	}

	log.Println(sum)
}
