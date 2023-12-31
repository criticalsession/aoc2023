package main

import (
	"fmt"
	"strings"
	"time"

	. "github.com/criticalsession/aoc2023/day05/data"
	"github.com/criticalsession/aoc2023/utils"
)

func main() {
	s, err := utils.GetInput(utils.InputOptions{
		Path: "input.txt",
	})
	utils.Catch(err)

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
		t1 := time.Now()
		seedRanges := d.GetSeedRanges()

		chans := make([]chan []StartEnd, len(seedRanges))
		for i, sr := range seedRanges {
			chans[i] = make(chan []StartEnd)
			go processSeedRange(sr, d.RangeBlocks, chans[i])
		}

		for _, c := range chans {
			for _, r := range <-c {
				if minLoc < 0 || r.Start < minLoc {
					minLoc = r.Start
				}
			}
		}

		fmt.Println(time.Since(t1))
	}

	fmt.Println(minLoc)
}

func updateMinloc(seed *int, d *ParsedData, minLoc *int) {
	fin := d.Walk(*seed)
	if *minLoc < 0 || fin < *minLoc {
		*minLoc = fin
	}
}

func processSeedRange(sr StartEnd, rb []RangeBlock, c chan []StartEnd) {
	updatedSeedRange := []StartEnd{sr}
	for _, rb := range rb {
		updatedSeedRange = rb.FindOverlappingRanges(updatedSeedRange)
	}

	c <- updatedSeedRange
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
				SourceStart: ss,
				SourceEnd:   ss + w - 1,
				Offset:      ds - ss,
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
