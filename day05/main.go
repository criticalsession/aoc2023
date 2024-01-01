package main

import (
	"fmt"
	"strings"

	"github.com/criticalsession/aoc2023/utils"
)

type parsedData struct {
	seeds       []int
	rangeBlocks []rangeBlock
}

type valRange struct {
	destStart   int
	sourceStart int
	width       int
}

type rangeBlock []valRange

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path: "sample.txt",
	})
	utils.Catch(err)

	solve(s, false)
}

func solve(s []string, partTwo bool) {
	d := parseInput(&s)
	fmt.Println(d)
}

func parseInput(s *[]string) parsedData {
	d := parsedData{
		rangeBlocks: make([]rangeBlock, 7),
	}

	currVals := rangeBlock{}
	currBlock := -1

	for _, l := range *s {
		if strings.HasPrefix(l, "seeds:") {
			// extract seeds
			l = strings.Replace(l, "seeds: ", "", 1)
			parts := strings.Split(l, " ")
			for _, seed := range parts {
				if seed != "" {
					d.seeds = append(d.seeds, utils.ConvertToNumber(seed))
				}
			}
		} else if l == "" {
			// store and reset current range block
			if currBlock >= 0 {
				d.rangeBlocks[currBlock] = currVals
			}

			currVals = rangeBlock{}
		} else if utils.IsNumber(l[0]) {
			// append range to existing block
			parts := strings.Split(l, " ")
			currVals = append(currVals, valRange{
				destStart:   utils.ConvertToNumber(parts[0]),
				sourceStart: utils.ConvertToNumber(parts[1]),
				width:       utils.ConvertToNumber(parts[2]),
			})
		} else {
			currBlock++
		}
	}

	if len(currVals) > 0 {
		d.rangeBlocks[currBlock] = currVals
	}

	return d
}
