package main

import (
	"fmt"

	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path: "sample.txt",
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
			} else {
				if nColl != "" {
					// process collected
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

					if hasAdjacentSymbols(line, prev,
						next, started, started+len(nColl),
						lineLen) {

						total += utils.ConvertToNumber(nColl)
						fmt.Println("Found:", nColl)
					}

					started = -1
					nColl = ""
				}
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
	symbols := []string{"+", "-", "*", "/", "$", "%", "#", "@", "&"}
	sameLine, _ := workCoords(start, end, lineLen)

	for _, sym := range symbols {
		for _, cSame := range sameLine {
			if string(l[cSame]) == sym {
				return true
			}
		}
	}

	return false
}

func workCoords(start, end, lineLen int) ([]int, []int) {
	sameLine := make([]int, 0, 2)
	adjLine := make([]int, 0, (end-start)+2)

	if start > 0 {
		sameLine = append(sameLine, start-1)
	}

	if end < lineLen {
		sameLine = append(sameLine, end)
	}

	return sameLine, adjLine
}
