package day7

import (
	"advent-of-code-2022/utils"
	"fmt"
	"math"
	"strings"
)

type Puzzle struct{}

const deviceTotalDiskSpace = 70000000
const spaceNeededForUpdate = 30000000
const maxDirectorySize = 100000

// Solve solves day 7's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day7/day7-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	rootDir := createDirectoryStructure(lines)

	return sumDirectorySizesLessThan(rootDir, maxDirectorySize)
}

func solvePart2(lines []string) int {
	rootDir := createDirectoryStructure(lines)

	unusedDiskSpace := deviceTotalDiskSpace - rootDir.size()
	spaceToClear := spaceNeededForUpdate - unusedDiskSpace

	toDelete := findDirectoryToDelete(rootDir, spaceToClear)

	return toDelete.size()
}

func createDirectoryStructure(lines []string) directory {
	rootDir := newDirectory("/", nil)
	curDir := &rootDir

	for i := 0; i < len(lines); {
		if lines[i] == "$ ls" {
			i = parseListing(curDir, lines, i+1)
		} else if strings.HasPrefix(lines[i], "$ cd") {
			dirName := strings.Split(lines[i], " ")[2]
			if dirName == "/" {
				curDir = &rootDir
			} else {
				curDir = curDir.changeDirectory(dirName)
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
			dir.makeDirectory(entry[1])
		} else {
			dir.addFile(entry[1], utils.ToInt(entry[0]))
		}
		i++
	}

	return i
}

func sumDirectorySizesLessThan(dir directory, max int) int {
	size := 0

	subDirs := dir.allSubDirectories()
	for _, d := range subDirs {
		dirSize := d.size()
		if dirSize <= max {
			size += dirSize
		}
	}

	return size
}

func findDirectoryToDelete(dir directory, spaceToClear int) directory {
	var toDelete directory
	smallestSize := math.MaxInt

	subDirs := dir.allSubDirectories()
	for _, s := range subDirs {
		size := s.size()
		if size > spaceToClear && size < smallestSize {
			toDelete = s
			smallestSize = size
		}
	}

	return toDelete
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

func (d *directory) addFile(name string, size int) {
	d.files[name] = size
}

func (d *directory) makeDirectory(name string) {
	dir := newDirectory(name, d)
	d.directories[name] = &dir
}

func (d *directory) changeDirectory(name string) *directory {
	if name == ".." {
		if d.parent == nil {
			return d
		} else {
			return d.parent
		}
	} else {
		cd, found := d.directories[name]
		if found {
			return cd
		} else {
			return d
		}
	}
}

func (d *directory) size() int {
	size := 0
	for _, f := range d.files {
		size += f
	}
	for _, d := range d.directories {
		size += d.size()
	}

	return size
}

func (d *directory) allSubDirectories() []directory {
	result := make([]directory, 0)

	for _, s := range d.directories {
		result = append(result, *s)
		result = append(result, s.allSubDirectories()...)
	}

	return result
}
