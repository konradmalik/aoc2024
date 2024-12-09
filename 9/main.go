package main

import (
	"bufio"
	"log"
	"os"
	"slices"

	"github.com/konradmalik/aoc2024/lib"
)

const InputFile string = "./input.txt"
const Empty int = -1

type blob struct {
	Id   int
	Size int
}

func readInput(file *os.File) []int {
	rows := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, lib.ParseNumbersSep(line, "")...)
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return rows
}

func unwrap(pattern []int) []blob {
	layout := make([]blob, 0)

	id := 0
	for i, n := range pattern {
		if i%2 == 0 {
			layout = append(layout, slices.Repeat([]blob{{id, n}}, n)...)
			id++
		} else {
			layout = append(layout, slices.Repeat([]blob{{Empty, n}}, n)...)
		}
	}
	return layout
}

func compact(layout []blob) {
	i := 0
	j := len(layout) - 1
	for j > 0 {
		if i >= j {
			i = 0
			j--
			continue
		}

		space := layout[i]
		blob := layout[j]

		if blob.Id == Empty {
			j--
			continue
		}

		if space.Id != Empty || blob.Size > space.Size {
			i++
			continue
		}

		// shrink empty space
		for x := 0; x < space.Size; x++ {
			layout[i+x].Size = space.Size - blob.Size
		}

		// swap
		for x := 0; x < blob.Size; x++ {
			tmp := layout[i]
			layout[i] = layout[j]
			layout[j] = tmp
			i++
			j--
		}
		// reset left hand for the next blob
		i = 0
	}
}

func checksum(layout []blob) int {
	sum := 0
	for i, b := range layout {
		if b.Id == Empty {
			continue
		}
		sum += i * b.Id
	}
	return sum
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows := readInput(file)
	layout := unwrap(rows)
	log.Println(layout)
	compact(layout)
	log.Println(layout)
	check := checksum(layout)
	log.Println(check)
}
