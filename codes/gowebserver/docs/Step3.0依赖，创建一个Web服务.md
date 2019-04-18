# 0依赖，创建一个Web服务

## 先从代码开始

打开之前创建好的`main.go`文件，修改内容如下:

```go

package main

import (
    "fmt"
    "net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "我是一个开始")
}

func main() {
    http.HandleFunc("/", myWeb)
    fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("服务器开启错误: ", err)
    }
}

```

保存文件，然后在命令行工具下输入命令，运行程序

`$ go run main.go`

这时候，你会看到用 `fmt.Println`打印出来的提示，那么在浏览器中访问 `http://localhost:8080`你将访问到一个页面，显示 "**我是一个开始**"

## 解读

我们从程序运行的顺序去了解它的工作流程

首先，定义`package main`，然后导入包。

这里，导入了一个新的包 `net/http`，这个包是内置的，实现http客户端和服务端的各种功能。Web服务的所有功能就是基于这个包（以后接触到的所有Go语言Web框架也都基于这个包，没有例外）

接下来，先看`main`函数

第一句，绑定路由和处理函数

`http.HandleFunc("/", myWeb)`

调用http包的HandleFunc方法，绑定一个路由到一个函数myWeb。

这句代码的意思是，当通过访问地址 http://localhost/ 时，就等同于调用了 myWeb 函数。

第二句，仅是在控制台打印一句话

第三句，开启服务并且监听端口

` err := http.ListenAndServe(":8080", nil)`