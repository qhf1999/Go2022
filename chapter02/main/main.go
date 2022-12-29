package main

import (
	"Go2022/chapter02/model"
	"fmt"
	"strconv"
)

func main() {
	p := model.NewPerson("史密斯")
	p.SetAge(23)
	p.SetSalary(99999)
	fmt.Println(p)
	fmt.Println(p.Name + "年龄为" + strconv.Itoa(p.GetAge()) +
		",月薪为" + strconv.FormatFloat(p.GetSalary(), 'f', 1, 64))

}
