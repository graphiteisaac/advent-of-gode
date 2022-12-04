package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Row struct {
	L string
	R string
}

const ROCK = "A"
const PAPER = "B"
const SCISSORS = "C"

var ShapeScore = map[string]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

var WinningMoves = map[string]string{
	ROCK:     PAPER,
	PAPER:    SCISSORS,
	SCISSORS: ROCK,
}

var LosingMoves = map[string]string{
	PAPER:    ROCK,
	SCISSORS: PAPER,
	ROCK:     SCISSORS,
}

func getRPSOutcome(lhs, rhs string) int {
	if LosingMoves[lhs] == rhs {
		return 0
	}
	if lhs == rhs {
		return 3
	}
	return 6
}

func getScoreOne(lhs string, rhs string) int {
	// In this round, we are assuming X Y and Z are our counter move.
	counter := map[string]string{
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}

	return ShapeScore[counter[rhs]] + getRPSOutcome(lhs, counter[rhs])
}

func getScoreTwo(lhs string, rhs string) int {
	// In this round, we know the X Y and Z values refer to a desired outcome, and we need to shift our move respectively.
	const LOSE = "X"
	const DRAW = "Y"
	const WIN = "Z"

	if rhs == WIN {
		move := WinningMoves[lhs]
		return ShapeScore[move] + getRPSOutcome(lhs, move)
	} else if rhs == LOSE {
		move := LosingMoves[lhs]
		return ShapeScore[move] + getRPSOutcome(lhs, move)
	} else if rhs == DRAW {
		return ShapeScore[lhs] + getRPSOutcome(lhs, lhs)
	}

	return 3247239892
}

func processData() (rows []Row) {
	file, err := os.Open("./day-02/input.txt")
	if err != nil {
		log.Fatal("could not open file, ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		rows = append(rows, Row{
			L: moves[0],
			R: moves[1],
		})
	}

	return
}

func partOne(rows *[]Row) int {
	fmt.Println("Advent of Code #2 - Part 1")

	var score = 0
	for _, row := range *rows {
		score = score + getScoreOne(row.L, row.R)
	}

	return score
}

func partTwo(rows *[]Row) int {
	fmt.Println("Advent of Code #2 - Part 2")

	var score = 0
	for _, row := range *rows {
		score = score + getScoreTwo(row.L, row.R)
	}

	return score
}

func main() {
	rows := processData()

	firstScore := partOne(&rows)
	fmt.Printf("First: %d\n", firstScore)

	secondScore := partTwo(&rows)
	fmt.Printf("Second: %d\n", secondScore)
}
