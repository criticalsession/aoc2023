package main

import (
	"fmt"
	"strings"

	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path:  "day2input.txt",
		Split: ":",
	})
	if err != nil {
		panic(err)
	}

	solve(s)
	solvePart2(s)
}

func solve(s []string) {
	gameTotal := 0
	for gameNo, line := range s {
		setsOk := true
		for _, set := range utils.SplitAndTrim(line, ";") {
			col := getColorsFromSet(set)
			if col[Red] > 12 || col[Green] > 13 || col[Blue] > 14 {
				setsOk = false
			}
		}

		if setsOk {
			gameTotal += gameNo + 1
		}
	}

	fmt.Println("Total:", gameTotal)
}

func solvePart2(s []string) {
	gameTotal := 0
	for _, line := range s {
		var mins ColorMap = ColorMap{
			Red: 0, Green: 0, Blue: 0,
		}
		for _, set := range utils.SplitAndTrim(line, ";") {
			col := getColorsFromSet(set)
			mins[Red] = max(mins[Red], col[Red])
			mins[Blue] = max(mins[Blue], col[Blue])
			mins[Green] = max(mins[Green], col[Green])
		}

		gameTotal += mins.Power()
	}

	fmt.Println("Total:", gameTotal)
}

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type ColorMap map[Color]int

func (m *ColorMap) Power() int {
	return (*m)[Red] * (*m)[Blue] * (*m)[Green]
}

func getColorsFromSet(s string) ColorMap {
	colorsInSet := ColorMap{
		Red: 0, Green: 0, Blue: 0,
	}

	for _, c := range utils.SplitAndTrim(s, ",") {
		if strings.Contains(c, "blue") {
			colorsInSet[Blue] += utils.ReplaceAndGetInt(c, "blue")
		} else if strings.Contains(c, "red") {
			colorsInSet[Red] += utils.ReplaceAndGetInt(c, "red")
		} else if strings.Contains(c, "green") {
			colorsInSet[Green] += utils.ReplaceAndGetInt(c, "green")
		}
	}

	return colorsInSet
}
