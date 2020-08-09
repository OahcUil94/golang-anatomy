# Delve调试器

[https://github.com/go-delve/delve](https://github.com/go-delve/delve)

## 安装

把代码仓库直接拷贝到`$GOPATH/src/github.com/go-delve/delve`, 然后执行`make install`

> 如果github访问过慢, 可以尝试`github.com.cnpmjs.org`访问, `git clone https://github.com.cnpmjs.org/go-delve/delve.git`
> 在`make install`过程中, 会下载其他的依赖, 如果下载不成功, 需要配置`GOPROXY="https://goproxy.cn,direct"`
