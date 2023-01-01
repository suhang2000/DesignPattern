package main

import "fmt"

// >>>>抽象策略<<<<

type WeaponStrategy interface {
	UseWeapon()
}

// 具体的策略

type AK47 struct{}

func (a *AK47) UseWeapon() {
	fmt.Println("Use AK47 to fight")
}

type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("use Knife to fight")
}

// 环境类

type Hero struct {
	strategy WeaponStrategy
}

func (h *Hero) SetStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func main() {
	hero := Hero{}
	// strategy 1
	hero.SetStrategy(new(AK47))
	hero.Fight()
	// strategy 2
	hero.SetStrategy(new(Knife))
	hero.Fight()
}
