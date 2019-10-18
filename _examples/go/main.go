package main

import (
	"github.com/esilva-everbridge/yaml"
)

func main() {
	settings := yaml.New()
	err := settings.Set("success", true)
	if err != nil {
		panic(err)
	}
	err = settings.Set("nested", "tree", 1)
	if err != nil {
		panic(err)
	}
	err = settings.Set("another", "nested", "tree", []int{1, 2, 3})
	if err != nil {
		panic(err)
	}
	err = settings.Write("test.yaml")
	if err != nil {
		panic(err)
	}
}
