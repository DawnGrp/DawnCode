package main

import "fmt"

func main() {

	m := map[string]string{
		"a": "b",
	}

	foo(&m)
}

func foo(m *map[string]string) {

	fmt.Println(&m)
}
