package main

import (
	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path:  "input.txt",
		Split: ":",
	})
	utils.Catch(err)

	solve(s, false)
}

func solve(s []string, partTwo bool) {
}
