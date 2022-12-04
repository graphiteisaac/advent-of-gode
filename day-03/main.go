package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rucksack struct {
	First   string
	Second  string
	Overlap string
	Score   int
}

func getScore(provided rune) int {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i, char := range chars {
		if char == provided {
			return i + 1
		}
	}
	return 0
}

func getCommonScore(sacks []Rucksack) rune {
	var chars []rune
	for i, s := range sacks {
		if len(chars) == 0 {
			chars = []rune(s.First + s.Second)
		}

		newChars := ""
		for _, c := range chars {
			if i != len(sacks) && strings.Contains(sacks[i].First+sacks[i].Second, string(c)) && !strings.Contains(newChars, string(c)) {
				newChars = newChars + string(c)
			}
		}
		chars = []rune(newChars)
	}

	return chars[0]
}

func main() {
	file, err := os.Open("./day-03/input.txt")
	if err != nil {
		log.Fatal("could not open file, ", err)
	}
	defer file.Close()

	var sacks []Rucksack
	var total int
	var badgeTotal int

	i := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		letters := []rune(scanner.Text())
		firstItem := letters[:len(letters)/2]
		secondItem := string(letters[len(letters)/2:])
		var matched rune

		for _, r := range firstItem {
			if strings.Contains(secondItem, string(r)) {
				matched = r
			}
		}

		score := getScore(matched)

		sacks = append(sacks, Rucksack{
			First:   string(firstItem),
			Second:  secondItem,
			Overlap: string(matched),
			Score:   score,
		})

		total = total + score

		if i%3 == 0 {
			badgeTotal = badgeTotal + getScore(getCommonScore(sacks[len(sacks)-3:]))
		}
		i++
	}

	fmt.Println("Total: ", total)
	fmt.Println("Badge Total: ", badgeTotal)
}
