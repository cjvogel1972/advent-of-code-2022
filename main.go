package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/utils"
	"fmt"
	"os"
)

type Puzzle interface {
	Solve()
}

var puzzles = []Puzzle{
	day1.Puzzle{},
}

func main() {
	if len(os.Args) > 1 {
		day := utils.ConvertToInt(os.Args[1]) - 1
		puzzles[day].Solve()
	} else {
		for i, puzzle := range puzzles {
			fmt.Printf("Day %d\n", i+1)
			fmt.Printf("---------------\n")
			puzzle.Solve()
			fmt.Printf("\n")
		}
	}

}
