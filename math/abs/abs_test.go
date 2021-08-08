package abs_test

import (
	"math"
	"testing"

	"github.com/HowJMay/golang-arm-assembly-demo/math/abs"
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

var fabs = []float64{
	4.9790119248836735e+00,
	7.7388724745781045e+00,
	2.7688005719200159e-01,
	5.0106036182710749e+00,
	9.6362937071984173e+00,
	2.9263772392439646e+00,
	5.2290834314593066e+00,
	2.7279399104360102e+00,
	1.8253080916808550e+00,
	8.6859247685756013e+00,
}
var vffabsSC = []float64{
	math.Inf(-1),
	math.Copysign(0, -1),
	0,
	math.Inf(1),
	math.NaN(),
}
var fabsSC = []float64{
	math.Inf(1),
	0,
	0,
	math.Inf(1),
	math.NaN(),
}

func alike(a, b float64) bool {
	switch {
	case math.IsNaN(a) && math.IsNaN(b):
		return true
	case a == b:
		return math.Signbit(a) == math.Signbit(b)
	}
	return false
}

func TestAbs(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := abs.Abs(vf[i]); fabs[i] != f {
			t.Errorf("Abs(%g) = %g, want %g", vf[i], f, fabs[i])
		}
	}
	for i := 0; i < len(vffabsSC); i++ {
		if f := abs.Abs(vffabsSC[i]); !alike(fabsSC[i], f) {
			t.Errorf("Abs(%g) = %g, want %g", vffabsSC[i], f, fabsSC[i])
		}
	}
}

var (
	GlobalI int
	GlobalB bool
	GlobalF float64
)

var absPos = .5

func BenchmarkAbs(b *testing.B) {
	x := 0.0
	for i := 0; i < b.N; i++ {
		x = abs.ArchAbs(absPos)
	}
	GlobalF = x

}

func BenchmarkAbsGo(b *testing.B) {
	x := 0.0
	for i := 0; i < b.N; i++ {
		x = abs.AbsGo(absPos)
	}
	GlobalF = x

}
