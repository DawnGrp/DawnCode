# 处理好js、css和图片，才能做漂亮的网页，Go http静态文件的处理办法

以在index.html文件里引用一个index.js文件为例。

## 从代码开始

```go
func main() {
     http.HandleFunc("/", myWeb)

     //指定相对路径./static 为文件服务路径
     staticHandle := http.FileServer(http.Dir("./static"))
     //将/js/路径下的请求匹配到 ./static/js/下
     http.Handle("/js/", staticHandle)

     fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
     err := http.ListenAndServe(":8080", nil)
     if err != nil {
          fmt.Println("服务器开启错误: ", err)
     }
}
```

在项目的根目录下创建static目录，进入static目录，创建js目录，然后在js目录里创建一个index.js文件。

```javascript
alert("Javascript running...");
```

打开之前的index.html文件,在</body>后面加上 `<script src="/js/index.js"></script>`

运行 `$ go run main.go`，访问 http://localhost:8080，页面会弹出提示框。

## 解读

页面在浏览器中运行时，当运行到`<script src="/js/index.js"></script>`浏览器会请求 `/js/index.js`这个路径

程序检查到第一层路由匹配`/js/`，于是用文件服务处理这次请求，匹配到程序运行的路径下相对路径`./static/js`。

匹配的设置是 `main.go`文件中这两句

```go
     //指定相对路径./static 为文件服务路径
     staticHandle := http.FileServer(http.Dir("./static"))
     //将/js/路径下的请求匹配到 ./static/js/下
     http.Handle("/js/", staticHandle)
```

也可以写成一句，更容易理解

```go
//浏览器访问/js/ 将会以静态文件形式访问目录 ./static/js
http.Handle("/js/", http.FileServer(http.Dir("./static")))
```

很简单...但是，可能还是不满足需求，因为, 如果

`http.Handle("/js/", http.FileServer(http.Dir("./static")))` 对应到 ./static/js

`http.Handle("/css/", http.FileServer(http.Dir("./static")))` 对应到 ./static/css

`http.Handle("/img/", http.FileServer(http.Dir("./static")))` 对应到 ./static/img

`http.Handle("/upload/", http.FileServer(http.Dir("./static")))` 对应到 ./static/upload

这样所有请求的路径都必须匹配一个static目录下的子目录。

如果，我就想访问static目录下的文件，或者，js、css、img、upload目录就在项目根目录下怎么办？

http包下，还提供了一个函数 `http.StripPrefix` 剥开前缀，如下：

```go
    //http.Handle("/js/", http.FileServer(http.Dir("./static")))
    //加上http.StripPrefix 改为 ：
    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static"))))
```

这样，浏览器中访问/js/时，直接对应到./static目录下，不需要再加一个/js/子目录。

所以，如果需要再根目录添加多个静态目录，并且和URL的路径匹配，可以这样：

`http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))` 对应到 ./js

`http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))` 对应到 ./css

`http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))` 对应到 ./img

`http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload"))))` 对应到 ./upload


到这里，一个从流程上完整的Web服务程序就介绍完了。

整理一下，一个Go语言的Web程序基本的流程：

1. 定义请求处理函数
2. 用http包的HandleFunc匹配处理函数和路由
3. ListenAndServe开启监听

当有http请求时：

1. http请求到监听的的端口
2. 根据路由将请求对象和响应写入器传递给匹配的处理函数
3. 处理函数经过一番操作后，将数据写入到响应写入器
4. 响应给请求的浏览器

## 最后编译程序

之前调试都使用的是 `go run` 命令运行程序。

您会发现，每次运行`go run`都会重新编译源码，如何将程序运行在没有Go环境的计算机上？

使用 `go build` 命令，它会编译源码，生成可执行的二进制文件。

最简单的 `go build` 命令什么参数都不用加，它会自动查找目录下的main包下的main()函数，然后依次查找依赖包编译成一个可执行文件。

其他依赖文件的相对路径需要和编译成功后的可执行文件一致，例如范例中的templates文件夹和static文件夹。

默认情况下，`go build`会编译为和开发操作系统对应的可执行文件，如果要编译其他操作系统的可执行文件，需要用到交叉编译。

例如将Linux和MacOSX系统编译到windows

`GOOS=windows GOARCH=amd64 go build`

在Windows上需要使用SET命令, 例如在Windows上编译到Linux系统
```shell
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

