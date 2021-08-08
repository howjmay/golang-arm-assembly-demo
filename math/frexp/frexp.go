// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package frexp

import (
	"math"
)

// Frexp breaks f into a normalized fraction
// and an integral power of two.
// It returns frac and exp satisfying f == frac × 2**exp,
// with the absolute value of frac in the interval [½, 1).
//
// Special cases are:
//	Frexp(±0) = ±0, 0
//	Frexp(±Inf) = ±Inf, 0
//	Frexp(NaN) = NaN, 0
func Frexp(f float64) (frac float64, exp int) {
	if haveArchFrexp {
		return ArchFrexp(f)
	}
	return frexp(f)
}

func normalize(x float64) (y float64, exp int) {
	const SmallestNormal = 2.2250738585072014e-308 // 2**-1022
	if math.Abs(x) < SmallestNormal {
		return x * (1 << 52), -52
	}
	return x, 0
}

var FrexpGo = frexp

func frexp(f float64) (frac float64, exp int) {
	// special cases
	mask := uint64(0x7FF)
	shift := uint64(64 - 11 - 1)
	bias := uint64(1023)
	switch {
	case f == 0:
		return f, 0 // correctly return -0
	case math.IsInf(f, 0) || math.IsNaN(f):
		return f, 0
	}
	f, exp = normalize(f)

	x := math.Float64bits(f)
	exp += int(((x >> shift) & mask) - bias + 1)
	x &^= mask << shift
	x |= uint64((-1 + int(bias)) << shift)
	frac = math.Float64frombits(x)
	return
}

func main() {
	_, _ = Frexp(1.0)
}
