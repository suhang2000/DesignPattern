package main

import "fmt"

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("eat...")
}

// 继承添加方法

type Cat struct {
	Animal
}

func (c *Cat) Sleep() {
	fmt.Println("cat sleep")
}

// 组合添加方法

type Cat2 struct {
	a *Animal
}

func (c *Cat2) Sleep() {
	fmt.Println("cat2 sleep")
}

func main() {
	// extend
	c := new(Cat)
	c.Eat()
	c.Sleep()
	// combination
	c2 := new(Cat2)
	c2.a = new(Animal)
	c2.a.Eat()
	c2.Sleep()
}
