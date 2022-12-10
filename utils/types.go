package utils

import (
	"log"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) Copy() Point {
	return Point{p.X, p.Y}
}

// ToInt converts a string to an integer, exiting the app if there is an error
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

// Within checks to see if a number is within the range of two given numbers
func Within(i int, min int, max int) bool {
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

// Min returns the minimum of the two given numbers
func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of the two given numbers
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
