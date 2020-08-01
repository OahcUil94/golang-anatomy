# Plan9汇编

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

var str = "hello world"
```

编译命令如下: 

- 命令1: `go tool compile -N -l -S main.go`
- 命令2: `go build -gcflags -S`
- 命令3: `go tool compile -N -l main.go`生成`main.o`文件, 再执行`go tool objdump main.o`

编译结果如下: 

```
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.string."hello world" SRODATA dupok size=11
	0x0000 68 65 6c 6c 6f 20 77 6f 72 6c 64                 hello world
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
"".str SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 0b 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."hello world"+0
```

### 将Go汇编代码转换为最终汇编代码



> 虽然go tool提供了asm这个工具, 但是该工具需要以Go汇编代码为输入, 输出的是编译之后的汇编代码.
> Go汇编语言中一个指令在最终的目标代码中可能会被编译为其它等价的机器指令.
> go tool asm -S main_amd64.s
 

AMD64的CPU是小端序的, 

%x, 9527

$表示常量

内存地址离的很近, 性能有可能会好

## 常见标识符

- TEXT: 告诉汇编器该数据放在TEXT区
- SB: 告诉汇编器这是基于静态地址的数据(static base)
- FP: 栈帧指针(Frame Pointer)

## 汇编指令

- LDP: Load double word pair 一次加载两个寄存器指令

## 参考资料

### 官方文档
- [官方文档asm](https://golang.org/doc/asm)

### Go语言中文网
- [Go 语言汇编快速入门](https://studygolang.com/articles/12828)

### 书籍
- [Go语言高级编程](http://books.studygolang.com/advanced-go-programming-book/)

### 其他文章
- [golang汇编](https://lrita.github.io/2017/12/12/golang-asm/)
- [得到Go程序的汇编代码的方法](https://colobu.com/2018/12/29/get-assembly-output-for-go-programs/)
- [Go Assembly by Example](https://www.davidwong.fr/goasm/)
- [Go 系列文章3 ：plan9 汇编入门](https://xargin.com/plan9-assembly/)
- [go 和 plan9 汇编](https://xargin.com/go-and-plan9-asm/)
