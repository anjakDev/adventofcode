package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

func partone() int {
	fileScanner, closer := util.NewScanner("input-day10.txt")
	defer closer.Close()

	sumSignalStrengths := 0

	xRegister := 1
	cycle := 0
	for fileScanner.Scan() {
		// instruction can have two parts: "command value" where value is optional
		instruction := strings.Split(fileScanner.Text(), " ")
		if len(instruction) == 1 {
			cycle++
			sumSignalStrengths += getSignalStrength(cycle, xRegister)
			continue
		}
		value, _ := strconv.Atoi(instruction[1])
		cycle++
		sumSignalStrengths += getSignalStrength(cycle, xRegister)
		cycle++
		sumSignalStrengths += getSignalStrength(cycle, xRegister)
		xRegister += value
	}
	return sumSignalStrengths
}

func getSignalStrength(cycle int, xRegister int) int {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return cycle * xRegister
	}
	return 0
}
