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
