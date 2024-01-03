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

func (p *ParsedData) Walk(seed int) int {
	for i := range p.RangeBlocks {
		seed = p.RangeBlocks[i].GetDestValue(seed)
	}

	return seed
}

type StartEnd struct {
	Start int
	End   int
}

func (p *ParsedData) GetSeedRanges() []StartEnd {
	ranges := []StartEnd{}
	seedStart := -1
	for _, s := range p.Seeds {
		if seedStart < 0 {
			seedStart = s
			continue
		} else {
			ranges = append(ranges, StartEnd{seedStart, seedStart + s})
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

func (r *RangeBlock) ConvertToOverlappingRanges(s int, e int) {
	fmt.Println(s, e)

	results := []StartEnd{}
	for i := range *r {
		sourceStart := (*r)[i].SourceStart
		sourceEnd := (*r)[i].SourceStart + (*r)[i].Width
		fmt.Println("checking:", sourceStart, sourceEnd)

		if e >= sourceStart && s < sourceEnd {
			fmt.Println("overlap")

			// overlaps
			overlapStart := max(s, sourceStart)
			overlapEnd := min(e, sourceEnd)
			results = append(results, StartEnd{
				overlapStart - sourceStart + (*r)[i].DestStart,
				overlapEnd - sourceStart + (*r)[i].DestStart,
			})
		}
	}

	fmt.Println(results)
}
