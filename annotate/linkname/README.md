# go:linkname注释

该指令告诉编译器, 当前文件中的变量或函数在编译时链接到指定的变量或者函数上

## 注意事项

### 导入unsafe包

由于该指令破坏了类型系统和包的模块化, 因此使用时必须导入unsafe包, 例如源码里的`runtime/timeasm.go`文件:

```
package runtime

import _ "unsafe"

//go:linkname time_now time.now
func time_now() (sec int64, nsec int32, mono int64)
```

### linkname用法

- linkname指定的两个函数, 必须一个有函数体, 另一个没有函数体
- 使用linkname的文件需要导入`unsafe`包, 没有函数体的文件, 需要导入有函数体的包名(内置包就不需要了)
- linkname可以链接自己, `//go:linkname aSay a.say`, 在其它地方使用时, 不需要加包路径前缀了, 例如: `//go:linkname bSay a.say`, 而不是`//go:linkname bSay pkgpath/a.say`  

### 添加空asm文件

如果linkname加在了函数体文件上, 并且没有函数体的函数名是大写的, 运行代码会报`missing function body`, 原因在于Go在编译的时候会启用-complete编译器flag, 它要求所有的函数必需包含函数体, 创建一个空的汇编语言文件绕过这个限制 

## 参考资料

- [突破限制,访问其它Go package中的私有函数](https://colobu.com/2017/05/12/call-private-functions-in-other-packages/)
- [在 golang 中如何调用私有函数（绑定隐藏的标识符）](https://zhuanlan.zhihu.com/p/135224673)
