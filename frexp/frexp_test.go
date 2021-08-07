package frexp_test

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

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

var vffrexpSC = []float64{
	math.Inf(-1),
	math.Copysign(0, -1),
	0,
	math.Inf(1),
	math.NaN(),
}

type fi struct {
	f float64
	i int
}

var frexpSC = []fi{
	{math.Inf(-1), 0},
	{math.Copysign(0, -1), 0},
	{0, 0},
	{math.Inf(1), 0},
	{math.NaN(), 0},
}
var frexp = []fi{
	{6.2237649061045918750e-01, 3},
	{9.6735905932226306250e-01, 3},
	{-5.5376011438400318000e-01, -1},
	{-6.2632545228388436250e-01, 3},
	{6.02268356699901081250e-01, 4},
	{7.3159430981099115000e-01, 2},
	{6.5363542893241332500e-01, 3},
	{6.8198497760900255000e-01, 2},
	{9.1265404584042750000e-01, 1},
	{-5.4287029803597508250e-01, 4},
}

const (
	MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
)

const (
	SmallestNormalFloat64   = 2.2250738585072014e-308 // 2**-1022
	LargestSubnormalFloat64 = SmallestNormalFloat64 - SmallestNonzeroFloat64
)

var vffrexpBC = []float64{
	SmallestNormalFloat64,
	LargestSubnormalFloat64,
	SmallestNonzeroFloat64,
	MaxFloat64,
	-SmallestNormalFloat64,
	-LargestSubnormalFloat64,
	-SmallestNonzeroFloat64,
	-MaxFloat64,
}
var frexpBC = []fi{
	{0.5, -1021},
	{0.99999999999999978, -1022},
	{0.5, -1073},
	{0.99999999999999989, 1024},
	{-0.5, -1021},
	{-0.99999999999999978, -1022},
	{-0.5, -1073},
	{-0.99999999999999989, 1024},
}

func TestFrexp(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f, j := frexp.Frexp(vf[i]); !veryclose(frexp[i].f, f) || frexp[i].i != j {
			t.Errorf("Frexp(%g) = %g, %d, want %g, %d", vf[i], f, j, frexp[i].f, frexp[i].i)
		}
	}
	for i := 0; i < len(vffrexpSC); i++ {
		if f, j := frexp.Frexp(vffrexpSC[i]); !alike(frexpSC[i].f, f) || frexpSC[i].i != j {
			t.Errorf("Frexp(%g) = %g, %d, want %g, %d", vffrexpSC[i], f, j, frexpSC[i].f, frexpSC[i].i)
		}
	}
	for i := 0; i < len(vffrexpBC); i++ {
		if f, j := frexp.Frexp(vffrexpBC[i]); !veryclose(frexpBC[i].f, f) || frexpBC[i].i != j {
			t.Errorf("Frexp(%g) = %g, %d, want %g, %d", vffrexpBC[i], f, j, frexpBC[i].f, frexpBC[i].i)
		}
	}
}

func TestFrexpSelf(t *testing.T) {
	type test struct {
		f    float64
		frac float64
		exp  int
	}

	tests := []test{
		{
			f:    0,
			frac: 0,
			exp:  0,
		},
		{
			f:    math.Inf(0),
			frac: math.Inf(0),
			exp:  0,
		},
		{
			f:    math.NaN(),
			frac: math.NaN(),
			exp:  0,
		},
	}

	for _, tt := range tests {
		frac, exp := frexp.Frexp(tt.f)
		if math.IsNaN(tt.frac) {
			if !math.IsNaN(frac) || exp != 0 {
				t.Errorf("frac = %f, exp = %d, tt.frac = %f, tt.exp = %d\n", frac, exp, tt.frac, tt.exp)
			}
		} else if math.IsInf(tt.frac, 0) {
			if !math.IsInf(frac, 0) || exp != 0 {
				t.Errorf("frac = %f, exp = %d, tt.frac = %f, tt.exp = %d\n", frac, exp, tt.frac, tt.exp)
			}
		} else if frac != tt.frac || exp != tt.exp {
			t.Errorf("frac = %f, exp = %d, tt.frac = %f, tt.exp = %d\n", frac, exp, tt.frac, tt.exp)
		}
	}

	for i := 0; i < len(vffrexpBC); i++ {
		if f, j := frexp.ArchFrexp(vffrexpBC[i]); !veryclose(frexpBC[i].f, f) || frexpBC[i].i != j {
			_f, _j := frexp.frexp(vffrexpBC[i])
			t.Errorf("frexp(%g) = %g, %d, want %g, %d", vffrexpBC[i], _f, _j, frexpBC[i].f, frexpBC[i].i)
			t.Errorf("ArchFrexp(%g) = %g, %d, want %g, %d", vffrexpBC[i], f, j, frexpBC[i].f, frexpBC[i].i)
		}
	}

	// random test
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		randF := rand.Float64()
		randI := rand.Int31n(32)
		af, aexp := frexp.Frexp(randF * float64(randI))
		mf, mexp := math.Frexp(randF * float64(randI))
		if af != mf || aexp != mexp {
			t.Errorf("af = %f, aexp = %d, mf = %f, mexp  = %d\n", af, aexp, mf, mexp)
		}
	}
}

var (
	GlobalI int
	GlobalB bool
	GlobalF float64
)

func BenchmarkFrexp(b *testing.B) {
	x := 0.0
	y := 0
	for i := 0; i < b.N; i++ {
		x, y = frexp.ArchFrexp(8)
	}
	GlobalF = x
	GlobalI = y
}

func BenchmarkFrexpGO(b *testing.B) {
	x := 0.0
	y := 0
	for i := 0; i < b.N; i++ {
		x, y = frexp.frexp(8)
	}
	GlobalF = x
	GlobalI = y
}
