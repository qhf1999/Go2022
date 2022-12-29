package main

import "fmt"

var name string
var x = &name

func main() {
	name := "tom"
	y := &name
	fmt.Println("我是🐯")
	fmt.Println(x)
	fmt.Println(y)
}
