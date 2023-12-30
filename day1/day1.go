package main

import (
	"fmt"
	"strings"

	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path: "day1input.txt",
	})
	utils.Catch(err)

	solve(s, false)
	solve(s, true)
}

func solve(s []string, part2 bool) {
	total := 0
	for _, l := range s {
		lineLen := len(l)
		var j, k int = -1, -1

		for i := 0; i < lineLen; i++ {
			x := l[i]
			y := l[lineLen-i-1]

			if j < 0 {
				if x >= '0' && x <= '9' {
					j = int(x - '0')
				} else {
					j = textToNum(l[i:], UsePrefix)
				}
			}

			if k < 0 {
				if y >= '0' && y <= '9' {
					k = int(y - '0')
				} else {
					k = textToNum(l[:len(l)-i], UseSuffix)
				}
			}

			if j > -1 && k > -1 {
				break
			}
		}

		total += j*10 + k
	}
	fmt.Println("total:", total)
}

type PrefixCheck int

const (
	UsePrefix PrefixCheck = iota
	UseSuffix
)

func textToNum(l string, usePrefix PrefixCheck) int {
	if usePrefix == UsePrefix {
		if strings.HasPrefix(l, "zero") {
			return 0
		} else if strings.HasPrefix(l, "one") {
			return 1
		} else if strings.HasPrefix(l, "two") {
			return 2
		} else if strings.HasPrefix(l, "three") {
			return 3
		} else if strings.HasPrefix(l, "four") {
			return 4
		} else if strings.HasPrefix(l, "five") {
			return 5
		} else if strings.HasPrefix(l, "six") {
			return 6
		} else if strings.HasPrefix(l, "seven") {
			return 7
		} else if strings.HasPrefix(l, "eight") {
			return 8
		} else if strings.HasPrefix(l, "nine") {
			return 9
		}
	} else {
		if strings.HasSuffix(l, "zero") {
			return 0
		} else if strings.HasSuffix(l, "one") {
			return 1
		} else if strings.HasSuffix(l, "two") {
			return 2
		} else if strings.HasSuffix(l, "three") {
			return 3
		} else if strings.HasSuffix(l, "four") {
			return 4
		} else if strings.HasSuffix(l, "five") {
			return 5
		} else if strings.HasSuffix(l, "six") {
			return 6
		} else if strings.HasSuffix(l, "seven") {
			return 7
		} else if strings.HasSuffix(l, "eight") {
			return 8
		} else if strings.HasSuffix(l, "nine") {
			return 9
		}
	}

	return -1
}
