package main

import (
	"adventofcode/util"
	"bufio"
)

func partone() int {
	fileScanner, closer := util.NewScanner("input-day6.txt")
	fileScanner.Split(bufio.ScanBytes)
	defer closer.Close()

	var position int
	fourLetterSlice := make([]string, 0)
	var fourLetterMap map[string]int
	for fileScanner.Scan() {
		letter := fileScanner.Text()
		position++
		if len(fourLetterSlice) < 4 {
			fourLetterSlice = append(fourLetterSlice, letter)
		}
		if len(fourLetterSlice) == 4 {
			uniqueFour := true
			fourLetterMap = make(map[string]int)
			for i, elem := range fourLetterSlice {
				if _, ok := fourLetterMap[elem]; ok {
					delete(fourLetterMap, fourLetterSlice[0])
					fourLetterSlice = fourLetterSlice[1:]
					uniqueFour = false
					break
				}
				fourLetterMap[elem] = i
			}
			if uniqueFour {
				break
			}
		}
	}

	return position
}
