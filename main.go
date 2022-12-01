package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day11"
	"advent-of-code-2022/day12"
	"advent-of-code-2022/day13"
	"advent-of-code-2022/day14"
	"advent-of-code-2022/day15"
	"advent-of-code-2022/day16"
	"advent-of-code-2022/day17"
	"advent-of-code-2022/day18"
	"advent-of-code-2022/day19"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day20"
	"advent-of-code-2022/day21"
	"advent-of-code-2022/day22"
	"advent-of-code-2022/day23"
	"advent-of-code-2022/day24"
	"advent-of-code-2022/day25"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
	"advent-of-code-2022/utils"
	"fmt"
	"os"
)

type Puzzle interface {
	Solve()
}

var puzzles = []Puzzle{
	day1.Puzzle{},
	day2.Puzzle{},
	day3.Puzzle{},
	day4.Puzzle{},
	day5.Puzzle{},
	day6.Puzzle{},
	day7.Puzzle{},
	day8.Puzzle{},
	day9.Puzzle{},
	day10.Puzzle{},
	day11.Puzzle{},
	day12.Puzzle{},
	day13.Puzzle{},
	day14.Puzzle{},
	day15.Puzzle{},
	day16.Puzzle{},
	day17.Puzzle{},
	day18.Puzzle{},
	day19.Puzzle{},
	day20.Puzzle{},
	day21.Puzzle{},
	day22.Puzzle{},
	day23.Puzzle{},
	day24.Puzzle{},
	day25.Puzzle{},
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
