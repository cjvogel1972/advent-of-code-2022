package utils

import (
	"log"
	"strconv"
)

// ConvertToInt converts a string to an integer, exiting the app if there is an error
func ConvertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

func InBetween(i int, min int, max int) bool {
	return min <= i && i <= max
}
