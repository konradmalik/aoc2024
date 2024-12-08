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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findAntinodes(area [][]rune, positions [][]int) [][]int {
	nodes := make([][]int, 0)

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			pos := positions[i]
			ppos := positions[j]

			dx := pos[0] - ppos[0]
			dy := pos[1] - ppos[1]
			step := gcd(dx, dy)

			stepx := dx / step
			stepy := dy / step

			x := pos[0]
			y := pos[1]
			for {
				if _, ok := getSafely(area, x, y); !ok {
					break
				}
				nodes = append(nodes, []int{x, y})
				x = x - stepx
				y = y - stepy
			}

			x = pos[0]
			y = pos[1]
			for {
				if _, ok := getSafely(area, x, y); !ok {
					break
				}
				nodes = append(nodes, []int{x, y})
				x = x + stepx
				y = y + stepy
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
