package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func priority(ch rune) int {
	var res int

	if ch > 'a' {
		res = int(ch-'a') + 1
	} else {
		res = int(ch-'A') + 27
	}

	return res
}

// Part 1: find repeating items in two compartments and sum up
// priorities for repeating items for each rucksack and in total
func reorganizeRucksacks(input []byte) {
	prioritySum := 0
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		repeatedChars := make(map[rune]int)

		fmt.Printf("Line length: %d\nLine full: %s\nLine half: %s\n", len(line), line, line[0:len(line)/2])

		for _, left := range line[0 : len(line)/2] {
			for _, right := range line[len(line)/2:] {
				if left == right {
					repeatedChars[left] = priority(left)
				}
			}
		}
		for k, v := range repeatedChars {
			fmt.Printf("(k=%q, v=%d) ", k, v)
			prioritySum += v
		}
	}

	fmt.Printf("\nTotal sum of priorities of repeated types is %d\n", prioritySum)
}

// Part 2: find a repeating character in the groups of three lines,
// and find the total sum of all priorities of such characters
func groupBadges(input []byte) {
	prioritySum := 0

	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		rucksackHash := make(map[rune]int)
		for _, ch := range scanner.Text() {
			rucksackHash[ch] = 0
		}

		// read the next two lines
		otherTwoRucksacks := make([]string, 2)
		for i := 0; i < 2; i++ {
			if !scanner.Scan() {
				fmt.Fprintln(os.Stderr, "Error reading input: wrong number of elf rucksacks")
				os.Exit(1)
			}
			otherTwoRucksacks[i] = scanner.Text()
		}

		// scan the other two rucksacks and check if they contain a item from the first one (stored in rucksackHash)
		// then update the value in rucksackHash to save inforamtion about which item can be found in all three rucksacks
		for i := 0; i < 2; i++ {
			for _, ch := range otherTwoRucksacks[i] {
				if val, exists := rucksackHash[ch]; exists {
					// This one is tricky, if the character exists in the rucksack hash and its value equals current index,
					// it means it was added to the hash when checking the previous rucksack. So if the value equals 0 then
					// it was in the original hash, if it equals 1, it exists in the original hash and the second rucksack
					if val == i {
						rucksackHash[ch] = i + 1
					}
				}
			}
		}

		for k, v := range rucksackHash {
			// values could be 0, 1, 2, if it equals 2 it means it exists in all three rucksacks
			if v == 2 {
				prioritySum += priority(k)
			}
		}
	}

	fmt.Printf("Total sum of priorities for group badges is %d\n", prioritySum)
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Can't read input file: %s\n", err)
		os.Exit(1)
	}

	reorganizeRucksacks(input)
	groupBadges(input)
}
