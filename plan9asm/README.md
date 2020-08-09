# Plan9汇编

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Plan9汇编](#plan9汇编)
  - [学习汇编目的](#学习汇编目的)
  - [Go汇编说明](#go汇编说明)
  - [编译命令](#编译命令)
    - [将Go代码编译为最终汇编代码](#将go代码编译为最终汇编代码)
    - [将Go汇编代码转换为最终汇编代码](#将go汇编代码转换为最终汇编代码)
    - [指令标识说明](#指令标识说明)
      - [GLOBAL指令](#global指令)
      - [DATA指令](#data指令)
      - [SB标识](#sb标识)
      - [SDWARFINFO标识](#sdwarfinfo标识)
      - [SNOPTRDATA标识](#snoptrdata标识)
      - [dupok标识](#dupok标识)
      - [size标识](#size标识)
  - [汇编指令](#汇编指令)
  - [参考资料](#参考资料)
    - [官方文档](#官方文档)
    - [Go语言中文网](#go语言中文网)
    - [书籍](#书籍)
    - [其他文章](#其他文章)

<!-- /code_chunk_output -->

## 学习汇编目的

- 分析代码底层实现
- 压榨CPU性能

> 不要为了炫技而做一些有违常规的事情

## Go汇编说明

- Go汇编语言并不是一个独立的语言, 因此Go汇编程序无法独立使用, 需要与Go语言混编
- Go汇编代码用来定义变量和函数, 如果定义的变量和函数需要在Go语言代码中使用, 则需要在Go语言文件中声明
- Go汇编文件的后缀名是`.s`, 由于汇编代码依赖硬件平台架构, 所以在文件名后面最好附加硬件平台信息, 例如: `pkg_amd64.s`
- Go汇编相关资料官方仅提供了一篇文章说明[asm](https://golang.org/doc/asm), 更多的内容需要自己去查看源码以及通过工具进行调试和反编译
- Go汇编语言是一种高级的汇编语言, 和最终真实执行的代码并不完全等价

> 如何查看硬件架构: 通过`go env`查看环境变量`GOHOSTARCH`的内容即可

## 编译命令

### 将Go代码编译为最终汇编代码

准备如下代码: 

main.go文件

```go
package main

var ID = 1024
```

编译命令如下: 

- 命令1: `go tool compile -N -l -S main.go`
- 命令2: `go build -gcflags -S main.go`
- 命令3: `go tool compile -N -l main.go`生成`main.o`文件, 再执行`go tool objdump main.o`

> `compile`的`-N`参数表示禁用优化, `-l`参数表示禁用内联

编译结果如下:

```
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
"".ID SNOPTRDATA size=8
	0x0000 00 04 00 00 00 00 00 00                          ........
```

### 将Go汇编代码转换为最终汇编代码

准备如下代码: 

main.go文件

```go
package main

var ID int
```

main_amd64.s文件

```go
GLOBL ·ID(SB),$8

DATA ·ID+0(SB)/1,$0x00
DATA ·ID+1(SB)/1,$0x04
DATA ·ID+2(SB)/1,$0x00
DATA ·ID+3(SB)/1,$0x00
DATA ·ID+4(SB)/1,$0x00
DATA ·ID+5(SB)/1,$0x00
DATA ·ID+6(SB)/1,$0x00
DATA ·ID+7(SB)/1,$0x00

```

> 注意: 每一个`.s`文件的最后必须有一个空行, 否则就会报类似的错误`./main_amd64.s:xx: unexpected EOF`, 上面代码最后6个字节必须进行初始化, 否则可能会读取到内存中的脏值

编译命令如下: 

- `go tool asm -S main_amd64.s`

编译结果如下: 

```
"".ID SDATA size=8
	0x0000 00 04 00 00 00 00 00 00                          ........
```

### 指令标识说明

- `·`是一个特殊的unicode符号, `·`开头表示的是当前包的变量 
- Go汇编中`$`符号表示常量
- `amd64`的CPU是小端序的, 所以1024表示的十六进制不是`04 00`, 而是`00 04`

#### GLOBL指令

GLOBL指令用于将变量声明为global, 使用DATA结合GLOBL来定义一个变量, GLOBL指令语法`GLOBL symbol(SB), width`:

- symbol: 变量标识符, 如果变量需要导出, 需要和Go代码中的变量名保持一致
- width: 占用内存大小, 单位字节, 可以设置的值为1, 2, 4, 8

> 注意: 是`GLOBL`, 不要写成`GLOBAL`

#### DATA指令

用于把数据定义在内存的DATA段, 初始化变量, Data指令语法`DATA symbol+offset(SB)/width, value`:

- offset: 地址偏移量, 单位字节
- width: 初始化内存的宽度大小
- value: 要初始化的值, 常量以`$`符号开头

#### SB标识

Static Base Pointer静态基指针, 用来表示全局变量和函数(函数变量貌似也是带有SB标识的, 测试过但不确定)

#### SDWARFINFO标识

- S: static(猜测)
- DWARF: 全称`Debugging With Attributed Record Formats`, 编译器编译源代码生成调试信息的规范, 在Linux中被普遍使用
- INFO: 信息

#### SNOPTRDATA标识 

- NOPTR: no pointer, 不包含指针
- DATA: 数据

#### dupok标识

Duplicates Is OK, 标识是否允许符号冗余(同名符号), 暂时还未做进一步理解

#### size标识

表示占用的字节数大小

## 汇编指令

- LDP: Load double word pair 一次加载两个寄存器指令

## 参考资料

### 官方文档
- [官方文档asm](https://golang.org/doc/asm)

### Go语言中文网
- [Go 语言汇编快速入门](https://studygolang.com/articles/12828)
- [Go语言内幕(3):链接器,链接器,重定位](https://studygolang.com/articles/7208)

### 书籍
- [Go语言高级编程](http://books.studygolang.com/advanced-go-programming-book/)

### 其他文章
- [golang汇编](https://lrita.github.io/2017/12/12/golang-asm/)
- [得到Go程序的汇编代码的方法](https://colobu.com/2018/12/29/get-assembly-output-for-go-programs/)
- [Go Assembly by Example](https://www.davidwong.fr/goasm/)
- [Go 系列文章3 ：plan9 汇编入门](https://xargin.com/plan9-assembly/)
- [go 和 plan9 汇编](https://xargin.com/go-and-plan9-asm/)
