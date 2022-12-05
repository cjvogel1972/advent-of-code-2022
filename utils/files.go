package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadLines reads lines from the given file
func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var values []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	return values
}

// SplitFile groups the lines of a file into two, based on a blank line
func SplitFile(lines []string) ([]string, []string) {
	var top = make([]string, 0)
	linesRead := 0
	for i := 0; len(lines[i]) > 0; i++ {
		top = append(top, lines[i])
		linesRead++
	}

	// take blank line into account
	linesRead++

	var bottom = make([]string, len(lines)-linesRead)
	for i := 0; i < len(lines)-linesRead; i++ {
		bottom[i] = lines[linesRead+i]
	}

	return top, bottom
}
