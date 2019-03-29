package main

import (
	"hello/utils"

	"github.com/astaxie/beego"
)

func main() {

	utils.PrintText("Hi")

	beego.Run()
}
