package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./puzzleinput.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
}
