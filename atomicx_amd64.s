// +build amd64,!gccgo,!appengine

#include "textflag.h"

TEXT ·CompareAndSwapUint128(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	MOVQ old_0+8(FP), AX
	MOVQ old_1+16(FP), DX
	MOVQ new_0+24(FP), BX
	MOVQ new_1+32(FP), CX
	LOCK
	CMPXCHG16B (R8)
	SETEQ swapped+40(FP)
	RET

TEXT ·LoadUint128(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	XORQ AX, AX
	XORQ DX, DX
	XORQ BX, BX
	XORQ CX, CX
	LOCK
	CMPXCHG16B (R8)
	MOVQ AX, val_0+8(FP)
	MOVQ DX, val_1+16(FP)
	RET

TEXT ·BitTestAndSetUint32(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVL offset+8(FP), CX
	LOCK
	BTSL CX, (DX)
	RET

TEXT ·BitTestAndSetUint64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVQ offset+8(FP), CX
	LOCK
	BTSQ CX, (DX)
	RET
