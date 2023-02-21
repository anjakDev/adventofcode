package main

import (
	"adventofcode/util"
	"strconv"
)

func partone() int {
	fileScanner, closer := util.NewScanner("input-day1.txt")
	defer closer.Close()

	var currentMax int
	var sumCaloriesForElf int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			if currentMax < sumCaloriesForElf {
				currentMax = sumCaloriesForElf
			}
			sumCaloriesForElf = 0
			continue
		}
		calorie, _ := strconv.Atoi(line)
		sumCaloriesForElf += calorie
	}

	return currentMax
}
