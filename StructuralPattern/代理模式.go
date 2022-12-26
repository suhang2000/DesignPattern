package main

import "fmt"

// >>>>抽象层<<<<
// 抽象的购物主题Subject

type Shopping interface {
	Buy()
}

// >>>>实现层<<<<
// 具体的购物主题，实现Shopping接口

type JapanShopping struct{}

func (j *JapanShopping) Buy() {
	fmt.Println("Japan shopping")
}

type AmericanShopping struct{}

func (a *AmericanShopping) Buy() {
	fmt.Println("American shopping")
}

// >>>>代理<<<<

type Proxy struct {
	shopping Shopping
}

func (p *Proxy) Buy() {
	p.before()
	p.shopping.Buy()
	p.after()
}

func (p *Proxy) before() {
	fmt.Println("before...")
}

func (p *Proxy) after() {
	fmt.Println("after...")
}

func NewProxy(shopping Shopping) Shopping {
	return &Proxy{shopping: shopping}
}

func main() {
	var s Shopping = new(JapanShopping)
	var proxy Shopping = NewProxy(s)
	proxy.Buy()
	proxy = NewProxy(new(AmericanShopping))
	proxy.Buy()
}
