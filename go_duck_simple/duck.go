package main

import (
	"fmt"
)

type Quacker interface {
	Quack() string
}

// Duck implements a Quack() string method and count as a Quacker
type Duck struct{}

func (d Duck) Quack() string {
	return "quack"
}

// Cat implements a Quack() string method and count as a Quacker.
type Cat struct{}

func (c Cat) Quack() string {
	return "meow"
}

// Dog does not implement a Quack() string method and does not count as a Quacker.
type Dog struct{}

func (d Dog) Bark() string {
	return "bark"
}

func MakeQuack(x Quacker) {
	fmt.Println(x.Quack())
}

func main() {
	duck := Duck{}
	cat := Cat{}
	dog := Dog{}

	MakeQuack(duck)
	MakeQuack(cat)
	MakeQuack(dog) // Compile error because Dog cannot be used as a Quacker type
}
