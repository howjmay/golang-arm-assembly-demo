//go:build gc
// +build gc

#include "textflag.h"

// func decl_array() (uint64, uint64)
TEXT ·decl_array(SB),NOSPLIT,$16-16
	MOVD	$0x3, R0
	MOVD	$0x2, R1
	MOVD	R0, arr-8(SP)
	MOVD	R1, arr-16(SP)

	// Take the addr from the lowest addr of the array
	MOVD	$arr-16(SP), R3
	VLD1	(R3), [V0.D2]
	VSHL	$1, V0.D2, V0.D2

	VMOV	V0.D[0], R4
	VMOV	V0.D[1], R5
	MOVD	R4, ret+0(FP)
	MOVD	R5, ret+8(FP)

	RET
