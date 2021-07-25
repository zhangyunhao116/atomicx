// +build arm64,!gccgo,!appengine

#include "textflag.h"

TEXT ·CompareAndSwapUint128(SB),NOSPLIT,$0
	MOVD  addr+0(FP), R6
	MOVD  old_0+8(FP), R2
	MOVD  old_1+16(FP), R3
	MOVD  new_0+24(FP), R8
	MOVD  new_1+32(FP), R9
	CASPD (R2, R3), (R6), (R8, R9)
    CSET  EQ, R0
    MOVB  R0, swapped+40(FP)
	RET

TEXT ·LoadUint128(SB),NOSPLIT,$0
	RET
