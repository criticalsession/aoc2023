package data

type ParsedData struct {
	Seeds       []int
	RangeBlocks []RangeBlock
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
