package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
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

func getSafely(area [][]rune, x, y int) (rune, bool) {
	if x < 0 || y < 0 || x >= len(area) || y >= len(area[0]) {
		return ' ', false
	}
	return area[x][y], true
}

type node struct {
	X int
	Y int
}

func floodFillWithPerimeter(area [][]rune, s node) ([]node, int) {
	color := area[s.X][s.Y]
	queue := []node{s}
	// if visited key present, if same color bool true
	visited := make(map[node]bool)
	perimeter := 0

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if visited[n] {
			continue
		}

		value, ok := getSafely(area, n.X, n.Y)

		if !ok || value != color {
			perimeter++
		}

		if !ok {
			continue
		}

		visited[n] = false

		if value == color {
			visited[n] = true
			// NOTE also setting area to lowercase
			area[n.X][n.Y] = unicode.ToLower(area[n.X][n.Y])
			queue = append(queue, []node{{n.X + 1, n.Y}, {n.X - 1, n.Y}, {n.X, n.Y + 1}, {n.X, n.Y - 1}}...)
		}
	}

	ret := make([]node, 0)
	for k, v := range visited {
		if v {
			ret = append(ret, k)
		}
	}

	return ret, perimeter
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	area := readInput(file)
	printMap(area)

	sum := 0
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[i]); j++ {
			if !unicode.IsLower(area[i][j]) {
				region, perimeter := floodFillWithPerimeter(area, node{i, j})
				log.Println(string(area[i][j]))
				log.Println(len(region))
				log.Println(perimeter)
				sum += len(region) * perimeter
			}
		}
	}
	log.Println(sum)
}
