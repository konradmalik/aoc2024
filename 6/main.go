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

func move(area [][]rune, start []int) int {
	count := 0
	x := start[0]
	y := start[1]
	dir := 'N'

	for {
		value, ok := getSafely(area, x, y)
		if !ok {
			return count
		}
		if value != 'X' {
			count++
		}

		area[x][y] = 'X'

		if dir == 'N' {
			value, ok := getSafely(area, x-1, y)
			if !ok {
				return count
			}
			if value == '#' {
				dir = 'E'
				y++
			} else {
				x--
			}
		} else if dir == 'E' {
			value, ok := getSafely(area, x, y+1)
			if !ok {
				return count
			}
			if value == '#' {
				dir = 'S'
				x++
			} else {
				y++
			}
		} else if dir == 'S' {
			value, ok := getSafely(area, x+1, y)
			if !ok {
				return count
			}
			if value == '#' {
				dir = 'W'
				y--
			} else {
				x++
			}
		} else if dir == 'W' {
			value, ok := getSafely(area, x, y-1)
			if !ok {
				return count
			}
			if value == '#' {
				dir = 'N'
				x--
			} else {
				y--
			}
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

	count := move(area, start)
	printMap(area)
	log.Println(count)
}
