# 第一章 Plan9汇编介绍

## SB寄存器

Static Base Pointer的缩写, 静态基地址寄存器, 表示静态内存的开始地址.

可以将SB想象为一个和内存容量有相同大小的字节数组, 所有的静态全局符号可以通过SB加偏移量进行定位.

Go汇编中定义的符号symbol其实就是相对于SB内存起始地址的偏移量, 也就是symbol = SB + offset

## 查看支持的汇编指令

### x86架构

定义在源码: `cmd/internal/obj/x86/anames.go`
