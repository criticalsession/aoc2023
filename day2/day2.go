package main

import (
	"fmt"

	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path:  "day2sample.txt",
		Split: ":",
	})
	if err != nil {
		panic(err)
	}

	solve(s, false)
}

func solve(s []string, part1 bool) {
	gameIdTotal := 0
	for gameNo, line := range s {
		fmt.Println(line)
		gameIdTotal += gameNo
	}

	fmt.Println("Total:", gameIdTotal)
}
