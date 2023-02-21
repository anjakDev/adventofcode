package main

import "adventofcode/util"

func parttwo() int {
	fileScanner, closer := util.NewScanner("input-day2.txt")
	defer closer.Close()

	choiceMatrix := make([][]int, 3)
	choiceMatrix[0] = []int{3, 1, 2}
	choiceMatrix[1] = []int{1, 2, 3}
	choiceMatrix[2] = []int{2, 3, 1}

	var totalScore int

	for fileScanner.Scan() {
		line := []rune(fileScanner.Text())
		gameScore := getRowIdx(line[2]) * 3
		letterScore := choiceMatrix[getColIdx(line[0])][getRowIdx(line[2])]
		totalScore += letterScore + gameScore
	}

	return totalScore
}
