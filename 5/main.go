package main

import (
	"bufio"
	"github.com/konradmalik/aoc2024/lib"
	"log"
	"os"
	"slices"
)

const InputFile string = "./input.txt"

func readInput(file *os.File) ([][]int, [][]int) {
	rules := make([][]int, 0)
	updates := make([][]int, 0)

	isRules := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRules = false
			continue
		}
		if isRules {
			rules = append(rules, lib.ParseNumbersSep(line, "|"))
		} else {
			updates = append(updates, lib.ParseNumbersSep(line, ","))
		}
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("scanner encountered an err: %s", err)
	}

	return rules, updates
}

func sortInPlace(rules [][]int, update []int) bool {
	modified := false
	for {
		changedIteration := false
		for _, rule := range rules {
			a := rule[0]
			b := rule[1]
			idxa := slices.Index(update, a)
			idxb := slices.Index(update, b)

			if idxa == -1 || idxb == -1 {
				continue
			}

			if idxa > idxb {
				tmp := update[idxa]
				update[idxa] = update[idxb]
				update[idxb] = tmp
				modified = true
				changedIteration = true
			}
		}
		if !changedIteration {
			return modified
		}
	}
}

func checkRules(rules [][]int, update []int) bool {
	for _, rule := range rules {
		a := rule[0]
		b := rule[1]
		idxa := slices.Index(update, a)
		idxb := slices.Index(update, b)

		if idxa == -1 || idxb == -1 {
			continue
		}

		if idxa > idxb {
			return false
		}
	}
	return true
}

func getMiddle(update []int) int {
	return update[(len(update)-1)/2]
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules, updates := readInput(file)

	sum := 0
	for _, update := range updates {
		log.Println(update)
		if sortInPlace(rules, update) {
			mid := getMiddle(update)
			log.Println(mid)
			sum += mid
		}
	}
	log.Println(sum)
}
