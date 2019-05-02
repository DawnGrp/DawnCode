package main

import "fmt"

func main() {

	var s string

	println(*foo(&s))

	println(bar(s))

	a()

	i := c2()
	println(i)
}

func foo(i *string) *string {

	defer func() {
		*i = "defer result in foo1"
		println(*i)
	}()
	defer func() {
		*i = "defer result in foo2"
		println(*i)
	}()
	defer func() {
		*i = "defer result in foo3"
		println(*i)
	}()

	*i = "function result in foo"
	println(*i)
	return i
}

func bar(i string) string {

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
