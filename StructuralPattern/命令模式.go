package main

import "fmt"

/*
场景：
联想路边撸串烧烤场景， 有烤羊肉，烧鸡翅命令，有烤串师傅，和服务员。根据命令模式，设计烤串场景。
*/

type Cooker struct{}

func (c *Cooker) MakeChicken() {
	fmt.Println("烤鸡肉")
}

func (c *Cooker) MakeMutton() {
	fmt.Println("烤羊肉")
}

// 抽象的命令

type Command interface {
	Make()
}

// 具体的命令单子

type CommandCookChicken struct {
	cooker *Cooker
}

func (c *CommandCookChicken) Make() {
	c.cooker.MakeChicken()
}

type CommandCookMutton struct {
	cooker *Cooker
}

func (c *CommandCookMutton) Make() {
	c.cooker.MakeMutton()
}

// 服务员

type Waiter struct {
	cmdList []Command
}

func (w *Waiter) Notify() {
	if w.cmdList == nil {
		return
	}
	for _, cmd := range w.cmdList {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)
	cmdChicken := &CommandCookChicken{cooker}
	cmdMutton := &CommandCookMutton{cooker}

	waiter := new(Waiter)
	waiter.cmdList = append(waiter.cmdList, cmdChicken, cmdMutton)

	waiter.Notify()
}
