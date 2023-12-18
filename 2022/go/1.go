package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {

	input, err := os.Open("../1.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	sc := bufio.NewScanner(input)

	var totals []int
	var cur_index int = 0
	var cur_value int = 0
	for sc.Scan() {
		if v := len(sc.Text()); v > 0 {
			val, _ := strconv.Atoi(sc.Text())
			cur_value += val
		} else {
            totals = append(totals, cur_value)
            cur_index += 1
            cur_value = 0
		}
	}
	sort.Ints(totals)
	fmt.Printf("Max Value: %v\n", totals[len(totals) - 1])	
	top3 := totals[len(totals) - 1] + totals[len(totals) - 2] + totals[len(totals) - 3]
	fmt.Printf("Top 3 Value: %v\n", top3)	

}