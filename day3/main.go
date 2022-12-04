package main

import (
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

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Can't read input file: %s\n", err)
		os.Exit(1)
	}

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

	fmt.Printf("\nSum of priorities of repeated types is %d\n", prioritySum)
}
