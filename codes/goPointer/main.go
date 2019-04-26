package main

import "fmt"

func main() {
	list := []int{1, 2, 3}

	println("list的地址: ", &list)

	for i := range list {
		fmt.Printf("第%d项的地址：%x \n", i, &list[i])
	}
}
