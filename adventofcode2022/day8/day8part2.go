package main

import (
	"adventofcode/util"
	"strconv"
)

func parttwo() int {
	fileScanner, closer := util.NewScanner("input-day8.txt")
	defer closer.Close()

	treeGrid := make([][]int, 0)

	row := 0
	for fileScanner.Scan() {
		treeGrid = append(treeGrid, make([]int, 0))
		line := fileScanner.Text()
		for i := range line {
			digit, _ := strconv.Atoi(string(line[i]))
			treeGrid[row] = append(treeGrid[row], digit)
		}
		row++
	}

	counterArray := make([]int, 0)

	row = 0
	col := 0
	for row < len(treeGrid) {
		for col < len(treeGrid[0]) {
			num := treeGrid[row][col]
			var countLeft, countRight, countUp, countDown int

			if col > 0 {
				countLeft = countScoreLeft(treeGrid, col, row, num)
			}
			if col < len(treeGrid[0])-1 {
				countRight = countScoreRight(treeGrid, col, row, num)
			}
			if row > 0 {
				countUp = countScoreUp(treeGrid, col, row, num)
			}
			if row < len(treeGrid)-1 {
				countDown = countScoreDown(treeGrid, col, row, num)
			}
			counterArray = append(counterArray, countLeft*countRight*countUp*countDown)
			col++
		}
		col = 1
		row++
	}

	var max, temp int
	for _, element := range counterArray {
		if element > temp {
			temp = element
			max = temp
		}
	}
	return max
}

func countScoreLeft(treeGrid [][]int, col int, row int, num int) int {
	counter := 0
	col = col - 1
	for col >= 0 {
		counter++
		if treeGrid[row][col] >= num {
			return counter
		}
		col--
	}
	return counter
}

func countScoreRight(treeGrid [][]int, col int, row int, num int) int {
	counter := 0
	col = col + 1
	for col < len(treeGrid[0]) {
		counter++
		if treeGrid[row][col] >= num {
			return counter
		}
		col++
	}
	return counter
}

func countScoreUp(treeGrid [][]int, col int, row int, num int) int {
	counter := 0
	row = row - 1
	for row >= 0 {
		counter++
		if treeGrid[row][col] >= num {
			return counter
		}
		row--
	}
	return counter
}

func countScoreDown(treeGrid [][]int, col int, row int, num int) int {
	counter := 0
	row = row + 1
	for row < len(treeGrid) {
		counter++
		if treeGrid[row][col] >= num {
			return counter
		}
		row++
	}
	return counter
}
