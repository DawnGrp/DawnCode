# 处理http请求参数

## 还是先从代码开始

打开main.go文件，只修改myWeb函数，如下:

```go

func myWeb(w http.ResponseWriter, r *http.Request) {

    r.ParseForm() //它还将请求主体解析为表单，获得POST Form表单数据，必须先调用这个函数

    for k, v := range r.URL.Query() {
        fmt.Println("key:", k, ", value:", v[0])
    }

    for k, v := range r.PostForm {
        fmt.Println("key:", k, ", value:", v[0])
    }

    fmt.Fprintln(w, "我是一个开始")
}

```

只是增加了两个循环，分别将请求的URL里的参数和POST表单的里参数打印出来。

运行程序

`$ go run main.go`

然后用任何工具（推荐Postman）提交一个POST请求，并且带上URL参数，或者在命令行中用cURL提交

```shell

curl --request POST \
  --url 'http://localhost:8080/?name=zeta' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data description=hello

```

这时候，可以看到，命令行工具会答应出以下内容：

```shell

key: name , value: zeta
key: description , value: hello

```

## 解读

http请求的所有内容，都保存在http.Request对象中，也就是myWeb获得的参数 r 。

首先，调用`r.ParseForm()`，他的作用是将 r.Form and r.PostForm 对象填充数据，如果不调用这个函数，我们将无法通过r.Form和r.PostForm获得数据。

接下来，分别循环获取 r.URL.Query() 函数返回的对象 和 r.PostForm 对象里的数据。

r.URL.Query() 函数将返回本次请求中的URL参数的对象

r.PostForm 对象保存的是Form表单提交的数据的对象

r.URL.Query() 和 r.PostForm 字典都是url.Vaules对象，它是一个键值对对象，键的类型是string，值的类型是string数组，因为同名参数是被组合成数组的。