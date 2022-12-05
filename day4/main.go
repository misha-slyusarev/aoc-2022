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

	total := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		// line example: 2-8,3-7
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
		fLeft := assignments[0]
		fRight := assignments[1]
		sLeft := assignments[2]
		sRight := assignments[3]

		if fLeft >= sLeft && fRight <= sRight || sLeft >= fLeft && sRight <= fRight {
			total++
		}
	}
	fmt.Printf("Total number of assignments that fully included by other assignments is %d\n", total)
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s", err)
		os.Exit(1)
	}

	repeatedAssignments(input)
}
