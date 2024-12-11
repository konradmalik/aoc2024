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

func blinkAtStone(stone, blinksLeft int) int {
	if blinksLeft == 0 {
		return 1
	}

	if stone == 0 {
		return blinkAtStone(1, blinksLeft-1)
	}

	str := strconv.Itoa(stone)
	if n := len(str); n%2 == 0 {
		a, err := strconv.Atoi(str[:n/2])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(str[n/2:])
		if err != nil {
			panic(err)
		}
		return blinkAtStone(a, blinksLeft-1) + blinkAtStone(b, blinksLeft-1)
	}

	return blinkAtStone(stone*2024, blinksLeft-1)

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
		sum += blinkAtStone(stone, 25)
	}
	log.Println(sum)
}
