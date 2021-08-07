// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hypot

import (
	"math"
)

/*
	Hypot -- sqrt(p*p + q*q), but overflows only if the result does.
*/

// Hypot returns Sqrt(p*p + q*q), taking care to avoid
// unnecessary overflow and underflow.
//
// Special cases are:
//	Hypot(±Inf, q) = +Inf
//	Hypot(p, ±Inf) = +Inf
//	Hypot(NaN, q) = NaN
//	Hypot(p, NaN) = NaN
func Hypot(p, q float64) float64 {
	if haveArchHypot {
		return archHypot(p, q)
	}
	return hypot(p, q)
}

var HypotGo = hypot

func hypot(p, q float64) float64 {
	// special cases
	switch {
	case math.IsInf(p, 0) || math.IsInf(q, 0):
		return math.Inf(1)
	case math.IsNaN(p) || math.IsNaN(q):
		return math.NaN()
	}
	p, q = math.Abs(p), math.Abs(q)
	if p < q {
		p, q = q, p
	}
	if p == 0 {
		return 0
	}
	q = q / p
	return p * math.Sqrt(1+q*q)
}
