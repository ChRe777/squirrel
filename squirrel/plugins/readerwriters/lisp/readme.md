# Plugin Parser

## Build

	> go build -buildmode=plugin -o plugin_name.so


## Example simple plugin

	package main
	
	import "fmt"
	
	var V int
	
	func F() { fmt.Printf("Hello, number %d\n", V) }

## Example how to use plugin

	p, err := plugin.Open("plugin_name.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
