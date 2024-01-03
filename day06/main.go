package main

import (
	"fmt"
	"strings"

	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path:  "input.txt",
		Split: ":",
	})
	utils.Catch(err)

	solve2(s)
}

func solve(s []string) {
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

func solve2(s []string) {
	time := utils.ConvertToNumber(strings.ReplaceAll(s[0], " ", ""))
	distance := utils.ConvertToNumber(strings.ReplaceAll(s[1], " ", ""))

	min := -1
	max := -1
	i := 1

	for {
		if min > 0 && max > 0 {
			break
		}

		if min < 0 {
			if i*(time-i) > distance {
				min = i
			}
		}

		if max < 0 {
			j := time - i
			if j*(time-j) > distance {
				max = j
			}
		}

		i++
	}

	fmt.Println(max - min + 1)
}
