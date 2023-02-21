package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func NewScanner(filename string) (*bufio.Scanner, io.Closer) {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	return fileScanner, readFile
}
