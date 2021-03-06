package main

import (
	"fmt"
)

type Quacker interface {
	Quack() string
}

type Duck struct{}

func (d Duck) Quack() string {
	return "quack"
}

type Cat struct{}

func (c *Cat) Quack() string {
	return "meow"
}

type Dog struct{}

func (d Dog) Bark() string {
	return "bark"
}

func MakeQuack(x Quacker) {
	fmt.Println(x.Quack())
}

func main() {
	var x Quacker
	// Both instances of Duck are Quackers. This is because the receive type for the Quack()
	// method is a reference type (d Duck) and pointers have the method set of both T and *T.
	x = Duck{}
	MakeQuack(x)
	x = &Duck{}
	MakeQuack(x)

	// Only the pointer type of Cat is a Quacker. This is because the receiver type for then
	// Quack() method is a pointer type (c *Cat). Reference types (T) only have the method
	// set of T whereas pointer types have both T and *T.
	x = Cat{} // Compile error
	MakeQuack(x)
	x = &Cat{}
	MakeQuack(x)

	// Dogs don't quack.
	x = Dog{} // Compile error
	MakeQuack(x)
	x = &Dog{} // Compile error
	MakeQuack(x)
}
