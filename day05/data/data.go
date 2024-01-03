package data

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
			ranges = append(ranges, StartEnd{seedStart, seedStart + s - 1})
			seedStart = -1
		}
	}

	return ranges
}

type ValRange struct {
	DestStart   int
	SourceStart int
	SourceEnd   int
	Offset      int
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

func (r *RangeBlock) FindOverlappingRanges(rngs []StartEnd) []StartEnd {
	results := []StartEnd{}
	newRanges := []StartEnd{}

	for _, rng := range rngs {
		found := false
		for _, v := range *r {
			if rng.End >= v.SourceStart && rng.Start < v.SourceEnd {
				// overlaps
				overlapStart := max(rng.Start, v.SourceStart)
				overlapEnd := min(rng.End, v.SourceEnd)
				results = append(results, StartEnd{overlapStart + v.Offset, overlapEnd + v.Offset})

				if overlapStart > rng.Start {
					newRanges = append(newRanges, StartEnd{rng.Start, overlapStart - 1})
				}

				if overlapEnd < rng.End {
					newRanges = append(newRanges, StartEnd{overlapEnd, rng.End})
				}

				found = true
				break
			}
		}

		if !found { // fallthrough
			results = append(results, rng)
		}
	}

	if len(newRanges) > 0 { // there are unprocessed ranges
		results = append(results, (*r).FindOverlappingRanges(newRanges)...)
	}

	return results
}
