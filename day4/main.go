package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func repeatedAssignments(input []byte) {
	var err error

	rangesIncluded := 0
	rangesOverlap := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		// convert string input into an array of integers
		// which represent assignment ranges. Line example "2-8,3-7"
		assignments := make([]int, 4)
		for i, assignment := range strings.Split(line, ",") {
			for j, section := range strings.Split(assignment, "-") {
				assignments[j+i*2], err = strconv.Atoi(section)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error reading file: %s", err)
					os.Exit(1)
				}
			}
		}

		// the first elf assignment range
		fLeft := assignments[0]
		fRight := assignments[1]

		// the second elf assignment range
		sLeft := assignments[2]
		sRight := assignments[3]

		if fLeft >= sLeft && fRight <= sRight || sLeft >= fLeft && sRight <= fRight {
			rangesIncluded++
		}

		if fLeft <= sRight && sLeft <= fRight {
			rangesOverlap++
		}
	}
	fmt.Printf("Total number of assignments in which one range includes the other one is %d\n", rangesIncluded)
	fmt.Printf("Total number of assignments in which assignment ranges overlap is %d\n", rangesOverlap)
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s", err)
		os.Exit(1)
	}

	repeatedAssignments(input)
}
