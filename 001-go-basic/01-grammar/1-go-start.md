# GO入门

> [Go 学习文档](http://www.topgoer.com/go%E5%9F%BA%E7%A1%80/)

## 1. 安装

```code
// github安装，网上自行找链接

// 查看是否安装成功
go version
```

## 2. Go程序

```go
package main // 可运行的应用程序

import "fmt"

func main() { // 应用程序的主入口
    fmt.Println("Hello, 世界");
}

```

Go语言程序的默认入口函数(主函数),在定义时不能有任何的参数和返回值，且Go程序自动调用。





## 3. Go命令行

```js
// 初始化一个目录
$ go mod init example.com/m
// 编译生成可执行文件main于当前目录下
$ go build main.go
// 测试 可执行文件 是否可用
./main
// 安装到 $GOBIN 目录，打开终端输入main都会执行 main
go install ./main.go
// go 跨平台编译
// GOOS: 要编译的目标操作系统
// GOARCH: 要编译的目标处理器架构
GOOS=linux GOARCH=amd64 go build ./main.go
```
