package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func partone() int {
	fileScanner, closer := util.NewScanner("input-day9.txt")
	defer closer.Close()

	input := make([]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		input = append(input, line)
	}

	visitedMap := make(map[Position]int)

	posH := Position{0, 0}
	posT := Position{0, 0}
	visitedMap[posT] = 1

	for _, line := range input {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		amount, _ := strconv.Atoi(splitLine[1])
		switch direction {
		case "R":
			posH, posT = moveRight(amount, posH, posT, visitedMap)
		case "L":
			posH, posT = moveLeft(amount, posH, posT, visitedMap)
		case "U":
			posH, posT = moveUp(amount, posH, posT, visitedMap)
		case "D":
			posH, posT = moveDown(amount, posH, posT, visitedMap)
		}
	}

	return len(visitedMap)
}

func moveRight(amount int, posH Position, posT Position, visitedMap map[Position]int) (Position, Position) {
	i := 0
	for i < amount {
		posH.x++
		if posH.y > posT.y && (posH.x > posT.x+1 || posH.x < posT.x-1) {
			posT.y++
			posT.x++
			visitedMap[posT] = 1
			i++
			continue
		}
		if posH.y < posT.y && (posH.x > posT.x+1 || posH.x < posT.x-1) {
			posT.y--
			posT.x++
			visitedMap[posT] = 1
			i++
			continue
		}
		if posH.y == posT.y && (posH.x > posT.x+1 || posH.x < posT.x-1) {
			posT.x++
			visitedMap[posT] = 1
			i++
			continue
		}
		i++
	}
	return posH, posT
}

func moveLeft(amount int, posH Position, posT Position, visitedMap map[Position]int) (Position, Position) {
	i := 0
	for i < amount {
		posH.x--
		if posH.y > posT.y && (posH.x > posT.x+1 || posH.x < posT.x-1) {
			posT.y++
			posT.x--
			visitedMap[posT] = 1
			i++
			continue
		}
		if posH.y < posT.y && (posH.x > posT.x+1 || posH.x < posT.x-1) {
			posT.y--
			posT.x--
			visitedMap[posT] = 1
			i++
			continue
		}
		if posH.y == posT.y && (posH.x > posT.x+1 || posH.x < posT.x-1) {
			posT.x--
			visitedMap[posT] = 1
			i++
			continue
		}
		i++
	}
	return posH, posT
}

func moveUp(amount int, posH Position, posT Position, visitedMap map[Position]int) (Position, Position) {
	i := 0
	for i < amount {
		posH.y++
		if posH.x > posT.x && (posH.y > posT.y+1 || posH.y < posT.y-1) {
			posT.x++
			posT.y++
			visitedMap[posT] = 1
			i++
			continue
		}
		if posH.x < posT.x && (posH.y > posT.y+1 || posH.y < posT.y-1) {
			posT.x--
			posT.y++
			visitedMap[posT] = 1
			i++
			continue
		}
		if posH.x == posT.x && (posH.y > posT.y+1 || posH.y < posT.y-1) {
			posT.y++
			visitedMap[posT] = 1
			i++
			continue
		}
		i++
	}
	return posH, posT
}

func moveDown(amount int, posH Position, posT Position, visitedMap map[Position]int) (Position, Position) {
	i := 0
	for i < amount {
		posH.y--
		if posH.x > posT.x && (posH.y > posT.y+1 || posH.y < posT.y-1) {
			posT.x++
			posT.y--
			visitedMap[posT] = 1
			i++
			continue
		}
		if posH.x < posT.x && (posH.y > posT.y+1 || posH.y < posT.y-1) {
			posT.x--
			posT.y--
			visitedMap[posT] = 1
			i++
			continue
		}
		if posH.x == posT.x && (posH.y > posT.y+1 || posH.y < posT.y-1) {
			posT.y--
			visitedMap[posT] = 1
			i++
			continue
		}
		i++
	}
	return posH, posT
}
