package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {

	input, err := os.Open("../2.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	sc := bufio.NewScanner(input)

	responses := map[string]int{ "X": 1, "Y": 2, "Z": 3,}
	part1 := map[string]map[string]int{
		"A": { "X": 3, "Y": 6, "Z": 0, },
        "B": { "X": 0, "Y": 3, "Z": 6, },
        "C": { "X": 6, "Y": 0, "Z": 3, },
	}

	results := map[string]int{ "X": 0, "Y": 3, "Z": 6,}
	part2 := map[string]map[string]int{
		"A": { "X": 3, "Y": 1, "Z": 2, },
        "B": { "X": 1, "Y": 2, "Z": 3, },
        "C": { "X": 2, "Y": 3, "Z": 1, },
	}

	var score_part1 int = 0
	var score_part2 int = 0

	for sc.Scan() {
		if v := len(sc.Text()); v > 0 {
			line := strings.Split(sc.Text(), " ")
			opponent := line[0]
            response := line[1]

			score_part1 += responses[response]
			score_part2 += results[response]

			score_part1 += part1[opponent][response]
			score_part2 += part2[opponent][response]

		}
	}
	
	fmt.Printf("Score 1: %v\n", score_part1)	
	fmt.Printf("Score 2: %v\n", score_part2)	
}