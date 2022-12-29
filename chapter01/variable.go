package main

import "fmt"

var name string
var x = &name

func main() {
	name := "tom"
	y := &name
	fmt.Println("我是齐红飞")
	fmt.Println(x)
	fmt.Println(y)
}
