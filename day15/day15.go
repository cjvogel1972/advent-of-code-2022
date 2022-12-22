package day15

import (
	"advent-of-code-2022/utils"
	_range "advent-of-code-2022/utils/range"
	"fmt"
)

type Puzzle struct{}

// Solve solves day 15's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day15/day15-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines, 2000000))
	fmt.Printf("Part 2: %d\n", solvePart2(lines, 4000000))
}

func solvePart1(lines []string, rowNum int) int {
	scanners, beacons := parseScannersBeacons(lines)

	row := make(map[utils.Point]bool)
	for _, s := range scanners {
		dist := utils.Abs(s.loc.Y - rowNum)
		for i := 0; i <= s.dist-dist; i++ {
			row[utils.Point{s.loc.X + i, rowNum}] = false
			row[utils.Point{s.loc.X - i, rowNum}] = false
		}
	}

	for p, _ := range beacons {
		if p.Y == rowNum {
			row[p] = true
		}
	}

	numNonBeaconsRow := 0
	for _, b := range row {
		if !b {
			numNonBeaconsRow++
		}
	}

	return numNonBeaconsRow
}

func solvePart2(lines []string, maxCoord int) int {
	scanners, _ := parseScannersBeacons(lines)

	rows := make([][]_range.Range, maxCoord+1)
	for _, s := range scanners {
		if s.loc.Y+s.dist < 0 || s.loc.Y-s.dist > maxCoord+1 {
			continue
		}
		for i := 0; i < s.dist+1; i++ {
			above := s.loc.Y - i
			below := s.loc.Y + i
			start := utils.Max(0, s.loc.X-s.dist+i)
			end := utils.Min(maxCoord, s.loc.X+s.dist-i)
			if utils.Within(below, 0, maxCoord) {
				processRow(rows, below, start, end)
			}
			if i != 0 && utils.Within(above, 0, maxCoord) {
				processRow(rows, above, start, end)
			}

		}
	}

	var distressBeacon utils.Point

	for y := 0; y < len(rows); y++ {
		if len(rows[y]) == 1 {
			continue
		}
		if rows[y][0].End < rows[y][1].Start {
			distressBeacon = utils.Point{rows[y][0].End + 1, y}
			break
		} else {
			distressBeacon = utils.Point{rows[y][1].End + 1, y}
			break
		}
	}

	return distressBeacon.X*4000000 + distressBeacon.Y
}

func processRow(rows [][]_range.Range, i int, start int, end int) {
	newRange := _range.New(start, end)
	if len(rows[i]) == 0 {
		rows[i] = make([]_range.Range, 0)
		rows[i] = append(rows[i], newRange)
	} else {
		within := false
		for j := 0; j < len(rows[i]); {
			r := rows[i][j]
			if r.Within(newRange) {
				within = true
				break
			} else if newRange.Within(r) {
				rows[i][j] = rows[i][len(rows[i])-1]
				rows[i] = rows[i][:len(rows[i])-1]
				continue
			} else if r.Overlap(newRange) || r.End+1 == newRange.Start || newRange.End+1 == r.Start {
				newRange = newRange.Merge(r)
				rows[i][j] = rows[i][len(rows[i])-1]
				rows[i] = rows[i][:len(rows[i])-1]
				continue
			}
			j++
		}
		if !within {
			rows[i] = append(rows[i], newRange)
		}
	}
}

func parseScannersBeacons(lines []string) ([]scanner, map[utils.Point]bool) {
	scanners := make([]scanner, 0)
	beacons := make(map[utils.Point]bool)
	for _, l := range lines {
		var sx int
		var sy int
		var bx int
		var by int

		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		s := utils.Point{sx, sy}
		b := utils.Point{bx, by}
		dist := s.ManhattanDistance(b)
		scanners = append(scanners, scanner{s, b, dist})
		beacons[b] = true
	}
	return scanners, beacons
}

type scanner struct {
	loc           utils.Point
	closestBeacon utils.Point
	dist          int
}

func (s scanner) inRange(p utils.Point) bool {
	return s.loc.ManhattanDistance(p) <= s.dist
}
