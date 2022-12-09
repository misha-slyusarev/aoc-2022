package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Find the number of characters that needed to be processed
// before the message of arbitrary length was found
func findMessage(input []byte, messageLength uint) uint {
	var i, j uint

	h := make(map[byte]byte)
	n := uint(len(input))

	i, j = 0, 0
	for i < n {
		if j < n && h[input[j]] == 0 {
			if j+1-i == messageLength {
				break
			}
			h[input[j]]++
			j++
		} else {
			h[input[i]]--
			i++
		}
	}
	return j + 1
}

// Start of Packet
func findSOP(input []byte) {
	fmt.Printf("The number of characters processed before SOP found is %d\n", findMessage(input, 4))
}

// Start of Message
func findSOM(input []byte) {
	fmt.Printf("The number of characters processed before SOM found is %d\n", findMessage(input, 14))
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading input: %s", err)
		os.Exit(1)
	}

	findSOP(input)
	findSOM(input)
}
