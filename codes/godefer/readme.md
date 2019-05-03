# Golang研学：必须掌握的defer延迟执行

defer：在函数A内用defer关键字调用的函数B会在在函数A `return`后执行。

先看一个基础的例子，了解一下defer的效果

```go
func main() {
	fmt.Println("in main func:", foo())
}

func foo() int {
	i := 0
	defer fmt.Println("in defer :", i)
	i = 1000
	fmt.Println("in foo:", i)
	return i+24
}
```

这段代码运行后会打印出

```shell
in foo: 1000
in defer : 0
in main func: 1024
```

变量`i`初始化为`0`，`defer`指定`fmt.Println`函数延迟到`return`后执行，`main`函数中调用`foo`打印返回值，它是在`defer`之后执行。

## 有什么用途?

函数中会申明使用很多变量资源，函数结束时，我们通常会对它们做一些处理：销毁、释放（例如数据库链接、文件句柄、流）。

一般情况下，我们会在`return`语句之前处理这些事情。

但是，如果函数中包含多个`return`，这些处理我们需要在每个`return`之前都操作一次，实际工作中经常出现遗漏，代码维护时也很麻烦。

例如，不用`defer`的时候，可能会出现这样的情况：

```go
func foo(i int) int {
	if i > 100 {
		fmt.Println("不是期待的数字")
		return 0
	}

	if i < 50 {
		fmt.Println("不是期待的数字")
		return 0
	}

	return i
}
```

但是，如果用 `defer` 可以这样

```go
func foo(i int) int {
	defer func() {
		fmt.Println("不是期待的数字")
	}()

	if i > 100 {
		return 0
	}

	if i < 50 {
		return 0
	}

	return i
}
```

## 一个函数中多个defer的执行顺序是什么?

`defer`在同一个函数中可以使用多次。

多个`defer`指定的函数执行顺序是"先进后出"。

为什么呢 ？

可以这样理解：`defer`关键字会**在其以下的代码**执行到`return`后再执行，依次类推，第一个`defer`在其下面的第二个`defer`执行完后执行，第二个`defer`会在它下面的第三个`defer`执行完后执行...

这个顺序非常必要，因为在函数中，后面定义的对象可能依赖前面的对象，如果前面的`defer`先执行了，很可能造成后面的`defer`执行的时候出现异常。

所以，Go语言设计defer的时候是按先进后出的顺序执行的。

例子：

```go
func foo() {
	i := 0
	defer func() {
		i--
		fmt.Println("第一个defer", i)
	}()

	i++
	fmt.Println("+1后的i：", i)

	defer func() {
		i--
		fmt.Println("第二个defer", i)
	}()

	i++
	fmt.Println("再+1后的i：", i)

	defer func() {
		i--
		fmt.Println("第三个defer", i)
	}()

	i++
	fmt.Println("再+1后的i：", i)
}
```

运行后可以看到

```shell
+1后的i： 1
再+1后的i： 2
再+1后的i： 3
第三个defer 2
第二个defer 1
第一个defer 0
```

这个过程可以看出函数执行后，先进后出执行defer逐步处理变量的过程。

## 传递参数给defer指定的函数时，参数值是该函数执行处理后的值吗？

有一些总结是说：**defer指定的函数的参数在 defer 时确定**，但，这只是一个总结，真正的原因是， **Go语言除了map、slice、chan都是值传递**。

改造一下上面这个例子

```go
func foo() {
	i := 0
	defer func(k int) {
		fmt.Println("第一个defer", k)
	}(i)

	i++
	fmt.Println("+1后的i：", i)

	defer func(k int) {
		fmt.Println("第二个defer", k)
	}(i)

	i++
	fmt.Println("再+1后的i：", i)

	defer func(k int) {
		fmt.Println("第三个defer", k)
	}(i)

	i++
	fmt.Println("再+1后的i：", i)
}
```

得到的结果

```shell
+1后的i： 1
再+1后的i： 2
再+1后的i： 3
第三个defer 2
第二个defer 1
第一个defer 0
```

可能会有人觉得有一点出乎预料，代码里没有减数操作，为什么后执行的defer的`i`为什么defer打印出来的不是`3`呢？

defer关键字指定的函数是在`return`后执行的，这很容易让人想象在`return`后调用函数。

但是，defer指定的函数是在当前行就调用了的，只是**延迟**到`return`后执行，而不等同于“**移动**”到`return`后执行，因此调用时传递的是当前的参数的值。


## 传递指针参数会是什么情况？

那么如果希望`defer`指定的的函数参数的值是经过后面的代码处理过的，可以传递指针参数给`defer`指定的函数。

改造一下代码：

```go
func foo() {
	i := 0
	defer func(k *int) {
		fmt.Println("第一个defer", *k)
	}(&i)

	i++
	fmt.Println("+1后的i：", i)

	defer func(k *int) {
		fmt.Println("第二个defer", *k)
	}(&i)

	i++
	fmt.Println("再+1后的i：", i)

	defer func(k *int) {
		fmt.Println("第三个defer", *k)
	}(&i)

	i++
	fmt.Println("再+1后的i：", i)
}
```

运行后得到
```shell
+1后的i： 1
再+1后的i： 2
再+1后的i： 3
第三个defer 3
第二个defer 3
第一个defer 3
```

## defer会影响返回值吗？

在开头的第一个例子中可以看到，`defer`是在`foo`执行完，`main`里打印返回值之前执行的，但是没有影响到`main`里的打印结果。

还是相同的道理 **Go语言除了map、slice、chan都是值传递**

比较一下`foo1`和`foo2`两个函数的结果：

```go
func main() {

	fmt.Println("foo1 return :", foo1())
	fmt.Println("foot return :", foo2())

}

func foo1() int {

	i := 0

	defer func() {
		i = 1
	}()

	return i
}

func foo2() map[string]string {

	m := map[string]string{}

	defer func() {
		m["a"] = "b"
	}()

	return m
}
```

运行后，打印出

```shell
foo1 return : 0
foot return : map[a:b]
```

唯一不同的返回类型，**int类型return后，defer不会影响返回结果，但是map类型是引用传递，所以defer会改变返回结果。**

这说明，在return时，除了`map、slice、chan`，其他类型`return`时是将值拷贝到一个临时变量空间，因此，`defer`指定的函数内对函数内的变量的操作不会影响返回结果的。

**还有一种情况，给函数返回值申明变量名，**，这时，变量空间是在函数执行前申明出来，`return`时只是返回这个变量空间的内容，因此`defer`能够改变返回值。

例如，改造一下`foo1`函数，给它的返回值申明一个变量名`i`：

```go
func foo1() (i int) {

	i = 0

	defer func() {
		i = 1
	}()

	return i
}
```

再运行，可以看到 :

```shell
foo1 return : 1
```

返回值被`defer`指定的函数修改了。

## defer在panic和recover处理上的使用

在Go语言里，`defer`有一个经典的使用场景就是`recover`.

在函数执行过程中，有可能在很多地方都会出现`panic`，`panic`后如果不调用`recover`，程序会退出，为了不让程序退出，我们需要在`panic`后调用`recover`，但，`panic`后的代码不会执行，`recover`是不可能在`panic`后调用，然而`panic`所在的函数内`defer`指定的函数可以执行，所以`recover`只能在`defer`指定的函数中被调用。

例如:

```go
func panicfunc() {
	defer func() {
		fmt.Println("before recover")
		recover()
		fmt.Println("after recover")
	}()

	fmt.Println("before panic")
	panic(0)
	fmt.Println("after panic")
}
```

运行后，打印出：

```shell
before panic
before recover
after recover
```

## 总结以下

1. defer语句非常重要，非常常用，必须掌握
2. 在统一处理多个`return`和`panic/recover`场景下使用defer
3. 记住Go语言的函数参数传递的都是值，除了map、slice、chan，以便于正确的判断defer指定函数的参数值
4. defer不影响返回值，除非是map、slice和chan，或者返回值定义了变量名
5. 执行顺序：先进后出
