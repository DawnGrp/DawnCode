# 四行代码的Hello World！所能表达出来的核心

**命令行代码仅适用于Linux和MacOS系统，Windows根据说明在视窗下操作即可。**

## 1.创建项目

创建一个文件夹，进入该文件夹

`$ mkdir gowebserver && cd gowebserver`

新建一个文件 main.go

`$ touch main.go`

## 2. 用编辑器打开文件，并输入以下代码：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

## 3.打开命令行终端，输入以下命令

`$ go run main.go`

看到终端会输出：

`Hello, 世界`

**第一个Go代码就完成了**

这是一个很简单的Hello World，但是包含了Go语言编程的许多核心元素，接下来就详细讲解。


## 解读知识点: 包 与 函数


### `package`申明包 & `import`导入包

Go程序是由包构成的。

代码的第一行, 申明程序自己的包，用 `package` 关键字。`package`关键字必须是第一行出现的代码。

范例代码中，申明的本包名 `main`

在代码中第二行, 导入“fmt”包, 使用 `import` 关键字。默认情况下，导入包的包名与导入路径的最后一个元素一致，例如 `import "math/rand"`，在代码中使用这个包时，直接使用`rand`，例如 `rand.New()`

导入包的写法可以多行,也可以“分组”, 例如:

```go
import "fmt"
import "math/rand"
```

或者 分组

```go
import (
    "fmt"
    "math/rand"
)
```

> fmt包是Go语言内建的包,作用是输出打印。

## `func`关键字：定义函数

`func`是function的缩写, 在Go语言中是定义函数的关键字。

func定义函数的格式为：

```go
func 函数名(参数1 类型,参数2 类型){
    函数体
}
```

本例中定义了一个main函数。`main`函数没有参数。
然后在`main`函数体里调用`fmt`包的`Println`函数，在控制台输出字符串 “Hello, 世界”

**所有Go语言的程序的入口都是main包下的main函数** `main.main()`，所以每一个可执行的Go程序都应该有一个`main`包和一个`main函数`。


**我们已经介绍了九牛一毛中的一毛，接下来正式通过搭建一个简单的Web服务学习Go语言**