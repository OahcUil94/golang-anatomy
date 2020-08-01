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

> 如何查看硬件架构: 通过`go env`查看环境变量`GOHOSTARCH`的内容即可 

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

- [官方文档asm](https://golang.org/doc/asm)
- [书籍: Go语言高级编程](http://books.studygolang.com/advanced-go-programming-book/)
