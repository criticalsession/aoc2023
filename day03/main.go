package main

import (
	"fmt"
	"strings"

	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path: "input.txt",
	})
	utils.Catch(err)

	solve(s, false)
}

func solve(s []string, part2 bool) {
	total := 0
	lineLen := 0

	for i, line := range s {
		if i == 0 {
			lineLen = len(line)
		}
		var prev, next string
		nColl := ""
		started := -1
		for j, n := range line {
			if utils.IsNumber(byte(n)) {
				// collect
				if nColl == "" {
					started = j
				}
				nColl += string(n)

				if prev == "" && next == "" {
					if i > 0 {
						prev = s[i-1]
					} else {
						prev = ""
					}

					if i < len(s)-1 {
						next = s[i+1]
					} else {
						next = ""
					}
				}
			} else {
				if nColl != "" {
					// process collected
					if hasAdjacentSymbols(line, prev,
						next, started, started+len(nColl),
						lineLen) {

						total += utils.ConvertToNumber(nColl)
					}

					started = -1
					nColl = ""
				}
			}
		}

		if nColl != "" {
			// process collected
			if hasAdjacentSymbols(line, prev,
				next, started, started+len(nColl),
				lineLen) {

				total += utils.ConvertToNumber(nColl)
			}
		}
	}

	if part2 {
		fmt.Println("Total p2:", total)
	} else {
		fmt.Println("Total p1:", total)
	}
}

func hasAdjacentSymbols(l, prev, next string, start, end, lineLen int) bool {
	symbols := []string{"+", "-", "*", "/", "$", "%", "#", "@", "&", "="}
	coords := workCoords(start, end, lineLen)

	for _, sym := range symbols {
		if strings.Contains(l[coords[0]:coords[1]], sym) {
			return true
		}

		if prev != "" && strings.Contains(prev[coords[0]:coords[1]], sym) {
			return true
		}

		if next != "" && strings.Contains(next[coords[0]:coords[1]], sym) {
			return true
		}
	}

	return false
}

func workCoords(start, end, lineLen int) [2]int {
	coords := [2]int{start, end}

	if start > 0 {
		coords[0] = start - 1
	}

	if end < lineLen {
		coords[1] = end + 1
	}

	return coords
}
