package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}
type Move struct {
	P1 Position
	P2 Position
}

var letters map[string]int = make(map[string]int)
var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8}

func init() {
	letters["a"] = 1
	letters["b"] = 2
	letters["c"] = 3
	letters["d"] = 4
	letters["e"] = 5
	letters["f"] = 6
	letters["g"] = 7
	letters["h"] = 8
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	line := strings.Split(string(content), "\n")
	move, correctInput := parseInput(strings.Trim(line[0], "\r\n\t "))

	output := "ERROR\n"
	if !correctInput {
		ioutil.WriteFile("output.txt", []byte(output), 0)
		return
	}

	if !isPositionValid(move.P1) || !isPositionValid(move.P2) {
		ioutil.WriteFile("output.txt", []byte(output), 0)
		return
	}

	if checkTurn(move) {
		output = "YES\n"
	} else {
		output = "NO\n"
	}

	ioutil.WriteFile("output.txt", []byte(output), 0)
}

func parseInput(input string) (Move, bool) {
	var result Move

	if len(input) != 5 {
		return result, false
	}
	moves := strings.Split(input, "-")

	if len(moves) != 2 {
		return result, false
	}

	if len(moves[0]) != 2 || len(moves[1]) != 2 {
		return result, false
	}

	l := letters[strings.ToLower(moves[0][:1])]
	n, err := strconv.Atoi(moves[0][1:])

	if err != nil {
		return result, false
	}

	p1 := Position{l, n}

	l = letters[strings.ToLower(moves[1][:1])]
	n, err = strconv.Atoi(moves[1][1:])

	if err != nil {
		return result, false
	}

	p2 := Position{l, n}

	result = Move{p1, p2}

	return result, true
}

func checkTurn(turn Move) bool {
	dx := math.Abs(float64(turn.P1.X) - float64(turn.P2.X))
	dy := math.Abs(float64(turn.P1.Y) - float64(turn.P2.Y))

	if dx < 3 && dy < 3 && dx+dy == 3 {
		return true
	} else {
		return false
	}
}

func isPositionValid(p Position) bool {
	return p.X > 0 && p.X < 9 && p.Y > 0 && p.Y < 9
}
