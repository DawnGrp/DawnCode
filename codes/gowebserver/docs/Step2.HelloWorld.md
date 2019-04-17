# Hello World！

环境说明：

* 开发工具：Visual Studio Code 
* 开发环境：MacOS
* Go版本：1.12.4

**命令行代码仅适用于Linux和MacOS系统，Windows根据说明在视窗下操作即可。**

## 1.创建项目

在你希望的地方创建一个文件夹，进入

`$ mkdir gowebserver && cd gowebserver `

新建一个文件 main.go

`$ touch main.go`

## 2. 用编辑器打开文件。输入以下代码：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}

```

## 3.打开命令行工具，输入以下命令

`$ go run main.go`

看到命令行工具会输出：

`Hello, 世界`

**第一个Go代码就完成了**

虽然这是一个很简单的Hello World，但是包含了Go语言编程的很多核心元素，下面就来详细讲解。

# 包


## `package`申明包 & `import`导入包
 
Go程序是由包构成的。

申明程序的包，用 `package` 关键字

程序中导入别的包使用 `import` 关键字。默认情况下，导入包的包名以导入路径的最后一个元素一致，例如 `import "math/rand"`，在代码中就直接使用`rand`作为包名。

## `func`定义函数

程序的入口都是main包下的 main函数 `main.main()`
