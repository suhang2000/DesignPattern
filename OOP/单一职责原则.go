package main

import "fmt"

type ClothesShop struct{}

type ClothesWork struct{}

func (cs *ClothesShop) OnShop() {
	fmt.Println("On Shop")
}

func (cw *ClothesWork) OnWork() {
	fmt.Println("On Work")
}

func main() {
	cs := new(ClothesShop)
	cs.OnShop()

	cw := new(ClothesWork)
	cw.OnWork()
}
