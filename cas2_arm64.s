// +build arm64,!gccgo,!appengine

#include "textflag.h"

TEXT ·CompareAndSwapUint128(SB),NOSPLIT,$0
	MOVD  addr+0(FP), R7
	MOVD  old_0+8(FP), R2
	MOVD  old_1+16(FP), R3
	MOVD  R2, R4
	MOVD  R3, R5
	MOVD  new_0+24(FP), R8
	MOVD  new_1+32(FP), R9
	CASPD (R2, R3), (R7), (R8, R9)
	EOR   R4, R2
	EOR   R5, R3
	ORR   R2, R3
	CMP   $0, R3
	CSET  EQ, R0
	MOVB  R0, swapped+40(FP)
	RET

TEXT ·LoadUint128(SB),NOSPLIT,$0
	MOVD  addr+0(FP), R7
	MOVD  ZR, R2
	MOVD  ZR, R3
	MOVD  ZR, R8
	MOVD  ZR, R9
	CASPD (R2, R3), (R7), (R8, R9)
	MOVD  R2, val_0+8(FP)
	MOVD  R3, val_1+16(FP)
	RET
