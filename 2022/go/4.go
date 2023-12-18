package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	input, err := os.ReadFile("../4.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	sc := string(input)
	lines := strings.Split(sc, "\n")

	var score_part1 int = 0
	var score_part2 int = 0

	for i := 0; i < len(lines); i++ {
		pair := strings.Split(lines[i], ",")
		first := strings.Split(pair[0], "-")
		second := strings.Split(pair[1], "-")

		first_start,_ :=  strconv.Atoi(first[0])
        first_end,_ := strconv.Atoi(first[1])
        second_start,_ := strconv.Atoi(second[0]) 
        second_end,_ := strconv.Atoi(second[1]) 

		if (first_start <= second_start && first_end >= second_end) || ( second_start <= first_start && second_end >= first_end ) {
            score_part1 += 1;
        }

        if (first_start <= second_end && second_start <= first_end) {
            score_part2 += 1;
        }
	}

	
	fmt.Printf("Score 1: %v\n", score_part1)	
	fmt.Printf("Score 2: %v\n", score_part2)	
}