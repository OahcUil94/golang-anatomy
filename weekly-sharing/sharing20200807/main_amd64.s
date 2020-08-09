TEXT ·sum(SB), $0
    MOVQ arg1+0(FP), AX     // AX = a
    MOVQ arg2+8(FP), BX     // BX = b
    ADDQ AX, BX // BX += AX
    MOVQ BX, ret1+16(FP) // ret = BX
    RET
