package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s", err)
		os.Exit(1)
	}

	currentProvision := 0
	maxProvision := 0

	for lineNumber, line := range strings.Split(string(input), "\n") {
		if line == "" {
			if currentProvision > maxProvision {
				maxProvision = currentProvision
			}
			currentProvision = 0
		} else {
			provision, err := strconv.Atoi(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error while parsing the input file: expected a number on line %d, %s", lineNumber, err)
				os.Exit(1)
			}
			currentProvision += provision
		}
	}

	fmt.Printf("The largest provision carried by an elf is %d\n", maxProvision)
}
