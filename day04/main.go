package main

import (
	"fmt"
	"math"

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
	// PART 1
	// 1. build set using left side
	// 2. for each number owned, check set
	// 3. total points = 2^(count-1)

	total := 0
	for _, l := range s {
		card := utils.SplitAndTrim(l, "|")
		win := utils.SplitAndTrim(card[0], " ")
		nums := utils.SplitAndTrim(card[1], " ")

		test := make(map[int]bool, 5)
		for _, n := range win {
			if n != "" {
				test[utils.ConvertToNumber(n)] = true
			}
		}

		var c int8 = 0
		for _, n := range nums {
			if n != "" && test[utils.ConvertToNumber(n)] {
				c++ // hehe
			}
		}

		if c > 0 {
			total += int(math.Pow(2, float64(c-1)))
		}
	}

	fmt.Println(total)
}
