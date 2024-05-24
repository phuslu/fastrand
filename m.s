#include "textflag.h"

#ifdef GOARCH_amd64
TEXT ·getm(SB),NOSPLIT,$0-8
	MOVD g, R14
	MOVD 48(R14), R13
	MOVD R13, ret+0(FP)
	RET
#endif

#ifdef GOARCH_arm64
TEXT ·getm(SB),NOSPLIT,$0-8
	MOVD g, R14
	MOVD 48(R14), R13
	MOVD R13, ret+0(FP)
	RET
#endif
