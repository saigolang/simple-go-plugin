package main

import (
	"fmt"
	"path/filepath"
	"plugin"
)

func main() {

	plugins, err := filepath.Glob("add.so")

	if err != nil {
		panic(err)
	}

	p, err := plugin.Open(plugins[0])

	if err != nil {
		panic(err)
	}

	symbol, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}

	addFunc, ok := symbol.(func(int, int) int)
	if !ok {
		panic("Plugin has not add function")
	}

	addition := addFunc(3, 4)

	fmt.Println("add result is ", addition)
}
