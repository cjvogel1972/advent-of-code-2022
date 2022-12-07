package day7

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

type Puzzle struct{}

// Solve solves day 7's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day7/day7-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	rootDir := createRootDir(lines)

	return computeDirSize(rootDir)
}

func solvePart2(lines []string) int {
	rootDir := createRootDir(lines)
	totalDiskSpace := 70000000
	unusedDiskSpace := totalDiskSpace - rootDir.size()
	neededDiskSpace := 30000000
	spaceToClear := neededDiskSpace - unusedDiskSpace

	curDir := &rootDir
	smallestDir := findDirectoryToDelete(curDir, spaceToClear)

	return smallestDir.size()
}

func findDirectoryToDelete(dir *directory, spaceToClear int) *directory {
	var dirSizes = make(map[string]*directory)
	for _, d := range dir.directories {
		size := d.size()
		if size >= spaceToClear {
			result := findDirectoryToDelete(d, spaceToClear)
			dirSizes[result.name] = result
		}
	}

	if len(dirSizes) == 0 {
		return dir
	}

	var smallestName string
	smallestSize := 0
	for n, d := range dirSizes {
		if d.size() < smallestSize || smallestSize == 0 {
			smallestName = n
			smallestSize = d.size()
		}
	}

	return dirSizes[smallestName]
}

func createRootDir(lines []string) directory {
	rootDir := newDirectory("/", nil)
	curDir := &rootDir
	for i := 1; i < len(lines); {
		if lines[i] == "$ ls" {
			i++
			i = parseListing(curDir, lines, i)
		} else if strings.HasPrefix(lines[i], "$ cd") {
			entry := strings.Split(lines[i], " ")
			if entry[2] == ".." {
				if curDir.parent != nil {
					curDir = curDir.parent
				}
			} else {
				curDir = curDir.directories[entry[2]]
			}
			i++
		}
	}
	return rootDir
}

func parseListing(dir *directory, lines []string, i int) int {
	for i < len(lines) && !strings.HasPrefix(lines[i], "$") {
		entry := strings.Split(lines[i], " ")
		if entry[0] == "dir" {
			newDir := newDirectory(entry[1], dir)
			dir.directories[entry[1]] = &newDir
		} else {
			dir.files[entry[1]] = utils.ToInt(entry[0])
		}
		i++
	}
	return i
}

func computeDirSize(dir directory) int {
	size := 0
	dirSize := dir.size()
	if dirSize <= 100000 {
		size += dirSize
	}

	for _, d := range dir.directories {
		size += computeDirSize(*d)
	}

	return size
}

func newDirectory(name string, parent *directory) directory {
	return directory{name, make(map[string]int), make(map[string]*directory), parent}
}

type directory struct {
	name        string
	files       map[string]int
	directories map[string]*directory
	parent      *directory
}

func (d directory) size() int {
	size := 0
	for _, f := range d.files {
		size += f
	}
	for _, d := range d.directories {
		size += d.size()
	}

	return size
}
