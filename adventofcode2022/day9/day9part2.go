package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

type Knot struct {
	position Position
	next     *Knot
}

const knots = 10

func parttwo() int {
	fileScanner, closer := util.NewScanner("testinput-large.txt")
	defer closer.Close()

	input := make([]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		input = append(input, line)
	}

	visitedMap := make(map[Position]int)

	knotH := Knot{Position{0, 0}, nil}
	knotT := Knot{Position{0, 0}, nil}
	i := 0
	currKnot := &knotH
	for i < knots-2 {
		currKnot.next = &Knot{Position{0, 0}, nil}
		currKnot = currKnot.next
		i++
	}
	currKnot.position = Position{0, 0}
	currKnot.next = &knotT

	visitedMap[knotT.position] = 1

	for _, line := range input {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		amount, _ := strconv.Atoi(splitLine[1])
		knotH = move(direction, amount, knotH, visitedMap)
	}

	return len(visitedMap)
}

func move(direction string, amount int, knotH Knot, visitedMap map[Position]int) Knot {
	i := 0
	for i < amount {
		switch direction {
		case "R":
			knotH.position.x++
		case "L":
			knotH.position.x--
		case "U":
			knotH.position.y++
		case "D":
			knotH.position.y--
		}
		currentPrevKnot := knotH
		j := 0
		for j < knots {
			if currentPrevKnot.next == nil {
				visitedMap[currentPrevKnot.position] = 1
				break
			}

			xPosIsNotAdjacent := currentPrevKnot.position.x > currentPrevKnot.next.position.x+1 || currentPrevKnot.position.x < currentPrevKnot.next.position.x-1
			yPosIsNotAdjacent := currentPrevKnot.position.y > currentPrevKnot.next.position.y+1 || currentPrevKnot.position.y < currentPrevKnot.next.position.y-1
			knotPositionIsNotAdjacent := xPosIsNotAdjacent || yPosIsNotAdjacent

			if !knotPositionIsNotAdjacent {
				break
			}

			knotIsRightOfKnot := currentPrevKnot.position.x > currentPrevKnot.next.position.x
			knotIsLeftOfKnot := currentPrevKnot.position.x < currentPrevKnot.next.position.x

			knotIsAboveKnot := currentPrevKnot.position.y > currentPrevKnot.next.position.y
			knotIsBelowKnot := currentPrevKnot.position.y < currentPrevKnot.next.position.y

			knotIsSameRow := currentPrevKnot.position.y == currentPrevKnot.next.position.y
			knotIsSameColumn := currentPrevKnot.position.x == currentPrevKnot.next.position.x

			if knotIsRightOfKnot && knotIsAboveKnot {
				currentPrevKnot.next.position.x++
				currentPrevKnot.next.position.y++
			} else if knotIsRightOfKnot && knotIsBelowKnot {
				currentPrevKnot.next.position.x++
				currentPrevKnot.next.position.y--
			} else if knotIsLeftOfKnot && knotIsBelowKnot {
				currentPrevKnot.next.position.x--
				currentPrevKnot.next.position.y--
			} else if knotIsLeftOfKnot && knotIsAboveKnot {
				currentPrevKnot.next.position.x--
				currentPrevKnot.next.position.y++
			} else if knotIsRightOfKnot && knotIsSameRow {
				currentPrevKnot.next.position.x++
			} else if knotIsLeftOfKnot && knotIsSameRow {
				currentPrevKnot.position.x--
			} else if knotIsSameColumn && knotIsAboveKnot {
				currentPrevKnot.next.position.y++
			} else if knotIsSameColumn && knotIsBelowKnot {
				currentPrevKnot.next.position.y--
			}
			currentPrevKnot = *currentPrevKnot.next
			j++
		}
		i++
	}
	return knotH
}
