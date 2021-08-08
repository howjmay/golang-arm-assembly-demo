package abs

import "math"

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs(x float64) float64 {
	if haveArchAbs {
		return ArchAbs(x)
	}
	return abs(x)
}

var AbsGo = abs

func abs(x float64) float64 {
	return math.Float64frombits(math.Float64bits(x) &^ (1 << 63))
}
