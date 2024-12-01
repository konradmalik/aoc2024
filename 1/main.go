package main

import (
	"bufio"
	"log"
	"os"
	"slices"

	"github.com/konradmalik/aoc2024/lib"
)

const InputFile string = "./input.txt"

func readInput(file *os.File) [][]int {
	rows := make([][]int, 2)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := lib.ParseNumbers(scanner.Text())
		rows[0] = append(rows[0], row[0])
		rows[1] = append(rows[1], row[1])
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return rows
}

func distances(a, b []int) []int {
	ds := make([]int, len(a))
	for i, xa := range a {
		xb := b[i]

		d := xa - xb
		if d < 0 {
			d = -d
		}
		ds[i] = d
	}
	return ds
}

func sum(a []int) int {
	sum := 0
	for _, x := range a {
		sum += x
	}
	return sum
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	columns := readInput(file)
	slices.Sort(columns[0])
	slices.Sort(columns[1])

	distances := distances(columns[0], columns[1])

	sum := sum(distances)

	log.Println(sum)
}
