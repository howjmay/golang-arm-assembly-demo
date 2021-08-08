package hypot_test

import (
	"math"
	"testing"

	"github.com/HowJMay/golang-arm-assembly-demo/math/hypot"
)

var tanh = []float64{
	9.9990531206936338549262119e-01,
	9.9999962057085294197613294e-01,
	-2.7001505097318677233756845e-01,
	-9.9991110943061718603541401e-01,
	9.9999999146798465745022007e-01,
	9.9427249436125236705001048e-01,
	9.9994257600983138572705076e-01,
	9.9149409509772875982054701e-01,
	9.4936501296239685514466577e-01,
	-9.9999994291374030946055701e-01,
}
var vf = []float64{
	4.9790119248836735e+00,
	7.7388724745781045e+00,
	-2.7688005719200159e-01,
	-5.0106036182710749e+00,
	9.6362937071984173e+00,
	2.9263772392439646e+00,
	5.2290834314593066e+00,
	2.7279399104360102e+00,
	1.8253080916808550e+00,
	-8.6859247685756013e+00,
}

var vfhypotSC = [][2]float64{
	{math.Inf(-1), math.Inf(-1)},
	{math.Inf(-1), 0},
	{math.Inf(-1), math.Inf(1)},
	{math.Inf(-1), math.NaN()},
	{math.Copysign(0, -1), math.Copysign(0, -1)},
	{math.Copysign(0, -1), 0},
	{0, math.Copysign(0, -1)},
	{0, 0}, // +0, +0
	{0, math.Inf(-1)},
	{0, math.Inf(1)},
	{0, math.NaN()},
	{math.Inf(1), math.Inf(-1)},
	{math.Inf(1), 0},
	{math.Inf(1), math.Inf(1)},
	{math.Inf(1), math.NaN()},
	{math.NaN(), math.Inf(-1)},
	{math.NaN(), 0},
	{math.NaN(), math.Inf(1)},
	{math.NaN(), math.NaN()},
}
var hypotSC = []float64{
	math.Inf(1),
	math.Inf(1),
	math.Inf(1),
	math.Inf(1),
	0,
	0,
	0,
	0,
	math.Inf(1),
	math.Inf(1),
	math.NaN(),
	math.Inf(1),
	math.Inf(1),
	math.Inf(1),
	math.Inf(1),
	math.Inf(1),
	math.NaN(),
	math.Inf(1),
	math.NaN(),
}

func tolerance(a, b, e float64) bool {
	// Multiplying by e here can underflow denormal values to zero.
	// Check a==b so that at least if a and b are small and identical
	// we say they match.
	if a == b {
		return true
	}
	d := a - b
	if d < 0 {
		d = -d
	}

	// note: b is correct (expected) value, a is actual value.
	// make error tolerance a fraction of b, not a.
	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}
	return d < e
}
func close(a, b float64) bool      { return tolerance(a, b, 1e-14) }
func veryclose(a, b float64) bool  { return tolerance(a, b, 4e-16) }
func soclose(a, b, e float64) bool { return tolerance(a, b, e) }
func alike(a, b float64) bool {
	switch {
	case math.IsNaN(a) && math.IsNaN(b):
		return true
	case a == b:
		return math.Signbit(a) == math.Signbit(b)
	}
	return false
}
func TestHypot(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := math.Abs(1e200 * tanh[i] * math.Sqrt(2))
		if f := hypot.Hypot(1e200*tanh[i], 1e200*tanh[i]); !veryclose(a, f) {
			t.Errorf("Hypot(%g, %g) = %g, want %g", 1e200*tanh[i], 1e200*tanh[i], f, a)
		}
	}
	for i := 0; i < len(vfhypotSC); i++ {
		if f := hypot.Hypot(vfhypotSC[i][0], vfhypotSC[i][1]); !alike(hypotSC[i], f) {
			t.Errorf("Hypot(%g, %g) = %g, want %g", vfhypotSC[i][0], vfhypotSC[i][1], f, hypotSC[i])
		}
	}
}

var (
	GlobalI int
	GlobalB bool
	GlobalF float64
)

func BenchmarkHypot(b *testing.B) {
	x := 0.0
	for i := 0; i < b.N; i++ {
		x = hypot.Hypot(3, 4)
	}
	GlobalF = x
}

func BenchmarkHypotGo(b *testing.B) {
	x := 0.0
	for i := 0; i < b.N; i++ {
		x = hypot.HypotGo(3, 4)
	}
	GlobalF = x
}
