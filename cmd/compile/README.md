# go tool compile

## 设置软硬件编译平台

`GOOS=linux GOARCH=amd64 go tool compile -S main.go`

## 参数说明

1. `-V`: 查看编译器版本
2. `-S`: 打印编译出的汇编代码
3. `-N`: 编译时禁用优化
4. `-l`: 禁用内联
