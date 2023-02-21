package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

func parttwo() int {
	fileScanner, closer := util.NewScanner("input-day4.txt")
	defer closer.Close()

	var counter int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		assignments := strings.Split(line, ",")
		firstAssignment := strings.Split(assignments[0], "-")
		secondAssignment := strings.Split(assignments[1], "-")

		firstSectionFirst, _ := strconv.Atoi(firstAssignment[0])
		secondSectionFirst, _ := strconv.Atoi(secondAssignment[0])

		firstSectionSecond, _ := strconv.Atoi(firstAssignment[1])
		secondSectionSecond, _ := strconv.Atoi(secondAssignment[1])

		if firstSectionFirst <= secondSectionFirst && firstSectionSecond >= secondSectionFirst {
			counter++
			continue
		}
		if secondSectionFirst <= firstSectionFirst && secondSectionSecond >= firstSectionFirst {
			counter++
			continue
		}
	}

	return counter
}
