package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

func parttwo() {
	fileScanner, closer := util.NewScanner("input-day10.txt")
	defer closer.Close()

	output := &strings.Builder{}

	xRegister := 1
	counter := 0
	for fileScanner.Scan() {
		// instruction can have two parts: "command value" where value is optional
		instruction := strings.Split(fileScanner.Text(), " ")
		if len(instruction) == 1 {
			counter = drawSymbol(output, xRegister, counter)
			continue
		}
		value, _ := strconv.Atoi(instruction[1])
		counter = drawSymbol(output, xRegister, counter)
		counter = drawSymbol(output, xRegister, counter)
		xRegister += value
	}
	fmt.Printf(output.String())
}

func drawSymbol(output *strings.Builder, xRegister int, counter int) int {
	var err error
	if xRegister-1 <= counter && counter <= xRegister+1 {
		_, err = output.WriteString("#")
		if err != nil {
			panic(err)
		}
	} else {
		_, err = output.WriteString(".")
		if err != nil {
			panic(err)
		}
	}
	if counter == 39 {
		_, err = output.WriteString("\n")
		if err != nil {
			panic(err)
		}
		return 0
	}
	return counter + 1
}
