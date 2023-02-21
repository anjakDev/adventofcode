package main

import (
	"adventofcode/util"
	"strings"
)

func partone() int {
	fileScanner, closer := util.NewScanner("input-day3.txt")
	defer closer.Close()

	var sumOfPriorities int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		firstPart := line[0 : len(line)/2]
		secondPart := line[len(line)/2:]

		letterMap := make(map[rune]int)
		for _, letter := range firstPart {
			letterMap[letter] = 1
		}
		letterCollection := ""
		for _, letter := range secondPart {
			if strings.Contains(letterCollection, string(letter)) {
				continue
			}
			if letterMap[letter] == 1 {
				sumOfPriorities += getPriority(letter)
				letterCollection += string(letter)
			}
		}
	}

	return sumOfPriorities
}

func getPriority(itemType rune) int {
	if itemType < 91 {
		return int(itemType) - 38
	}
	return int(itemType) - 96
}
