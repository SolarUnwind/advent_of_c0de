package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func main() {

	input, err := os.ReadFile("../7.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	sc := string(input)
	lines := strings.Split(sc, "\n")

	paths := map[string]uint32{}

	path := []string{}
	for i := 0; i < len(lines); i++ {
		parsed_line := strings.Split(lines[i], " ")

		if parsed_line[1] == "cd" {
            if parsed_line[2] == ".." {
                path = path[:len(path)- 1]
            } else {
                path = append(path, parsed_line[2])
            }
        } else if parsed_line[1] == "ls" {
            continue
        } else if parsed_line[0] == "dir" {
            continue
        } else {
			conv_size,_ := strconv.ParseUint(parsed_line[0], 10, 32)
			size := uint32(conv_size)
			cur_path := ""
			for j := 0; j < len(path); j++ {
				cur_path += path[j]
				paths[cur_path] += size		
            }
        }     
	}


	var score_part1 uint32 = 0
	var score_part2 uint32 = 4294967295

	var limit uint32 = 100000
    var max uint32 = uint32(70000000) - uint32(30000000)
    var used uint32 = paths["/"]
    var need_to_free uint32 = used - max
	fmt.Printf("need_to_free: %v\n", need_to_free)	
	fmt.Printf("used: %v\n", used)	
	fmt.Printf("max: %v\n", max)	


	for _, val := range paths {
        if val <= limit {
			score_part1 += val
		} 
		if val >= need_to_free {
			score_part2 = min(val, score_part2)
		}
    }
	
	fmt.Printf("Score 1: %v\n", score_part1)	
	fmt.Printf("Score 2: %v\n", score_part2)
}