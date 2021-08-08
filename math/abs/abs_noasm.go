//go:build !arm64
// +build !arm64

package abs

const haveArchAbs = false

func archAbs(x float64) float64 {
	panic("not implemented")
}
