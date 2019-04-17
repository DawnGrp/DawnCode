/*

这是一个范例

*/
package main

import "fmt"

//main 主函数
func main() {
	SayHi()

	fmt.Printf("%d", "9")
}

//SayHi 打印字符串Hello world
//go:generate echo Hello! $GOPATH
func SayHi() {
	fmt.Println("Hello world!!")
}
