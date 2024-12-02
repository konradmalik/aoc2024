package main

import (
	"bufio"
	"log"
	"os"

	"github.com/konradmalik/aoc2024/lib"
)

const InputFile string = "./input.txt"

func readInput(file *os.File) [][]int {
	rows := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := lib.ParseNumbers(scanner.Text())
		rows = append(rows, row)
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return rows
}

func removeAsCopy(slice []int, s int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:s]...)
	return append(ret, slice[s+1:]...)
}

func isRowSafe(row []int) int {

	prevTrend := 0
	for i := range row {
		if i == 0 {
			continue
		}

		x1 := row[i-1]
		x2 := row[i]

		diff := x2 - x1

		if diff > 3 || diff < -3 || diff == 0 {
			return 0
		}

		trend := 0
		if diff < 0 {
			trend = -1
		} else if diff > 0 {
			trend = 1
		}

		if prevTrend != 0 && prevTrend != trend {
			return 0
		}
		prevTrend = trend
	}

	return 1
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reports := readInput(file)

	safe := 0
	for _, row := range reports {
		if isRowSafe(row) > 0 {
			safe++
			continue
		}

		for i := range row {
			if isRowSafe(removeAsCopy(row, i)) > 0 {
				safe++
				break
			}
		}
	}

	log.Println(safe)
}
