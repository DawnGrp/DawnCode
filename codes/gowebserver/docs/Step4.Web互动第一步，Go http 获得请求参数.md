# Web互动第一步，Go http 获得请求参数

## 还是先从代码开始

打开`main.go`文件，修改`myWeb`函数，如下:

```go
func myWeb(w http.ResponseWriter, r *http.Request) {

    r.ParseForm() //它还将请求主体解析为表单，获得POST Form表单数据，必须先调用这个函数

    for k, v := range r.URL.Query() {
        fmt.Println("key:", k, ", value:", v[0])
    }

    for k, v := range r.PostForm {
        fmt.Println("key:", k, ", value:", v[0])
    }

    fmt.Fprintln(w, "这是一个开始")
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

页面和终端命令行工具会答应出以下内容：

```shell
key: name , value: zeta
key: description , value: hello
```

## 解读

`http`请求的所有内容，都保存在`http.Request`对象中，也就是`myWeb`获得的参数 `r` 。

首先，调用`r.ParseForm()`，作用是填充数据到 `r.Form` 和 `r.PostForm`

接下来，分别循环获取遍历打印出 `r.URL.Query()` 函数返回的值 和 `r.PostForm` 值里的每一个参数。

`r.URL.Query()` 和 `r.PostForm` 分别是URL参数对象和表单参数对象
，它们都是键值对值，键的类型是字符串`string`，值的类型是`string`数组。

> 在http协议中，无论URL和表单，相同名称的参数会组成数组。

**循环遍历：for...range**

Go语言的循环只有`for`关键字，以下是Go中4种`for`循环

```go

//无限循环，阻塞线程，用不停息，慎用！
for{

}

//条件循环，如果a<b，循环，否则，退出循环
for a < b{

}

//表达式循环，设i为0，i小于10时循环，每轮循环后i增加1
for i:=0; i<10; i++{

}

//for...range 遍历objs，objs必须是map、slice、chan类型
for k, v := range objs{

}

```

前3种，循环你可以看作条件循环的变体（无限循环就是无条件的循环）。

本例种用到的是 `for...range` 循环，遍历可遍历对象，并且每轮循环都会将键和值分别赋值给变量 `k` 和 `v`

----

我们页面还是只是输出一句“**这是一个开始**”。我们需要一个可以见人的页面，这样可以不行

你也许也想到了，是不是可以在输出时，硬编码HTML字符串？当然可以，但是Go http包提供了更好的方式，HTML模版。

**接下来，我们就用HTML模版做一个真正的页面出来**