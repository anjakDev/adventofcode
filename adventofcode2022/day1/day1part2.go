package main

import (
	"adventofcode/util"
	"sort"
	"strconv"
)

type Elf struct {
	Calories int
}

func parttwo() int {
	fileScanner, closer := util.NewScanner("input-day1.txt")
	defer closer.Close()

	var sumCaloriesForElf int
	elves := make([]Elf, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			elf := Elf{Calories: sumCaloriesForElf}
			elves = append(elves, elf)
			sumCaloriesForElf = 0
			continue
		}
		calorie, _ := strconv.Atoi(line)
		sumCaloriesForElf += calorie
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})
	return elves[0].Calories + elves[1].Calories + elves[2].Calories
}
