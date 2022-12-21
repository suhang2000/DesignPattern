package main

import "fmt"

// AbstractBanker : 定义抽象接口和方法
type AbstractBanker interface {
	Business()
}

// 定义类及其实现方法

type SaveBanker struct{}

func (sb *SaveBanker) Business() {
	fmt.Println("Do save business")
}

type TransferBanker struct{}

func (tb *TransferBanker) Business() {
	fmt.Println("Do transfer business")
}

type PayBanker struct{}

func (pb *PayBanker) Business() {
	fmt.Println("Dp pay business")
}

func BankerBusiness(banker AbstractBanker) {
	banker.Business()
}

func main() {
	BankerBusiness(&SaveBanker{})
	BankerBusiness(&TransferBanker{})
	BankerBusiness(&PayBanker{})
}
