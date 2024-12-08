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

func collectAntennas(area [][]rune) map[rune][][]int {
	ret := make(map[rune][][]int)

	for i, row := range area {
		for j, r := range row {
			if r != '.' {
				positions, ok := ret[r]
				if !ok {
					positions = make([][]int, 0)
				}
				ret[r] = append(positions, []int{i, j})
			}
		}
	}

	return ret
}

func getSafely(area [][]rune, x, y int) (rune, bool) {
	if x < 0 || y < 0 || x >= len(area) || y >= len(area[0]) {
		return ' ', false
	}
	return area[x][y], true
}

func findAntinodes(area [][]rune, positions [][]int) [][]int {
	nodes := make([][]int, 0)

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			pos := positions[i]
			ppos := positions[j]
			xm := float64(ppos[0]+pos[0]) / 2
			ym := float64(ppos[1]+pos[1]) / 2

			dx := xm - float64(pos[0])
			dy := ym - float64(pos[1])

			x1 := int(xm + 3*dx)
			y1 := int(ym + 3*dy)
			x2 := int(xm - 3*dx)
			y2 := int(ym - 3*dy)

			if _, ok := getSafely(area, x1, y1); ok {
				nodes = append(nodes, []int{x1, y1})
			}
			if _, ok := getSafely(area, x2, y2); ok {
				nodes = append(nodes, []int{x2, y2})
			}
		}
	}

	return nodes
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	area := readInput(file)
	printMap(area)

	antennas := collectAntennas(area)

	antinodes := make([][]int, 0)
	for _, positions := range antennas {
		antinodes = append(antinodes, findAntinodes(area, positions)...)
	}

	unique := 0
	for _, node := range antinodes {
		if area[node[0]][node[1]] != '#' {
			area[node[0]][node[1]] = '#'
			unique++
		}
	}
	printMap(area)
	log.Println(unique)
}
