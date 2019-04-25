# 0依赖，创建一个Web服务

## 先从代码开始

打开之前创建好的`main.go`文件，修改代码如下:

```go
package main

import (
    "fmt"
    "net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "这是一个开始")
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

这时候，你会看到用 `fmt.Println`打印出来的提示，在浏览器中访问 `http://localhost:8080`你将访问到一个页面，显示 "**这是一个开始**"

## 解读

我们从程序运行的顺序去了解它的工作流程

首先，定义`package main`，然后导入包。

这里，导入了一个新的包 `net/http`，这个包是官方的，实现http客户端和服务端的各种功能。Go语言开发Web服务的所有功能就是基于这个包（其他第三方的Go语言Web框架也都基于这个包，没有例外）

### 先看`main`函数里发生了什么

第一句，匹配路由和处理函数

`http.HandleFunc("/", myWeb)`

调用http包的HandleFunc方法，匹配一个路由到一个处理函数`myWeb`。

这句代码的意思是，当通过访问地址 http://localhost/ 时，就等同于调用了 myWeb 函数。

第二句，用fmt在控制台打印一句话，纯属提示。

第三句，开启服务并且监听端口

```go
err := http.ListenAndServe(":8080", nil)
```

在这句，调用了`http`包中的`ListenAndServe`函数，该函数有两个参数，第一个是指定监听的端口号，第二个是指定处理请求的handler，通常这个参数填nil，表示使用默认的ServeMux作为handler。

**什么是nil？**

`nil`就是其他语言里的`null`。

> 什么是handler?什么是ServeMux？
> ServeMux就是一个HTTP请求多路由复用器。它将每个传入请求的URL与已注册模式的列表进行匹配，并调用与URL最匹配的模式的处理程序。
> 很熟悉吧？还记得前面的`http.HandleFunc`吗？他就是给http包中默认的ServeMux（DefaultServeMux）添加URL与处理函数匹配。
> 通常都是使用http包中的默认ServeMux，所以在`http.ListenAndServe`函数的第二个参数提供nil就可以了

`ListenAndServe`函数会一直监听，除非强制退出或者出现错误。

如果这句开启监听出现错误，函数会退出监听并会返回一个error类型的对象，因此用`err`变量接收返回对象。紧接着，判断`err`是否为空，打印出错误内容，程序结束。

---
这里有两个Go语言**知识点**

#### 1.定义变量

Go语言是静态语言，需要定义变量，定义变量用关键字`var`

``` go
 var   str   string = "my string"
//^      ^     ^
//关键字  变量名 类型
```

Go还提了一种简单的变量定义方式`:=`，**自动根据赋值的对象定义变量类型**，用起来很像脚本语言：

```go
str := "my string"
```

#### 2.错误处理

``` go
if err != nil{
    //处理....
}
```

在Go语言中，这是很常见的错误处理操作，另一种panic异常，官方建议不要使用或尽量少用，暂不做介绍，先从err开始。

Go语言中规定，如果函数可能出现错误，应该返回一个error对象，这个对象至少包含一个Error()方法错误信息。

因此，在Go中，是看不到try/catch语句的，函数使用error传递错误，用if语句判断错误对象并且处理错误。

---

#### 3. if 语句

与大多数语言使用方式一样，唯一的区别是，表达式不需要()包起来。

另外，Go语言中的if可以嵌入一个表达式，用;号隔开，例如范例中的代码可以改为：

```go
if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println("服务器开启错误: ", err)
}
```

err这个变量的生命周期只在`if`块中有效。

### 请求处理 myWeb函数

在`main`函数中，用`http.HandleFunc`将 myWeb与路由`/`匹配在一起。

`HandleFunc`函数定义了两个参数`w`,`r`，参数类型分别是`http.ResponseWriter`和`*http.Request`，`w`是响应留写入器，`r`是请求对象的指针。

**响应流写入器 w**: 用来写入http响应数据

**请求对象 * r**: 包含了http请求所有信息，注意，这里使用了指针，在定义参数时用`*`标记类型，说明这个参数需要的是这个类型的对象的**指针**。

当有请求路径`/`，请求对象和响应流写入器被传递给`myWeb`函数，并由`myWeb`函数负责处理这次请求。

**Go语言中红的指针**： 在Go语言中 除了map、slice、chan 其他函数传参都是**值传递**，所以，如果需要达到引用传递的效果，通过传递对象的指针实现。在Go语言中，取对象的指针用`&`，取值用`*`，例如：

```go
mystring := "hi"
//取指针
mypointer := &mystring
//取值
mystring2 := *mypointer

fmt.Println(mystring,mypointer,mystring2)
```
把这些代码放在`main`函数里，`$ go run main.go`运行看看


### myWeb函数体 

```go
fmt.Fprintf(w, "这是一个开始")
```

再一次遇到老熟人`fmt`，这次使用他的`Fprintf`函数将字符串“这是一个开始”,写入到`w`响应流写入器对象。`w`响应流写入器里写入的内容最后会被Response输出到用户浏览器的页面上。

### 总结一下，从编码到运行，你和它都干了些什么：

1. 定义一个函数myWeb，接收参数 响应流写入器和请求对象两个参数
2. 在main函数中，在默认的ServeMux中将路由/与myWeb绑定
3. 运行默认的ServeMux监听本地8080端口
4. 访问本地8080端口 `/` 路由
5. http将请求对象和响应写入器都传递给myWeb处理
6. myWeb向响应流中写入一句话，结束这次请求。

虽然代码很少很少，但是这就是一个最基本的Go语言Web服务程序了。

**下一节，看看如何获取请求参数**