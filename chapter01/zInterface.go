package main

import "fmt"

type datawriter interface {
	// DataWrite
	//方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
	DataWrite()
}

type Cat struct {
}

func (c Cat) DataWrite() {
	//TODO implement me
	fmt.Println("我是🐱")
}

func main() {
	var c datawriter = Cat{}
	c.DataWrite()

}
