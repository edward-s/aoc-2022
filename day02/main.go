package main

import (
	"fmt"
	"strings"

	"github.com/edward-s/aoc-2022/utils"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

const (
	WIN  = 6
	DRAW = 3
	LOSE = 0
)

var opponentMap = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}
var mineMap = map[string]int{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var outcomeMap = map[string]int{
	"X": LOSE,
	"Y": DRAW,
	"Z": WIN,
}

func calculateScore(me, opp int) int {
	if me == SCISSORS && opp == PAPER {
		return WIN + SCISSORS
	}
	if me == ROCK && opp == SCISSORS {
		return WIN + ROCK
	}
	if me == PAPER && opp == ROCK {
		return WIN + PAPER
	}
	if me == opp {
		return DRAW + me
	}
	return LOSE + me
}

func calculateMyMove(opp, outcome int) int {
	loseMap := map[int]int{
		ROCK:     SCISSORS,
		PAPER:    ROCK,
		SCISSORS: PAPER,
	}
	winMap := map[int]int{
		ROCK:     PAPER,
		PAPER:    SCISSORS,
		SCISSORS: ROCK,
	}

	if outcome == LOSE {
		mine := loseMap[opp]
		return calculateScore(mine, opp)
	} else if outcome == WIN {
		mine := winMap[opp]
		return calculateScore(mine, opp)
	}
	return DRAW + opp
}

func part1(lines []string) {
	score := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		opponent := opponentMap[parts[0]]
		mine := mineMap[parts[1]]

		score += calculateScore(mine, opponent)
	}

	fmt.Println(score)
}

func part2(lines []string) {
	score := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		opponent := opponentMap[parts[0]]
		outcome := outcomeMap[parts[1]]

		score += calculateMyMove(opponent, outcome)
	}

	fmt.Println(score)
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	part1(lines)
	part2(lines)
}
