package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	output := findMostCalories()
	fmt.Println(output)
}

func findMostCalories() int {
	content, err := os.ReadFile("./puzzleinput.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)

	}

	data := strings.Split(string(content), "\r\n\r\n")
	numData := make([]int, len(data))

	for i, v := range data {
		sum := 0
		reg := regexp.MustCompile(`[0-9]+`)
		matches := reg.FindAllString(v, -1)
		for _, v := range matches {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("str conversion failed")
				break
			}
			sum += num
		}
		numData[i] = sum
	}
	mostCalories := 0
	for _, v := range numData {
		if mostCalories < v {
			mostCalories = v
		}
	}
	return mostCalories

}
