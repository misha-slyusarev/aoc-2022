package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Find the start-of-packet marker
func findSop(input []byte) {
	i, j := 0, 0
	h := make(map[byte]byte)

	for i < len(input) {
		if j < len(input) && h[input[j]] == 0 {
			if j+1-i == 4 {
				break
			}
			h[input[j]]++
			j++
		} else {
			h[input[i]]--
			i++
		}
	}
	fmt.Printf("The number of characters processed before SOP found is %d\n", j+1)
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading input: %s", err)
		os.Exit(1)
	}

	findSop(input)
}
