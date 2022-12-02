package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Calories []int
	Total    int
}

func getMaxElf(arr []Elf) int {
	max := arr[0].Total
	for i := 0; i < len(arr); i++ {
		if arr[i].Total > max {
			max = arr[i].Total
		}
	}

	return max
}

func partOne() []Elf {
	fmt.Println("Advent Of Code #1 - Part 1")

	file, err := os.Open("./day-01/input.txt")
	if err != nil {
		log.Fatal("could not open file, ", err)
	}
	defer file.Close()

	var elves = []Elf{
		{
			Calories: []int{},
			Total:    0,
		},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		idx := len(elves) - 1

		calories, err := strconv.Atoi(line)
		if err == nil {
			elves[idx].Calories = append(elves[idx].Calories, calories)
			elves[idx].Total = elves[idx].Total + calories
		} else {
			elves = append(elves, Elf{
				Calories: []int{},
				Total:    0,
			})
		}
	}

	fmt.Println(getMaxElf(elves))
	return elves
}

func partTwo(elves []Elf) {
	fmt.Println("Advent of Code - #1 Part 2")

	sorted := elves
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Total > sorted[j].Total
	})

	remainder := sorted[:3]
	combinedTotal := 0
	for i, elf := range remainder {
		combinedTotal = combinedTotal + elf.Total
		fmt.Printf("#%d - %d\n", i, elf.Total)
	}
	fmt.Println(combinedTotal)
}

func main() {
	elves := partOne()

	partTwo(elves)
}
