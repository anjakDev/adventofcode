package main

import (
	"adventofcode/util"
	"bufio"
)

func parttwo() int {
	fileScanner, closer := util.NewScanner("input-day6.txt")
	fileScanner.Split(bufio.ScanBytes)
	defer closer.Close()

	var position int
	fourteenLetterSlice := make([]string, 0)
	var fourteenLetterMap map[string]int
	for fileScanner.Scan() {
		letter := fileScanner.Text()
		position++
		if len(fourteenLetterSlice) < 14 {
			fourteenLetterSlice = append(fourteenLetterSlice, letter)
		}
		if len(fourteenLetterSlice) == 14 {
			uniqueFour := true
			fourteenLetterMap = make(map[string]int)
			for i, elem := range fourteenLetterSlice {
				if _, ok := fourteenLetterMap[elem]; ok {
					delete(fourteenLetterMap, fourteenLetterSlice[0])
					fourteenLetterSlice = fourteenLetterSlice[1:]
					uniqueFour = false
					break
				}
				fourteenLetterMap[elem] = i
			}
			if uniqueFour {
				break
			}
		}
	}

	return position
}
