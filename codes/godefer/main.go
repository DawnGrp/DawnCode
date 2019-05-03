package main

import "fmt"

func main() {

	i := foo1()
	fmt.Println("foo1 return :", *i)

}

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

func foo1() *int {

	i := new(int)
	*i = 0

	defer func() {
		*i = 1
	}()

	return i
}

func foo3() int {

	//i = new(int)
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

func bar() string {
	var i = ""
	defer func() {
		i = "defer result in bar"
	}()

	i = "function result in bar"
	return i
}

func a() {
	i := 0
	defer fmt.Println(i)
	defer func(i *int) { fmt.Println(*i) }(&i)
	i++
	return
}

func c() int {
	i := 0
	defer func() { i = 3 }()
	i = 1
	return i
}

func c2() (i int) {
	defer func() { i = 3 }()
	i = 1
	return i
}
