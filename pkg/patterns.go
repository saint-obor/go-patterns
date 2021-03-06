package patterns

// Pattern ...
type Pattern struct {
	Rhythm []uint8
}

// NewEuclid ...
func NewEuclid(n, k, rotation int32, groove float64) (*Pattern, error) {

	p := new(Pattern)

	// flip n and k if n is greater than k
	if n > k {
		n, k = k, n
	}

	p.createEuclidPattern(n, k)

	if groove != 0 {
		p.setGroove(n, k, groove)
	}

	if rotation != 0 {
		p.setRotate(int(rotation))
	}

	return p, nil
}

// createPattern creates a new rhythmic pattern using Bresenham’s line algorithm
func (p *Pattern) createEuclidPattern(n, k int32) {

	p.Rhythm = []uint8{}

	previous := -1

	ratio := float64(n) / float64(k)

	var i int32
	for i < k {
		x := int(ratio * float64(i))
		if x != previous {
			p.Rhythm = append(p.Rhythm, 1)
		} else {
			p.Rhythm = append(p.Rhythm, 0)
		}
		previous = x
		i++
	}
}

func (p *Pattern) setRotate(rotation int) {

	np := []uint8{}

	offset := 0

	// If offset is negative
	if rotation < 0 {
		// subtract from length,
		// the positive of offset,
		// constrained to length via mod
		offset = len(p.Rhythm) - ((rotation * -1) % len(p.Rhythm))
	} else {
		offset = rotation % len(p.Rhythm)
	}

	np = append(np, p.Rhythm[offset:]...)
	np = append(np, p.Rhythm[:offset]...)

	p.Rhythm = np
}

func (p *Pattern) setGroove(n, k int32, groove float64) {

	tmpRhythm := make([]uint8, len(p.Rhythm))
	copy(tmpRhythm, p.Rhythm)

	gn := (int32(reRange(groove, 0, 100, 0, float64(k))) + n)
	if gn > k {
		gn = k
	}

	groovePattern := new(Pattern)
	groovePattern.createEuclidPattern(gn, k)

	// fmt.Println("groovePattern:", groovePattern)

	midPointIndex := int(k / 2)

	i := 1 // skip the first beat
	for i < len(p.Rhythm) {

		if p.Rhythm[i] == 1 && i != midPointIndex {

			// fmt.Println("pattern index:", i)

			tmpRhythm[i] = 0

			groovePatternIndex := i
			distance := 1
			direction := 1 // positive is forward, negative backwards

			found := false

			for !found {

				groovePatternIndex = (groovePatternIndex + (distance * direction)) % len(groovePattern.Rhythm)

				if groovePatternIndex < 0 {
					tmpRhythm[i] = 1
					break
				}

				// fmt.Println("groove pattern index:", groovePatternIndex)

				if groovePattern.Rhythm[groovePatternIndex] == 1 && tmpRhythm[groovePatternIndex] != 1 {
					tmpRhythm[groovePatternIndex] = 1
					found = true
				}

				distance++
				direction *= -1
			}
		}

		i++
	}

	p.Rhythm = tmpRhythm

	// fmt.Println("pattern:      ", p)
}

// reRange maps a value from one range to another
func reRange(value, aMin, aMax, bMin, bMax float64) float64 {

	if value < aMin {
		value = aMin
	}

	if value > aMax {
		value = aMax
	}

	newValue := (((value - aMin) * (bMax - bMin)) / (aMax - aMin)) + bMin
	return newValue
}
