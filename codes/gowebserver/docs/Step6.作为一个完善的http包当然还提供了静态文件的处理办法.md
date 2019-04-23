# 作为一个完善的http包当然还提供了静态文件的处理办法

以在index.html文件里引用一个index.js文件为例。

## 从代码开始，先来一个简单的例子

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

页面在浏览器中运行时，当运行到`<script src="/js/index.js"></script>`浏览器会请求 /js/index.js这个路径

程序检查到第一层路由匹配/js/，于是用文件服务处理这次请求，匹配到程序运行的路径下相对路径./static/js。

匹配的设置是 main.go文件中这两句

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

很简单...但是，可能还是不满足需求，因为

`http.Handle("/js/", http.FileServer(http.Dir("./static")))` 对应到 ./static/js

`http.Handle("/css/", http.FileServer(http.Dir("./static")))` 对应到 ./static/css

`http.Handle("/img/", http.FileServer(http.Dir("./static")))` 对应到 ./static/img

`http.Handle("/upload/", http.FileServer(http.Dir("./static")))` 对应到 ./static/upload

但是，这样所有请求的路径都必须匹配一个static目录下的子目录。如果，我就想访问static目录下的文件，或者，js、css、img、upload目录就在项目根目录下怎么办？

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

