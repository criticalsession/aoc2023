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
	solve(s, true)
}

func solve(s []string, partTwo bool) {
	// PART 1
	// 1. build set using left side
	// 2. for each number owned, check set
	// 3. total points = 2^(count-1)

	total, total2 := 0, 0
	totalCards := len(s)
	scratchCardCount := make([]int, totalCards)

	for i, l := range s {
		scratchCardCount[i]++

		card := utils.SplitAndTrim(l, "|")
		win := utils.SplitAndTrim(card[0], " ")
		nums := utils.SplitAndTrim(card[1], " ")

		test := make(map[int]bool, 5)
		for _, n := range win {
			if n != "" {
				test[utils.ConvertToNumber(n)] = true
			}
		}

		winnings := 0
		for _, n := range nums {
			if n != "" && test[utils.ConvertToNumber(n)] {
				winnings++ // hehe
			}
		}

		if winnings > 0 {
			total += int(math.Pow(2, float64(winnings-1)))
		}

		if partTwo {
			// PART 2
			// 1. multiply winnings by number of copies of card
			// 2. increase number of following cards by winnings
			// 3. count this card and its copies towards total2
			for j := 0; j < winnings; j++ {
				if i+j+1 < totalCards {
					scratchCardCount[i+j+1] += scratchCardCount[i]
				}
			}

			total2 += scratchCardCount[i]
		}
	}

	if !partTwo {
		fmt.Println("Part 1: ", total)
	} else {
		fmt.Println("Part 2:", total2)
	}

}
