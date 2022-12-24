package main

import "fmt"

// 抽象工厂模式，适合产品族的设计

// ====抽象层====

// 抽象产品

type AbstractCPU interface {
	ShowCPU()
}

type AbstractGPU interface {
	ShowGPU()
}

// 抽象工厂

type AbstractFactory_ interface {
	CreateCPU() AbstractCPU
	CreateGPU() AbstractGPU
}

// ====实现层====

// 英特尔产品族

type IntelCPU struct{}

func (i *IntelCPU) ShowCPU() {
	fmt.Println("intel cpu")
}

type IntelGPU struct{}

func (i *IntelGPU) ShowGPU() {
	fmt.Println("intel GPU")
}

// 英特尔工厂

type IntelFactory struct{}

func (i *IntelFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU = new(IntelCPU)
	return cpu
}

func (i *IntelFactory) CreateGPU() AbstractGPU {
	var gpu AbstractGPU = new(IntelGPU)
	return gpu
}

// 英伟达产品族

type NvidiaCPU struct{}

func (n *NvidiaCPU) ShowCPU() {
	fmt.Println("nvidia cpu")
}

type NvidiaGPU struct{}

func (n *NvidiaGPU) ShowGPU() {
	fmt.Println("nvidia gpu")
}

// 英伟达工厂

type NvidiaFactory struct{}

func (n *NvidiaFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU = new(NvidiaCPU)
	return cpu
}

func (n *NvidiaFactory) CreateGPU() AbstractGPU {
	var gpu AbstractGPU = new(NvidiaGPU)
	return gpu
}

func main() {
	// 需求：英特尔的CPU和GPU
	// 1. create intel factory
	var intelFactory AbstractFactory_ = new(IntelFactory)
	// 2. produce cpu and gpu
	var intelCpu AbstractCPU = intelFactory.CreateCPU()
	intelCpu.ShowCPU()
	var intelGpu AbstractGPU = intelFactory.CreateGPU()
	intelGpu.ShowGPU()
	// 需求：英伟达的CPU和英特尔的GPU
	// 1. create factory
	var nvidiaFactory AbstractFactory_ = new(NvidiaFactory)
	// 2. produce nvidia cpu
	nvidiaCpu := nvidiaFactory.CreateCPU()
	nvidiaCpu.ShowCPU()
	intelGpu = intelFactory.CreateGPU()
	intelGpu.ShowGPU()
}
