package main

import "fmt"

func main() {
	str := "Beginning of the string " +
		"second part of the string"
	str2 := `Beginning of the string 
second part of the string`
	fmt.Println(str)
	fmt.Println(str2)
}
