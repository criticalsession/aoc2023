package data

import "fmt"

type ParsedData struct {
	Seeds       []int
	RangeBlocks []RangeBlock
}

func (p *ParsedData) Walk(seed int) int {
	for i := range p.RangeBlocks {
		seed = p.RangeBlocks[i].GetDestValue(seed)
	}

	return seed
}

func (p *ParsedData) GetSeedRanges() [][]int {
	ranges := [][]int{}
	seedStart := -1
	for _, s := range p.Seeds {
		if seedStart < 0 {
			seedStart = s
			continue
		} else {
			ranges = append(ranges, []int{seedStart, seedStart + s - 1})
			seedStart = -1
		}
	}

	fmt.Println(ranges)

	return ranges
}

type ValRange struct {
	DestStart   int
	DestEnd     int
	SourceStart int
	SourceEnd   int
}

func (v *ValRange) GetDestValue(s int) int {
	if s < v.SourceStart || s > v.SourceEnd {
		// outside of range
		return -1
	}

	return v.DestStart + (s - v.SourceStart)
}

type RangeBlock []ValRange

func (r *RangeBlock) GetDestValue(s int) int {
	for i := range *r {
		if val := (*r)[i].GetDestValue(s); val >= 0 {
			return val
		}
	}

	// any source numbers that aren't mapped correspond to the same destination number.
	return s
}
