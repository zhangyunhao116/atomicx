// +build amd64,!gccgo,!appengine

#include "textflag.h"


TEXT ·OrUint32(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVL val+8(FP), CX
	LOCK
	ORL  CX, (DX)
	RET

TEXT ·OrUint64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVQ val+8(FP), CX
	LOCK
	ORQ  CX, (DX)
	RET


TEXT ·XorUint32(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVL val+8(FP), CX
	LOCK
	XORL  CX, (DX)
	RET

TEXT ·XorUint64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVQ val+8(FP), CX
	LOCK
	XORQ  CX, (DX)
	RET

TEXT ·AndUint32(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVL val+8(FP), CX
	LOCK
	ANDL CX, (DX)
	RET

TEXT ·AndUint64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), DX
	MOVQ val+8(FP), CX
	LOCK
	ANDQ CX, (DX)
	RET
