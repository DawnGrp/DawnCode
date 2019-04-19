# 获得http请求参数

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

首先，调用`r.ParseForm()`，他的作用填充数据到 r.Form 和 r.PostForm

接下来，分别循环获取遍历打印出 r.URL.Query() 函数返回的对象 和 r.PostForm 对象里的每一个参数。

`r.URL.Query()` 和 `r.PostForm` 分别保存的是URL参数对象和表单参数对象
，它们都是键值对对象，键的类型是字符串string，值的类型是string数组。

在http协议中，相同名称的参数会组成数组。

**循环遍历打印出参数 for...range**

Go语言的循环只有for，以下是Go中4种for循环

```go

//无限循环，阻塞线程，用不停息，慎用！
for{

}

//条件循环，如果a<b，循环，否则，退出循环
for a<b{

}

//初始化数据的条件循环，设i为0，i小于10时循环，每次轮循环后i=i+1
for i:=0;i<10;i++{

}

//for...range循环
for k,v:=range objs{

}

```

本例种用到的是 for...range循环，遍历可遍历对象，并且将键和值赋值给变量 k和v


但是，我们页面还是只是输出一句“我是一个开始”。

**接下来，我们就做一个真正的页面出来**