package main

import "github.com/criticalsession/aoc2023/utils"

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path: "sample.txt",
	})
	utils.Catch(err)

	solve(s, false)
}

func solve(s []string, partTwo bool) {
	// PART 1
	// 1. build set using left side
	// 2. for each number owned, check set
	// 3. total points = 2^(count-1)

}
