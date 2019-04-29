package main

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	sidebars := make(map[string]map[string]string)

	data := []byte(`
菜单1:
  子菜单1: m11.html
菜单2:
  子菜单1: m31.html
  子菜单2: m22.html
菜单3:
  子菜单1: m31.html
`)

	yaml.Unmarshal(data, &sidebars)

	// fmt.Printf("%v\n", sidebars)

	for key, children := range sidebars {
		log.Println(key)
		for title, child := range children {
			log.Println(title + ": " + child)
		}
	}
}
