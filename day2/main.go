package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
# General rules
A - Rock 		 	Rock 			- 1 point				Lost 	- 0
B - Paper   	Paper 		- 2 points			Draw 	- 3
C - Scissors	Scissors 	- 3 points			Win 	- 6
*/

type Game string

func newGame(input string) Game {
	return Game(strings.ReplaceAll(string(input), " ", ""))
}

/*
# First part rules
X - Rock
Y - Paper
Z - Scissors

# Game results
Win: CX, AY, BZ		- 6 points
Draw: AX, BY, CZ	- 3 points
Others: lost			- 0 points
*/
func (g Game) guesedResult() int {
	outcome := 0
	myShapeScore := 0

	switch g {
	case "CX", "AY", "BZ": // winning
		outcome = 6
	case "AX", "BY", "CZ": // draw
		outcome = 3
	}

	switch string(g[1]) {
	case "X":
		myShapeScore = 1
	case "Y":
		myShapeScore = 2
	case "Z":
		myShapeScore = 3
	}

	return outcome + myShapeScore
}

/*
# Second part rules
X - loose	- 0
Y - draw	- 3
Z - win		- 6

	X(0) - A(0) = 0 + 3
	  	 \ B(1) = 0 + 1
	  	 \ C(2) = 0 + 2

	Y(1) - A(0) = 3 + 1
	  	 \ B(1) = 3 + 2
	  	 \ C(2) = 3 + 3

	Z(2) - A(0) = 6 + 2
	  	 \ B(1) = 6 + 3
	  	 \ C(2) = 6 + 1
*/
func (g Game) decodedResult() int {
	oponnent := int(rune(g[0]) - 'A')
	me := int(rune(g[1]) - 'X')

	return me*3 + (me+oponnent+2)%3 + 1
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s", err)
		os.Exit(1)
	}

	guesedTotal := 0
	decodedTotal := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		game := newGame(line)
		guesedTotal += game.guesedResult()
		decodedTotal += game.decodedResult()
	}

	fmt.Printf("Guesed result: %d\nDecoded rules result: %d\n", guesedTotal, decodedTotal)
}
