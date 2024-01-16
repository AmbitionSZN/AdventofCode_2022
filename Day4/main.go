package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func comparePairs() int {
	content, err := os.ReadFile("./puzzleinput.txt")

	if err != nil {
		log.Fatalln("error reading file")
	}
	pairs := strings.Split(string(content), "\r\n")
	pairsArray := make([][]int, len(pairs))

	for i, pair := range pairs {
		split := strings.Split(pair, ",")
		vals1 := split[0]
		vals2 := split[1]
		vals1arr := strings.Split(vals1, "-")
		vals2arr := strings.Split(vals2, "-")
		vals1arr = append(vals1arr, vals2arr...)
		s := make([]int, len(vals1arr))
		for i, v := range vals1arr {
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("error")
			}
			s[i] = val
		}
		pairsArray[i] = s
	}
	fmt.Println(len(pairsArray))
	count := 0
	for _, pair := range pairsArray {
		boolean := true
		if pair[0] <= pair[2] && pair[1] >= pair[3] {
			count++
			boolean = false
		}
		if pair[2] <= pair[0] && pair[3] >= pair[1] && boolean {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(comparePairs())
}
