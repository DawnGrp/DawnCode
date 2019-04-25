# 动态响应数据给访客，Go http HTML模版+数据绑定

用数据渲染HTML模版，生成网页，这是绝大部分Web开发框架的常规操作。

Go内置了一个包就提供了这样的功能：

"`html/template`"

## 从代码开始，先来一个简单的例子

main函数不做改变，只增加导入`html/template`包，并修改myWeb函数，如下：

```go

import (
    "fmt"
    "net/http"
    "text/template" //导入模版包
)

func myWeb(w http.ResponseWriter, r *http.Request) {

    t := template.New("index")

    t.Parse("<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>")

    data := map[string]string{
        "name":    "zeta",
        "someStr": "这是一个开始",
    }

    t.Execute(w, data)

    // fmt.Fprintln(w, "这是一个开始")
}

```

在命令行中运行 `$ go run main.go` ，访问 `http://localhost:8080`

看，`<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>` 中的{{.name}}和{{.someStr}}被替换成了 `zeta`和`这是一个开始`。并且，不再使用fmt.Fprintln函数输出数据到Response了

但是...这还是在代码里硬编码HTML字符串啊...

别着急，template包还提供了通过文件渲染，继续修改代码：

1. 根目录下创建一个子目录存放模版文件 templates, 然后创建一个文件 index.html，并写入一些HTML代码 (我不是个好前端)

```html

<html>
<head></head>
<body>
    <div>Hello {{.name}}</div>
    <div>{{.someStr}}</div>
</body>
</html>

```

2. 修改myWeb函数

```go

func myWeb(w http.ResponseWriter, r *http.Request) {

    //t := template.New("index")
    //t.Parse("<div>Hi,{{.name}},{{.someStr}}<div>")
    //将上两句注释掉，用下面一句
    t, _ := template.ParseFiles("./templates/index.html")

    data := map[string]string{
        "name":    "zeta",
        "someStr": "这是一个开始",
    }

    t.Execute(w, data)

    // fmt.Fprintln(w, "这是一个开始")
}

```

在运行一下看看，页面按照HTML文件的内容输出了，并且{{.name}}和{{.someStr}}也替换了，对吧？

## 解读

从这个例子中可以看到，template包的核心功能其实就是将HTML字符串解析暂存起来，然后调用Execute的时候，用数据替换掉HTML字符串中的`{{}}`里面的内容

在第一个方式中 `t:=template.New("index")` 初始化一个template对象变量，然后用调用`t.Parse`函数解析字符串模版。

然后，定一个map对象，渲染的时候会用到。

最后，调用`t.Execute`函数，不仅用数据渲染模版，还替代了fmtFprintln函数的工作，将输出到Response数据流写入器中。

第二个方式中，直接调用 template包的ParseFiles函数，直接解析相对路径下的index.html文件并创建对象变量。

**知识点**

范例中除了`html/template`的使用外，还出现两个新东西 map类型 和 赋值给“_”

**map类型**

map类型也就是其他语言中的字典类型（键值对），之前的获取请求参数章节中出现的url/values类型其实就是从map类型中扩展出来的

map的初始化可以使用make：

```go

var data = make(map[string]string)
data = map[string]string{}

```

> make是内置函数，只能用来初始化 map、slice 和 chan，并且make函数和另一个内置函数New不同点在于，它返回的并不是对象的指针，而只是一个类型。

map赋值于其他语言的字典对象相同，取值有两种方式，请看下面的代码：

```go

data["name"]="zeta" //赋值

name := data["name"] //1.普通取值

name,ok := data["name"] //2.如果不存在name键，ok为false

```

代码中的变量ok，可以用来判断这一项是否设置过，因为取值时，如果项不存在，是不会异常的，取出来的值该类型的零值，比如 int类型的值，不存在的项就为0；string类型的值不存在就为空字符串。

所以，ok这个变量可以判断是本来这一项的值就是零值，还是这一项不存在于map中。

Go中的map还有几个特点需要注意：

1. map的项的顺序是不固定的，每次遍历排列的顺序都是不同的，所以不能用顺序判断内容
2. map可以用for...range 遍历
3. map在函数参数中是引用传递（Go语言中，只有map、slice、chan是引用传递，其他都是值传递）
   
**赋值给 “_”**

Go有一个特点，变量定义后如果没使用，会报错，无法编译。一般情况下没什么问题，但是极少情况下，我们调用函数，但是并不需要使用返回值，但是不使用，又无法编译，怎么办？“_”就是用来解决这个问题的， 用 “_” 丢弃函数的返回值。比如本例中，`template.ParseFiles("./templates/index.html")` 除了返回模版对象外，还会返回一个error对象，但是这样简单的例子，出错的可能性极小，所以我不想处理error了，将error返回值用“_”丢弃掉。

> 在实际项目中，请不要丢弃error，任何意外都是可能出现的，丢弃error会导致当出现罕见的意外情况时，非常难于Debug。所有的error都应该要处理，至少写入到日志或打印到控制台。（切记，不要丢弃 error ，很多Gopher们在这个问题上有大把的血泪史）

OK，到目前为止，用Go语言搭建一个简单的网页的核心部分就完成了。

### 等等 .js、.css、图片怎么办？

哦，对。例子里的模版全是HTML代码，一个漂亮的网页还必须用到图片、js脚本和css样式文件，可是...和PHP不同，请求路径是通过HandleFunc匹配到处理函数的，难道要把js、css和图片都通过函数输出后，再用HandleFunc和URL路径匹配？

不，下一节揭晓。

