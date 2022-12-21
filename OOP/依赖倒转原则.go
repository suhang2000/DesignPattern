package main

import "fmt"

// 抽象层 --> 实现层 <-- 业务逻辑层

// >>>>抽象层<<<<
// 接口

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// >>>>实现层<<<<
// 类

// Car

type Benz struct{}

func (b *Benz) Run() {
	fmt.Println("Benz is running")
}

type BMW struct{}

func (b *BMW) Run() {
	fmt.Println("BMW is running")
}

// Driver

type ZhangSan struct{}

func (z *ZhangSan) Drive(car Car) {
	fmt.Println("zhang san drives car")
	car.Run()
}

type LiSi struct{}

func (l *LiSi) Drive(car Car) {
	fmt.Println("li si drives car")
	car.Run()
}

// >>>>业务逻辑层<<<<

func main() {
	var bmw Car // 声明抽象接口，而非具体实现类
	bmw = &BMW{}
	var z3 Driver
	z3 = &ZhangSan{}
	z3.Drive(bmw)

	var benz Car = &Benz{}
	var l4 Driver = &LiSi{}
	l4.Drive(benz)
}
