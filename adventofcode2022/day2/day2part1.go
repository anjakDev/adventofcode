package main

import (
	"adventofcode/util"
)

func partone() int {
	fileScanner, closer := util.NewScanner("input-day2.txt")
	defer closer.Close()

	scoringMatrix := make([][]int, 3)
	scoringMatrix[0] = []int{3, 6, 0}
	scoringMatrix[1] = []int{0, 3, 6}
	scoringMatrix[2] = []int{6, 0, 3}

	var totalScore int

	for fileScanner.Scan() {
		line := []rune(fileScanner.Text())
		letterScore := getRowIdx(line[2]) + 1
		gameScore := scoringMatrix[getColIdx(line[0])][getRowIdx(line[2])]
		totalScore += letterScore + gameScore
	}

	return totalScore
}

func getColIdx(letter rune) int {
	var idx int
	switch letter {
	case 'A':
		idx = 0
	case 'B':
		idx = 1
	case 'C':
		idx = 2
	}
	return idx
}

func getRowIdx(letter rune) int {
	var idx int
	switch letter {
	case 'X':
		idx = 0
	case 'Y':
		idx = 1
	case 'Z':
		idx = 2
	}
	return idx
}
