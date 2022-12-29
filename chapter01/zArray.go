package main

import "fmt"

var a []int

func main() {
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	a = append(a, 1, 1, 1)

	fmt.Println(len(a))
}
