package main

import "fmt"

// 适配目标 target

type V5 interface {
	Use5V()
}

// 业务类，依赖V5接口

type PhoneAda struct {
	v V5
}

func NewPhoneAda(v V5) *PhoneAda {
	return &PhoneAda{v}
}

func (p *PhoneAda) Charge() {
	fmt.Println("Phone charge")
	p.v.Use5V()
}

// 被适配者，Adaptee

type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("use 220v")
}

// 适配器

type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("use adapter")
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func main() {
	phone := NewPhoneAda(NewAdapter(new(V220)))
	phone.Charge()
}
