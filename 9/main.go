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

func unwrap(pattern []int) []int {
	layout := make([]int, 0)

	id := 0
	for i, n := range pattern {
		if i%2 == 0 {
			layout = append(layout, slices.Repeat([]int{id}, n)...)
			id++
		} else {
			layout = append(layout, slices.Repeat([]int{Empty}, n)...)
		}
	}
	return layout
}

func compact(layout []int) {
	i := 0
	j := len(layout) - 1
	for i < j {
		for layout[i] != Empty {
			i++
		}

		for layout[j] == Empty {
			j--
		}

		layout[i] = layout[j]
		layout[j] = Empty
		i++
		j--
	}
}

func checksum(layout []int) int {
	sum := 0
	i := 0
	for _, n := range layout {
		if n == Empty {
			continue
		}
		sum += i * n
		i++
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
