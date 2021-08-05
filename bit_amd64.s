// +build amd64,!gccgo,!appengine

#include "textflag.h"

TEXT ·BitSetUint32(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVL offset+8(FP), CX
	LOCK
	BTSL CX, (DX)
	RET

TEXT ·BitSetUint64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVL offset+8(FP), CX
	LOCK
	BTSQ CX, (DX)
	RET

TEXT ·BitResetUint32(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVL offset+8(FP), CX
	LOCK
	BTRL CX, (DX)
	RET

TEXT ·BitResetUint64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVL offset+8(FP), CX
	LOCK
	BTRQ CX, (DX)
	RET
