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

func makeOccurrences(a []int) map[int]int {
	occurrences := make(map[int]int)
	for _, x := range a {
		val, ok := occurrences[x]
		if !ok {
			occurrences[x] = 1
		} else {
			occurrences[x] = val + 1
		}
	}
	return occurrences
}

func similarity(a, b []int) int {
	occurrences := makeOccurrences(b)
	s := 0
	for _, xa := range a {
		occ, ok := occurrences[xa]
		if !ok {
			occ = 0
		}

		s = s + xa*occ
	}

	return s
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

	sim := similarity(columns[0], columns[1])

	log.Println(sim)
}
