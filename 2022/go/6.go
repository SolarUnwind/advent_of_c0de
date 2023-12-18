package main

import (
	"fmt"
	"os"
)

func main() {

	input, err := os.ReadFile("../6.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	
	line := string(input)

	index := 0
	for i := 0; i < len(line); i++ {
		chunk := map[string]bool{}

		for j := i; j < i + 4; j++ {
			chunk[string(line[j])] = true
		}

		if len(chunk) == 4 {
			break
		}
		index += 1
	}
	//
	fmt.Printf("%v\n", index + 4)

	index = 0
	for i := 0; i < len(line); i++ {
		chunk := map[string]bool{}

		for j := i; j < i + 14; j++ {
			chunk[string(line[j])] = true
		}

		if len(chunk) == 14 {
			break
		}
		index += 1
	}

	fmt.Printf("%v\n", index + 14)

}