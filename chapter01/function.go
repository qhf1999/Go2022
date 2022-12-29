package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}

func main() {
	var a, _ = GetData()
	fmt.Println(a)
	//函数也是一种类型
	f := GetData
	fmt.Println(f())

}
