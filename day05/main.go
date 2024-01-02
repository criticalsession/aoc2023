package main

import (
	"fmt"
	"strings"

	. "github.com/criticalsession/aoc2023/day05/data"
	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path: "sample.txt",
	})
	utils.Catch(err)

	//solve(s, false)
	solve(s, true)
}

func solve(s []string, partTwo bool) {
	d := parseInput(&s)
	minLoc := -1

	if !partTwo {
		for i := range d.Seeds {
			updateMinloc(&d.Seeds[i], &d, &minLoc)
		}
	} else {
		d.RangeBlocks[0].ConvertToOverlappingRanges(d.GetSeedRanges()[0].Start, d.GetSeedRanges()[0].End)
	}

	fmt.Println(minLoc)
}

func updateMinloc(seed *int, d *ParsedData, minLoc *int) {
	fin := *seed
	for _, rb := range d.RangeBlocks {
		fin = rb.GetDestValue(fin)
	}

	if *minLoc < 0 || fin < *minLoc {
		*minLoc = fin
	}
}

func parseInput(s *[]string) ParsedData {
	d := ParsedData{
		RangeBlocks: make([]RangeBlock, 7),
	}

	currVals := RangeBlock{}
	currBlock := -1

	for _, l := range *s {
		if strings.HasPrefix(l, "seeds:") {
			// extract seeds
			l = strings.Replace(l, "seeds: ", "", 1)
			parts := strings.Split(l, " ")
			for _, seed := range parts {
				if seed != "" {
					d.Seeds = append(d.Seeds, utils.ConvertToNumber(seed))
				}
			}
		} else if l == "" {
			// store and reset current range block
			if currBlock >= 0 {
				d.RangeBlocks[currBlock] = currVals
			}

			currVals = RangeBlock{}
		} else if utils.IsNumber(l[0]) {
			// append range to existing block
			parts := strings.Split(l, " ")
			currVals = append(currVals, ValRange{
				DestStart:   utils.ConvertToNumber(parts[0]),
				SourceStart: utils.ConvertToNumber(parts[1]),
				Width:       utils.ConvertToNumber(parts[2]),
			})
		} else {
			currBlock++
		}
	}

	if len(currVals) > 0 {
		d.RangeBlocks[currBlock] = currVals
	}

	return d
}
