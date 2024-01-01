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

	solve(s, false)
}

func solve(s []string, partTwo bool) {
	d := parseInput(&s)
	minLoc := -1

	if !partTwo {
		for i := range d.Seeds {
			updateMinloc(&d.Seeds[i], &d, &minLoc)
		}
	} else {
		for _, sr := range d.GetSeedRanges() {
			for i := sr[0]; i < sr[1]; i++ {
				updateMinloc(&i, &d, &minLoc)
			}
		}
	}

	fmt.Println(minLoc)
}

func updateMinloc(seed *int, d *ParsedData, minLoc *int) {
	fin := d.Walk(*seed)
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
			ds, ss, w := utils.ConvertToNumber(parts[0]),
				utils.ConvertToNumber(parts[1]),
				utils.ConvertToNumber(parts[2])

			currVals = append(currVals, ValRange{
				DestStart:   ds,
				DestEnd:     ds + w,
				SourceStart: ss,
				SourceEnd:   ss + w,
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
