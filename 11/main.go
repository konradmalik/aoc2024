package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/konradmalik/aoc2024/lib"
)

const InputFile string = "./input.txt"
const Blinks int = 25

func readInput(file *os.File) []int {
	area := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		area = append(area, lib.ParseNumbersSep(line, " ")...)
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return area
}

type stone struct {
	value      int
	blinksLeft int
}

var memo map[stone]int = make(map[stone]int)

func blinkAtStone(value, blinksLeft int) int {
	if cached, ok := memo[stone{value, blinksLeft}]; ok {
		return cached
	}

	if blinksLeft == 0 {
		return 1
	}

	if value == 0 {
		ret := blinkAtStone(1, blinksLeft-1)
		memo[stone{1, blinksLeft - 1}] = ret
		return ret
	}

	str := strconv.Itoa(value)
	if n := len(str); n%2 == 0 {
		a, err := strconv.Atoi(str[:n/2])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(str[n/2:])
		if err != nil {
			panic(err)
		}
		ar := blinkAtStone(a, blinksLeft-1)
		memo[stone{a, blinksLeft - 1}] = ar
		br := blinkAtStone(b, blinksLeft-1)
		memo[stone{b, blinksLeft - 1}] = br
		return ar + br
	}

	ret := blinkAtStone(value*2024, blinksLeft-1)
	memo[stone{value * 2024, blinksLeft - 1}] = ret
	return ret
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	stones := readInput(file)
	for _, stone := range stones {
		sum += blinkAtStone(stone, 75)
	}
	log.Println(sum)
}
