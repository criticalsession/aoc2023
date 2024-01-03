package main

import (
	"fmt"

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
	time := utils.SplitAndTrim(s[0], " ")
	distance := utils.SplitAndTrim(s[1], " ")

	result := 0
	for i := range time {
		solutions := 0

		t := utils.ConvertToNumber(time[i])
		d := utils.ConvertToNumber(distance[i])
		for j := 1; j < t; j++ {
			if j*(t-j) > d {
				solutions++
			}
		}

		if result == 0 {
			result = solutions
		} else {
			result *= solutions
		}
	}

	fmt.Println(result)
}
