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

func getShapeScore(shape string) int {
	if shape == "rock" {
		return 1
	} else if shape == "paper" {
		return 2
	}
	// Has to be C
	return 3
}

func getRPSOutcome(lhs, rhs string) int {
	num := 6
	if (lhs == "paper" && rhs == "rock") ||
		(lhs == "scissors" && rhs == "paper") ||
		(lhs == "rock" && rhs == "scissors") {
		num = 0
	}

	if lhs == rhs {
		num = 3
	}

	return num
}

func getScoreOne(lhs string, rhs string) int {
	// In this round, we are assuming X Y and Z are our counter move.
	shapes := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	return getShapeScore(shapes[rhs]) + getRPSOutcome(shapes[lhs], shapes[rhs])
}

func getLosingMove(lhs string) string {
	if lhs == "A" {
		return "scissors"
	} else if lhs == "B" {
		return "rock"
	} else if lhs == "C" {
		return "paper"
	}

	return "paper"
}

func getWinningMove(lhs string) string {
	if lhs == "A" {
		return "paper"
	} else if lhs == "B" {
		return "scissors"
	} else if lhs == "C" {
		return "rock"
	}

	return "paper"
}

func getScoreTwo(lhs string, rhs string) int {
	// In this round, we know the X Y and Z values refer to a desired outcome, and we need to shift our move respectively.
	const LOSE = "X"
	const DRAW = "Y"
	const WIN = "Z"

	shapes := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	if rhs == WIN {
		move := getWinningMove(lhs)
		return getShapeScore(move) + getRPSOutcome(shapes[lhs], move)
	} else if rhs == LOSE {
		move := getLosingMove(lhs)
		return getShapeScore(move) + getRPSOutcome(shapes[lhs], move)
	} else if rhs == DRAW {
		return getShapeScore(shapes[lhs]) + getRPSOutcome(shapes[lhs], shapes[lhs])
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
