package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/konradmalik/aoc2024/lib"
)

const InputFile string = "./input.txt"

func readInput(file *os.File) [][]int {
	rows := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ":", "", 1)

		rows = append(rows, lib.ParseNumbers(line))
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return rows
}

func concat(a, b int) int {
	v, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		panic(err)
	}
	return v
}

func matchOps(expected int, nums []int, acc int) bool {
	if len(nums) == 0 {
		return expected == acc
	}

	return matchOps(expected, nums[1:], acc+nums[0]) || matchOps(expected, nums[1:], acc*nums[0]) || matchOps(expected, nums[1:], concat(acc, nums[0]))
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	rows := readInput(file)
	for _, row := range rows {
		if matchOps(row[0], row[2:], row[1]) {
			sum += row[0]
		}
	}
	log.Println(sum)
}
