package main

import (
	"fmt"
	"os"
	"strings"
)

func getTotalScore(rounds [][]rune) int {
	totalScore := 0
	for _, v := range rounds {
		fmt.Println(v)

		switch v[2] {
		case 88:
			switch v[0] {
			case 65:
				totalScore += 1 + 3
			case 66:
				totalScore += 1
			case 67:
				totalScore += 1 + 6
			}
		case 89:
			switch v[0] {
			case 65:
				totalScore += 2 + 6
			case 66:
				totalScore += 2 + 3
			case 67:
				totalScore += 2
			}
		case 90:
			switch v[0] {
			case 65:
				totalScore += 3
			case 66:
				totalScore += 3 + 6
			case 67:
				totalScore += 3 + 3
			}
		}

	}
	return totalScore
}

func parseInput() [][]rune {
	content, err := os.ReadFile("./puzzleinput.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	data := strings.Split(string(content), "\r\n")

	runeData := make([][]rune, len(data))

	for i, v := range data {
		runeData[i] = []rune(v)
	}

	return runeData
}

func main() {
	output := getTotalScore(parseInput())
	fmt.Println(output)
}
