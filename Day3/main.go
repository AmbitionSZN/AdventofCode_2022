package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInput() [][]rune {
	content, err := os.ReadFile("./puzzleinput.txt")

	if err != nil {
		log.Fatalln("error reading file")
	}

	data := strings.Split(string(content), "\r\n")
	runeData := make([][]rune, len(data))

	for i, v := range data {
		runeData[i] = []rune(v)
	}
	return runeData
}

func sumItemTypes(rucksacks [][]rune) int {
	matches := []rune{}
	for _, v := range rucksacks {
		compartmentOneLen := len(v) / 2
		compartmentOne := v[:compartmentOneLen]
		compartmentTwo := v[compartmentOneLen:]
		for _, valueOne := range compartmentOne {
			bool := false
			for _, valueTwo := range compartmentTwo {
				if valueOne == valueTwo {
					matches = append(matches, valueOne)
					bool = true
					break
				}
			}
			if bool {
				break
			}
		}
	}

	fmt.Println(matches)
	total := 0

	for _, v := range matches {
		priority := int(v) - 96
		if priority < 1 {
			total += priority + 58
		} else {
			total += priority
		}
	}
	return total
}

func main() {
	fmt.Println(sumItemTypes(parseInput()))
}
