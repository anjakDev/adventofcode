package main

import (
	"adventofcode/util"
	"strings"
)

func parttwo() int {
	fileScanner, closer := util.NewScanner("input-day3.txt")
	defer closer.Close()

	var sumOfPriorities int

	lines := make([]string, 0)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	for i := 0; i < len(lines); i += 3 {
		letterMap := make(map[rune]int)
		for _, letter := range lines[i] {
			letterMap[letter] = 1
		}
		letterCollection := ""
		for _, letter := range lines[i+1] {
			if strings.Contains(letterCollection, string(letter)) {
				continue
			}
			if letterMap[letter] == 1 {
				letterMap[letter] += 1
				letterCollection += string(letter)
			}
		}
		letterCollection = ""
		for _, letter := range lines[i+2] {
			if strings.Contains(letterCollection, string(letter)) {
				continue
			}
			if letterMap[letter] == 2 {
				sumOfPriorities += getPriority(letter)
				letterCollection += string(letter)
			}
		}
	}
	return sumOfPriorities
}
