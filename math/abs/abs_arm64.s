#include "textflag.h"

// func archAbs(x float64) float64
TEXT ·ArchAbs(SB), NOSPLIT, $0-16
    FMOVD   x+0(FP), F0
    FABSD   F0, F0
    FMOVD   F0, ret+8(FP)
    RET
