package main

import (
	"bufio"
	"log"
	"os"
	"slices"
)

const InputFile string = "./input.txt"

func readInput(file *os.File) []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return lines
}

func stringsToRunes(s []string) [][]rune {
	ret := make([][]rune, len(s))
	for i, line := range s {
		ret[i] = []rune(line)
	}
	return ret
}

func iterateHorizontally(lines []string, reversed bool) []string {
	if !reversed {
		return lines
	}

	ret := make([]string, len(lines))
	for i, line := range lines {
		runes := []rune(line)
		slices.Reverse(runes)

		ret[i] = string(runes)
	}
	return ret
}

func iterateVertically(lines []string, reversed bool) []string {
	ret := make([]string, len(lines))

	runes := stringsToRunes(lines)

	for i := 0; i < len(runes[0]); i++ {
		row := make([]rune, len(runes))
		for j := 0; j < len(runes); j++ {
			row[j] = runes[j][i]
		}
		if reversed {
			slices.Reverse(row)
		}
		ret[i] = string(row)
	}

	return ret
}

// assumes a rectangle NxN
func iterateDiagonallyLeft(lines []string, reversed bool) []string {
	ret := make([]string, 0)

	runes := stringsToRunes(lines)

	// 0,0;
	// 0,1; 1,0
	// 0,2; 1,1; 2,0
	// 0,3; 1,2; 2,1; 3,0

	// top half
	for k := 0; k < len(runes); k++ {
		row := make([]rune, 0)
		for j := 0; j <= k; j++ {
			i := k - j
			row = append(row, runes[i][j])
		}
		if reversed {
			slices.Reverse(row)
		}
		ret = append(ret, string(row))
	}

	// bottom half
	for k := len(runes) - 2; k >= 0; k-- {
		row := make([]rune, 0)
		for j := 0; j <= k; j++ {
			i := k - j
			row = append(row, runes[len(runes)-j-1][len(runes)-i-1])
		}
		if reversed {
			slices.Reverse(row)
		}
		ret = append(ret, string(row))
	}

	return ret
}

// assumes a rectangle NxN
func iterateDiagonallyRight(lines []string, reversed bool) []string {
	ret := make([]string, 0)

	runes := stringsToRunes(lines)

	// top half
	for k := 0; k < len(runes); k++ {
		row := make([]rune, 0)
		for j := 0; j <= k; j++ {
			i := k - j
			row = append(row, runes[i][len(runes)-j-1])
		}
		if reversed {
			slices.Reverse(row)
		}
		ret = append(ret, string(row))
	}

	// bottom half
	for k := len(runes) - 2; k >= 0; k-- {
		row := make([]rune, 0)
		for j := 0; j <= k; j++ {
			i := k - j
			row = append(row, runes[len(runes)-i-1][j])
		}
		if reversed {
			slices.Reverse(row)
		}
		ret = append(ret, string(row))
	}

	return ret
}

func countXmas(words []string) int {
	count := 0
	for _, word := range words {
		for i := range word {
			iend := i + 4
			if iend > len(word) {
				break
			}
			part := word[i:iend]
			if part == "XMAS" {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := readInput(file)

	count := 0
	words := iterateHorizontally(lines, false)
	count += countXmas(words)
	words = iterateHorizontally(lines, true)
	count += countXmas(words)
	words = iterateVertically(lines, false)
	count += countXmas(words)
	words = iterateVertically(lines, true)
	count += countXmas(words)
	words = iterateDiagonallyLeft(lines, false)
	count += countXmas(words)
	words = iterateDiagonallyLeft(lines, true)
	count += countXmas(words)
	words = iterateDiagonallyRight(lines, false)
	count += countXmas(words)
	words = iterateDiagonallyRight(lines, true)
	count += countXmas(words)
	log.Println(count)
}
