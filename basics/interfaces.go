package main

import "fmt"

type Animal interface {
	makesound() string
}

type Dog struct {
	name string
	age  int32
}

type Cat struct {
	name string
	age  int32
}

func (dog Dog) makesound() string {
	return "bark"
}

func (cat Cat) makesound() string {
	return "meow"
}

func main() {

	d := Dog{name: "John", age: 10}
	c := Cat{name: "Tobby", age: 11}

	fmt.Println(d.makesound())
	fmt.Println(c.makesound())

}
