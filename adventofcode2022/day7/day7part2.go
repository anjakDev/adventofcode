package main

import (
	"adventofcode/util"

	"golang.org/x/exp/slices"
)

const TOTAL_SIZE_FILESYSTEM = 70_000_000
const UNUSED_SPACE_REQ = 30_000_000

func parttwo() int {
	fileScanner, closer := util.NewScanner("input-day7.txt")
	defer closer.Close()

	input := make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		input = append(input, line)
	}
	cdCommand := "$ cd"

	rootDirectoryName := "/"
	rootDirectory := Directory{
		name:           rootDirectoryName,
		totalSize:      0,
		subdirectories: make([]*Directory, 0),
		parent:         nil,
	}

	rootDirectory, idx := getDirectContent(input, 2, rootDirectory)
	idx += 2
	activeDirectory := &rootDirectory
	for idx < len(input) {
		line := input[idx]
		if line[0:4] == cdCommand && line[5:] != ".." {
			dirName := line[5:]
			idxOfSubdirectory := slices.IndexFunc(activeDirectory.subdirectories, func(d *Directory) bool { return d.name == dirName })
			parentDirectory := activeDirectory
			activeDirectory = activeDirectory.subdirectories[idxOfSubdirectory]
			idx += 2
			tempDir, tempIdx := getDirectContent(input, idx, *activeDirectory)
			idx += tempIdx
			activeDirectory = &tempDir
			parentDirectory.subdirectories[idxOfSubdirectory] = activeDirectory
			continue
		}
		if line[0:4] == cdCommand && line[5:] == ".." {
			activeDirectory = activeDirectory.parent
			idx++
		}
	}
	resultSizes := fixSizes(&rootDirectory)
	rootDirectory.totalSize = resultSizes

	diffToTotal := TOTAL_SIZE_FILESYSTEM - resultSizes
	diffToUnused := UNUSED_SPACE_REQ - diffToTotal

	var getValidDirectories func(*Directory)

	valids := make([]int, 0)

	getValidDirectories = func(directory *Directory) {
		if directory.totalSize <= diffToUnused {
			return
		}
		valids = append(valids, directory.totalSize)
		for i := range directory.subdirectories {
			getValidDirectories(directory.subdirectories[i])
		}
	}

	getValidDirectories(&rootDirectory)

	min := valids[0]
	for _, v := range valids {
		if v < min {
			min = v
		}
	}
	return min
}
