package main

import (
	"bufio"
	"log"
	"os"
)

const InputFile string = "./input.txt"

func readInput(file *os.File) [][]rune {
	area := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		area = append(area, []rune(line))
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return area
}

func printMap(area [][]rune) {
	for _, row := range area {
		println(string(row))
	}
}

func getStart(area [][]rune) []int {
	for i, x := range area {
		for j, y := range x {
			if y == '^' {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func getSafely(area [][]rune, x, y int) (rune, bool) {
	if x < 0 || y < 0 || x >= len(area) || y >= len(area[0]) {
		return ' ', false
	}
	return area[x][y], true
}

func move(area [][]rune, start []int) ([][]int, bool) {
	positions := make([][]int, 0)
	x := start[0]
	y := start[1]
	dir := 'N'
	iterationsSinceNewField := 0

	for {
		iterationsSinceNewField++
		// stupid loop detector but well :')
		if iterationsSinceNewField > 1000000 {
			return positions, true
		}

		value, ok := getSafely(area, x, y)
		if !ok {
			return positions, false
		}
		if value != 'X' {
			positions = append(positions, []int{x, y})
			iterationsSinceNewField = 0
		}

		area[x][y] = 'X'

		if dir == 'N' {
			value, ok := getSafely(area, x-1, y)
			if !ok {
				return positions, false
			}
			if value == '#' {
				dir = 'E'
				continue
			}
			x--
		} else if dir == 'E' {
			value, ok := getSafely(area, x, y+1)
			if !ok {
				return positions, false
			}
			if value == '#' {
				dir = 'S'
				continue
			}
			y++
		} else if dir == 'S' {
			value, ok := getSafely(area, x+1, y)
			if !ok {
				return positions, false
			}
			if value == '#' {
				dir = 'W'
				continue
			}
			x++
		} else if dir == 'W' {
			value, ok := getSafely(area, x, y-1)
			if !ok {
				return positions, false
			}
			if value == '#' {
				dir = 'N'
				continue
			}
			y--
		}
	}
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	area := readInput(file)
	printMap(area)

	start := getStart(area)
	log.Println(start)

	positions, loop := move(area, start)
	printMap(area)
	log.Println(len(positions))
	log.Println(loop)

	loops := 0
	// skip start
	for _, pos := range positions[1:] {
		file.Seek(0, 0)
		area := readInput(file)
		area[pos[0]][pos[1]] = '#'
		if _, loop := move(area, start); loop {
			loops++
		}
	}
	log.Println(loops)
}
