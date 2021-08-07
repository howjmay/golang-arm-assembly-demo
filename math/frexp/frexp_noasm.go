// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !arm64
// +build !arm64

package frexp

const haveArchFrexp = false

func ArchFrexp(x float64) (frac float64, exp int)
	panic("not implemented")
}
