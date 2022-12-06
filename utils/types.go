package utils

import (
	"log"
	"strconv"
	"strings"
)

// ToInt converts a string to an integer, exiting the app if there is an error
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

// InBetween checks to see if a number is between two given numbers
func InBetween(i int, min int, max int) bool {
	return min <= i && i <= max
}

// HasDuplicates checks if a string has duplicate characters
func HasDuplicates(s string) bool {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) != 1 {
			return true
		}
	}
	return false
}
