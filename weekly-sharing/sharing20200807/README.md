# 2020-08-07分享内容


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [2020-08-07分享内容](#2020-08-07分享内容)
  - [go unsafe包使用](#go-unsafe包使用)
  - [内存对齐](#内存对齐)
    - [对齐规则](#对齐规则)
    - [参考资料](#参考资料)
  - [go asm字符串](#go-asm字符串)
    - [疑问](#疑问)
    - [参考资料](#参考资料-1)

<!-- /code_chunk_output -->

## go unsafe包使用

- [https://www.jianshu.com/p/10b8870a9e8e](https://www.jianshu.com/p/10b8870a9e8e)
- [https://www.bilibili.com/video/BV15V411d7WS](https://www.bilibili.com/video/BV15V411d7WS)

## 内存对齐

### 对齐规则

- 平台最大对齐边界
- 结构体整体需要进行对齐, 对齐后的字节数需要是对齐字节的倍数 

### 参考资料

- [https://www.bilibili.com/video/BV1iZ4y1j7TT](https://www.bilibili.com/video/BV1iZ4y1j7TT)
- [https://www.bilibili.com/video/BV1Ja4y1i7AF](https://www.bilibili.com/video/BV1Ja4y1i7AF)
- [http://lzz5235.github.io/2015/04/21/memory.html](http://lzz5235.github.io/2015/04/21/memory.html)
- [https://zhuanlan.zhihu.com/p/25863918](https://zhuanlan.zhihu.com/p/25863918)
- [https://blog.csdn.net/he_and/article/details/78959916](https://blog.csdn.net/he_and/article/details/78959916)
- [http://blog.chinaunix.net/uid-28541347-id-5795423.html](http://blog.chinaunix.net/uid-28541347-id-5795423.html)

## go asm字符串

### 疑问

- 学习汇编的目的
- 汇编真的可以提高性能?
- bool值相同的, 做比较一定为true?
- recover无法捕获的错误
- 构造string

### 参考资料

- [http://books.studygolang.com/advanced-go-programming-book/ch3-asm/readme.html](http://books.studygolang.com/advanced-go-programming-book/ch3-asm/readme.html)
