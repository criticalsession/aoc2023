package data

type ParsedData struct {
	Seeds       []int
	RangeBlocks []RangeBlock
}

func (p *ParsedData) GetSeedRanges() [][]int {
	ranges := [][]int{}
	seedStart := -1
	for _, s := range p.Seeds {
		if seedStart < 0 {
			seedStart = s
			continue
		} else {
			ranges = append(ranges, []int{seedStart, seedStart + s})
			seedStart = -1
		}
	}

	return ranges
}

type ValRange struct {
	DestStart   int
	SourceStart int
	Width       int
}

func (v *ValRange) GetDestValue(s int) int {
	if s < v.SourceStart || s > v.SourceStart+v.Width {
		// outside of range
		return -1
	}

	return v.DestStart + (s - v.SourceStart)
}

type RangeBlock []ValRange

func (r *RangeBlock) GetDestValue(s int) int {
	for _, v := range *r {
		if val := v.GetDestValue(s); val >= 0 {
			return val
		}
	}

	//sny source numbers that aren't mapped correspond to the same destination number.
	return s
}
