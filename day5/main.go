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
	elements []byte
}

func (s *stack) push(ch byte) {
	if s.top == len(s.elements)-1 {
		s.elements = append(s.elements, ch)
		s.top++
	} else {
		s.top++
		s.elements[s.top] = ch
	}
}

func (s *stack) pop() byte {
	var ch byte
	if s.top >= 0 {
		ch = s.elements[s.top]
		s.top--
	}

	return ch
}

func (s stack) copy() stack {
	newStack := stack{-1, make([]byte, len(s.elements))}
	for i, e := range s.elements {
		newStack.elements[i] = e
	}
	newStack.top = s.top
	return newStack
}

func newStack() stack {
	return stack{-1, make([]byte, 10)}
}

func copyStacks(stacks []stack) []stack {
	copies := make([]stack, len(stacks))
	for i, s := range stacks {
		copies[i] = s.copy()
	}
	return copies
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
			ch := initAr[lineNm][cratePos]
			if unicode.IsLetter(rune(ch)) {
				stacks[stkNm].push(ch)
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

func simulateRearrangement(input []byte) {
	var initialArrangement []string
	var commands []command

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

	initialStacks := parseInitialArrangement(initialArrangement)
	stacks := copyStacks(initialStacks)

	for scanner.Scan() {
		line := scanner.Text()
		command := parseCommand(line)
		commands = append(commands, command)

		for i := 0; i < command.count; i++ {
			stacks[command.dst].push(stacks[command.src].pop())
		}
	}

	fmt.Print("Top level names after moving each individual box: ")
	for i := range stacks {
		fmt.Print(string(stacks[i].pop()))
	}
	fmt.Println()

	stacks = copyStacks(initialStacks)
	for _, command := range commands {
		tmpStack := newStack()
		for i := 0; i < command.count; i++ {
			tmpStack.push(stacks[command.src].pop())
		}
		for i := 0; i < command.count; i++ {
			stacks[command.dst].push(tmpStack.pop())
		}
	}

	fmt.Print("Top level names after moving several boxes at once: ")
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

	simulateRearrangement(input)
}
