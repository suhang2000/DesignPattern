package main

import "fmt"

// 抽象类，定义骨架，包含一个模版的全部实现步骤

type Backbone interface {
	BoilWater()     // 煮开水
	Brew()          // 冲泡
	Pour()          // 倒入杯中
	AddTop()        // 添加配料
	IfAddTop() bool // 是否添加配料
}

// 封装流程模版，让子类继承且实现

type template struct {
	b Backbone
}

// 封装的固定模版

func (t *template) DoTemplate() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.Pour()
	if t.b.IfAddTop() {
		t.b.AddTop()
	}
}

// >>>>具体模版<<<<

// coffee

type MakeCoffee struct {
	template
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (m *MakeCoffee) BoilWater() {
	fmt.Println("boil water to 100")
}

func (m *MakeCoffee) Brew() {
	fmt.Println("brew coffee")
}

func (m *MakeCoffee) Pour() {
	fmt.Println("pour into cup")
}

func (m *MakeCoffee) AddTop() {
	fmt.Println("add milk and sugar")
}

func (m *MakeCoffee) IfAddTop() bool {
	return true
}

// tea

type MakeTea struct {
	template
}

func (m *MakeTea) BoilWater() {
	fmt.Println("boil water to 80")
}

func (m *MakeTea) Brew() {
	fmt.Println("brew tea")
}

func (m *MakeTea) Pour() {
	fmt.Println("pour into cup")
}

func (m *MakeTea) AddTop() {
	fmt.Println("add lemon")
}

func (m *MakeTea) IfAddTop() bool {
	return false
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func main() {
	// 1. make coffee
	NewMakeCoffee().DoTemplate()

	fmt.Println("--------")
	//2. make tea
	NewMakeTea().DoTemplate()
}
