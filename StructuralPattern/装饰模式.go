package main

import "fmt"

// >>>>抽象层<<<<

type Phone interface {
	ShowPhone()
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) ShowPhone() {}

// >>>>实现层<<<<

// 具体构件

type IPhone struct{}

func (i *IPhone) ShowPhone() {
	fmt.Println("IPhone...")
}

// 具体的装饰器类

type MoDecorator struct {
	// 继承基础装饰器类，继承phone成员属性
	Decorator
}

func (m *MoDecorator) ShowPhone() {
	m.phone.ShowPhone()
	fmt.Println("add mo")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (k *KeDecorator) ShowPhone() {
	k.phone.ShowPhone()
	fmt.Println("add ke")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}

func main() {
	var iphone Phone = new(IPhone)
	iphone.ShowPhone()
	fmt.Println("--------")
	var mp Phone = NewMoDecorator(iphone)
	mp.ShowPhone()
	fmt.Println("--------")
	var kp Phone = NewKeDecorator(iphone)
	kp.ShowPhone()
	fmt.Println("--------")
	var kmp Phone = NewKeDecorator(mp)
	kmp.ShowPhone()
}
