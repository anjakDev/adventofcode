package main

import (
	"adventofcode/util"
	"strconv"
)

func partone() int {
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
	countEdges := 2*len(treeGrid) + (len(treeGrid[0])-2)*2
	sumVisible := 0

	row = 1
	col := 1
	for row < len(treeGrid)-1 {
		for col < len(treeGrid[0])-1 {
			num := treeGrid[row][col]
			isVisibleLeft := checkLeft(treeGrid, col, row, num)
			if isVisibleLeft {
				sumVisible++
				col++
				continue
			}
			isVisbleRight := checkRight(treeGrid, col, row, num)
			if isVisbleRight {
				sumVisible++
				col++
				continue
			}
			isVisibleUp := checkUp(treeGrid, col, row, num)
			if isVisibleUp {
				sumVisible++
				col++
				continue
			}
			isVisibleDown := checkDown(treeGrid, col, row, num)
			if isVisibleDown {
				sumVisible++
				col++
				continue
			}
			col++
		}
		col = 1
		row++
	}

	return sumVisible + countEdges
}

func checkLeft(treeGrid [][]int, col int, row int, num int) bool {
	isVisible := true
	col = col - 1
	for col >= 0 {
		if treeGrid[row][col] >= num {
			return !isVisible
		}
		col--
	}
	return isVisible
}

func checkRight(treeGrid [][]int, col int, row int, num int) bool {
	isVisible := true
	col = col + 1
	for col < len(treeGrid[0]) {
		if treeGrid[row][col] >= num {
			return !isVisible
		}
		col++
	}
	return isVisible
}

func checkUp(treeGrid [][]int, col int, row int, num int) bool {
	isVisible := true
	row = row - 1
	for row >= 0 {
		if treeGrid[row][col] >= num {
			return !isVisible
		}
		row--
	}
	return isVisible
}

func checkDown(treeGrid [][]int, col int, row int, num int) bool {
	isVisible := true
	row = row + 1
	for row < len(treeGrid) {
		if treeGrid[row][col] >= num {
			return !isVisible
		}
		row++
	}
	return isVisible
}
