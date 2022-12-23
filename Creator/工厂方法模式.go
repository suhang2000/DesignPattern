package main

import "fmt"

// >>>>抽象层<<<<

// Fruit 抽象接口

type Fruit interface {
	Show()
}

// Factory 抽象接口

type AbstractFactory interface {
	CreateFruit() Fruit
}

// >>>>基础类模块<<<<

type Apple struct {
	Fruit
}

func (a *Apple) Show() {
	fmt.Println("apple")
}

type Banana struct {
	Fruit
}

func (b *Banana) Show() {
	fmt.Println("banana")
}

// 工厂模块

type AppleFactory struct {
	AbstractFactory
}

func (a *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (b BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)
	return fruit
}

// >>>>业务逻辑层<<<<

func main() {
	// 需求：创建一个苹果对象
	// 1. 创建苹果工厂
	var appleFactory AbstractFactory
	appleFactory = new(AppleFactory)
	// 2. 创建对应水果对象
	var apple Fruit
	apple = appleFactory.CreateFruit()
	apple.Show()

	// 2. 创建一个香蕉对象
	var bananaFactory AbstractFactory = new(BananaFactory)
	var banana Fruit = bananaFactory.CreateFruit()
	banana.Show()
}
