package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func parttwo() string {
	fileScanner, closer := util.NewScanner("input-day5.txt")
	defer closer.Close()

	var crateStacks []string
	var instructions []string

	fillInstructions := false

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			fillInstructions = true
		}
		if fillInstructions {
			instructions = append(instructions, line)
			continue
		}
		crateStacks = append(crateStacks, line)
	}

	// Get amount of stacks from last line
	stackAmount := crateStacks[len(crateStacks)-1]
	stackAmountSlice := delete_empty(strings.Split(stackAmount, " "))

	stacks := make([]*Stack, len(stackAmountSlice))
	for i := range stacks {
		stacks[i] = new(Stack)
	}

	crateStacks = crateStacks[0 : len(crateStacks)-1]
	for i := len(crateStacks) - 1; i >= 0; i-- {
		line := crateStacks[i]
		var counter int
		for j := 0; j < len(line)-2; j += 4 {
			if line[j+1] != ' ' {
				stacks[counter].Push(string(line[j+1]))
			}
			counter++
		}
	}

	instructions = instructions[1:]
	for _, line := range instructions {
		var amount int
		var sourceStack int
		var destStack int
		fmt.Sscanf(line, "move %d from %d to %d", &amount, &sourceStack, &destStack)

		tempStack := new(Stack)
		for i := 0; i < amount; i++ {
			value := stacks[sourceStack-1].Pop()
			tempStack.Push(value.(string))
		}
		for tempStack.Len() > 0 {
			value := tempStack.Pop()
			stacks[destStack-1].Push(value.(string))
		}
	}
	var valueTops string
	for i := range stacks {
		valueTops += stacks[i].top.value
	}
	return valueTops
}
