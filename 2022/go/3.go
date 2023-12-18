package main

import (
	"fmt"
	"os"
	"strings"
)

func char_to_priority(input int) int {
	lower_case := 96
    upper_case := 38
	result := 0

    if strings.ToLower(string(input)) != string(input) {
        result =  input - upper_case
    } else {
        result = input - lower_case
    }

	return result
}

func main() {

	input, err := os.ReadFile("../3.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	sc := string(input)
	lines := strings.Split(sc, "\n")

	var score_part1 int = 0
	var score_part2 int = 0

	for i := 0; i < len(lines); i++ {

		n := len(lines[i]) / 2
		left := lines[i][0:n]
		right := lines[i][n:]
		set := make(map[int]bool) 

		for _, char := range left {	
			set[int(char)] = true
		}

		for k := range set {
			if strings.Contains(right, string(k)) {
				score_part1 += char_to_priority(k)
			}
		}
	}

	for i := 0; i < len(lines); i += 3 {

		first := lines[i]
		set := make(map[int]bool) 

		for _, char := range first {	
			set[int(char)] = true
		}

		for k := range set {
			if strings.Contains(lines[i + 1], string(k)) && strings.Contains(lines[i + 2], string(k)){
				score_part2 += char_to_priority(k)
			}
		}
	}
	
	fmt.Printf("Score 1: %v\n", score_part1)	
	fmt.Printf("Score 2: %v\n", score_part2)	
}