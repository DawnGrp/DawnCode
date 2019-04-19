package main

import (
	"fmt"
	"net/http"
)

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

func main() {
	http.HandleFunc("/", myWeb)
	fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
