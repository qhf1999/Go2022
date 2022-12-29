package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
	c <- sum

}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	for i := 1; i < 10; i++ {
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
	}
	fmt.Println("我是main")
	//x, y := <-c, <-c
	//
	//fmt.Println(x, y, x+y)

}
