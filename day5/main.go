package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type stack struct {
	top      int
	elements []rune
}

func (s *stack) push(r rune) {
	if s.top == len(s.elements)-1 {
		s.elements = append(s.elements, r)
		s.top++
	} else {
		s.top++
		s.elements[s.top] = r
	}
}

func (s *stack) pop() rune {
	var r rune
	if s.top >= 0 {
		r = s.elements[s.top]
		s.top--
	}

	return r
}

func newStack() stack {
	return stack{-1, make([]rune, 10)}
}

func parseInitialArrangement(initAr []string) []stack {
	var stacks []stack
	var stackPositions []int

	lenAr := len(initAr)

	// parse the bottom line and remember position of every non-empty character,
	// which identifies where we can find data in the rest of the input
	for pos, char := range initAr[lenAr-1] {
		if !unicode.IsSpace(char) {
			stackPositions = append(stackPositions, pos)
		}
	}

	// allocate space for stacks
	stacks = make([]stack, len(stackPositions))
	for i := range stacks {
		stacks[i] = newStack()
	}

	// fill in stacks according to initial arrangement
	for lineNm := lenAr - 2; lineNm >= 0; lineNm-- {
		for stkNm, cratePos := range stackPositions {
			r := rune(initAr[lineNm][cratePos])
			if unicode.IsLetter(r) {
				stacks[stkNm].push(r)
			}
		}
	}

	return stacks
}

type command struct {
	count int
	src   int
	dst   int
}

func newCommand(count string, src string, dst string) command {
	c, err := strconv.Atoi(count)
	if err != nil {
		fmt.Printf("Error parsing command: %s", err)
	}
	s, err := strconv.Atoi(src)
	if err != nil {
		fmt.Printf("Error parsing command: %s", err)
	}
	d, err := strconv.Atoi(dst)
	if err != nil {
		fmt.Printf("Error parsing command: %s", err)
	}
	return command{
		count: c,     // number of objects to move
		src:   s - 1, // source index is 0-based
		dst:   d - 1, // destination index is 0-based
	}
}

func parseCommand(s string) command {
	re := regexp.MustCompile(`[[:digit:]]+`)
	found := re.FindAllString(s, -1)

	return newCommand(found[0], found[1], found[2])
}

func calculateRearrangement(input []byte) {
	var initialArrangement []string

	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		fmt.Println(line)
		initialArrangement = append(initialArrangement, line)
	}

	stacks := parseInitialArrangement(initialArrangement)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		command := parseCommand(line)

		for i := 0; i < command.count; i++ {
			stacks[command.dst].push(stacks[command.src].pop())
		}
	}

	fmt.Print("Resulting top level crates: ")
	for i := range stacks {
		fmt.Print(string(stacks[i].pop()))
	}
	fmt.Println()
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s", err)
	}

	calculateRearrangement(input)
}
