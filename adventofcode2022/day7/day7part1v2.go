package main

import (
	"adventofcode/util"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const limit = 100_000

type Directory struct {
	name           string
	totalSize      int
	subdirectories []*Directory
	parent         *Directory
}

func partoneV2() int {
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
	sumOfSizeLessThanOneK := getSumEQLTLimit(&rootDirectory)
	return sumOfSizeLessThanOneK
}

func getDirectContent(input []string, idx int, directory Directory) (Directory, int) {
	n := 0
	for n+idx < len(input) && input[idx+n][0:1] != "$" {
		line := input[idx+n]
		if line[0:3] == "dir" {
			subDirectory := Directory{
				name:           line[4:],
				totalSize:      0,
				subdirectories: make([]*Directory, 0),
				parent:         &directory,
			}
			directory.subdirectories = append(directory.subdirectories, &subDirectory)
			n++
			continue
		}
		fileDetails := strings.Split(line, " ")
		fileSize, _ := strconv.Atoi(fileDetails[0])
		directory.totalSize += fileSize
		n++
	}
	return directory, n
}

func fixSizes(directory *Directory) int {
	if len(directory.subdirectories) == 0 {
		return directory.totalSize
	}
	for i := range directory.subdirectories {
		sizesResult := fixSizes(directory.subdirectories[i])
		directory.totalSize += sizesResult
	}
	return directory.totalSize
}

func getSumEQLTLimit(directory *Directory) int {
	if len(directory.subdirectories) == 0 {
		if directory.totalSize <= limit {
			return directory.totalSize
		}
		return 0
	}
	sumLimitedSizesOfSubdirectories := 0
	for i := range directory.subdirectories {
		limitedSize := getSumEQLTLimit(directory.subdirectories[i])
		sumLimitedSizesOfSubdirectories += limitedSize
	}
	if directory.totalSize <= limit {
		return directory.totalSize + sumLimitedSizesOfSubdirectories
	}
	return sumLimitedSizesOfSubdirectories
}
