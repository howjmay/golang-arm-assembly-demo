//go:build gc
// +build gc

#include "textflag.h"

// func return_val(a uint64) (uint64, uint64)
TEXT ·return_val(SB),NOSPLIT,$0-8
	MOVD	a+0(FP), R0

	ADDS	R0, $1, R1
	MOVD	$0x3, R2
	MOVD	R1, ret+8(FP)
	MOVD	R2, ret+16(FP)

	RET
