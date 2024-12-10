package main

import (
	"bufio"
	"log"
	"os"

	"github.com/konradmalik/aoc2024/lib"
)

const InputFile string = "./input.txt"

func readInput(file *os.File) [][]int {
	area := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		area = append(area, lib.ParseNumbersSep(line, ""))
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return area
}

func printMap(area [][]int) {
	for _, row := range area {
		for _, point := range row {
			print(point)
		}
		println()
	}
}

func getSafely(area [][]int, x, y int) (int, bool) {
	if x < 0 || y < 0 || x >= len(area) || y >= len(area[0]) {
		return ' ', false
	}
	return area[x][y], true
}

type point struct {
	X int
	Y int
}

func tailheadScore(area [][]int, x, y, prev int, found map[point]bool) int {
	curr, ok := getSafely(area, x, y)
	if !ok {
		return 0
	}

	if curr-prev != 1 {
		return 0
	}

	if curr == 9 {
		if _, ok := found[point{x, y}]; ok {
			return 0
		}
		found[point{x, y}] = true
		return 1
	}

	return tailheadScore(area, x-1, y, curr, found) +
		tailheadScore(area, x+1, y, curr, found) +
		tailheadScore(area, x, y-1, curr, found) +
		tailheadScore(area, x, y+1, curr, found)
}

func scoreTrails(area [][]int) int {
	sum := 0
	for x, row := range area {
		for y, p := range row {
			if p == 0 {
				sum += tailheadScore(area, x, y, -1, make(map[point]bool))
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

	topo := readInput(file)
	printMap(topo)
	sum := scoreTrails(topo)
	log.Println(sum)
}
