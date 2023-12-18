package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"strconv"
)

func main() {

	input, err := os.ReadFile("../5.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	sc := string(input)
	lines := strings.Split(sc, "\n")
	
	cmd_index := 0
	var stack []string
	for i := 0; i < len(lines); i++ {
		stack = append(stack, lines[i])
		if v := len(lines[i]); v <= 0 {
			cmd_index = i;
			break
		}
	}
	tmp := strings.Fields(stack[cmd_index - 1])
	size := len(tmp)
	stacks :=  map[int][]string{}
	for i := 0; i < size; i++ {
		stacks[i] = []string{} 
	}

	j := 2

	for j <= cmd_index {
		chars := stack[cmd_index - j]
		for k := 0; k < size; k++ {
			ch := chars[1 + 4*k]
			if unicode.IsLetter(rune(ch)) {
				stacks[k] = append(stacks[k], string(ch))
			}
		}
		j += 1
	}

	for i := 0; i < size; i++ {
		fmt.Printf("%v: %v\n",i, stacks[i])
	}
	fmt.Printf("After\n")
	// Process commands

	for cmd_index < len(lines) - 1{
		cmd_index += 1
		cmd := strings.Fields(lines[cmd_index])
		n,_ := strconv.Atoi(cmd[1])
		from,_ := strconv.Atoi(cmd[3])
		to,_ := strconv.Atoi(cmd[5])

		move := []string{}
		l := len(stacks[from - 1])
		for k := 0; k < n; k++ {

			move = append(move, stacks[from - 1][l - n + k])
		}
		
		stacks[from - 1] = stacks[from - 1][:len(stacks[from - 1]) - n]
		stacks[to - 1] = append(stacks[to - 1], move...)
	}

	for i := 0; i < size; i++ {
		fmt.Printf("%v: %v\n",i, stacks[i])
	}
}