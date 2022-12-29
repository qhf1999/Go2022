// package学习
package model

type person struct {
	Name   string
	age    int
	salary float64
}

// 类似于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	p.age = age
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	p.salary = salary
}

func (p *person) GetSalary() float64 {
	return p.salary
}
