package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func myWeb(w http.ResponseWriter, r *http.Request) {

	//t := template.New("index")
	//t.Parse("<div>Hi,{{.name}},{{.someStr}}<div>")

	t, _ := template.ParseFiles("./templates/index.html")

	var data = make(map[string]string)
	data = map[string]string{
		//"name":    r.URL.Query()["name"][0],
		"name":    "zeta",
		"someStr": "这是一个开始",
	}

	t.Execute(w, data)

	// fmt.Fprintln(w, "这是一个开始")
}

func main() {

	http.HandleFunc("/", myWeb)

	staticHandle := http.FileServer(http.Dir("./static"))
	http.Handle("/js/", staticHandle)
	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static"))))

	fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
