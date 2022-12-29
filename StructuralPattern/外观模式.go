package main

import "fmt"

type TV struct{}

func (t *TV) On() {
	fmt.Println("TV on")
}

func (t *TV) Off() {
	fmt.Println("TV off")
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("Light on")
}

func (l *Light) Off() {
	fmt.Println("Light Off")
}

type Xbox struct{}

func (x *Xbox) On() {
	fmt.Println("打开 游戏机")
}

func (x *Xbox) Off() {
	fmt.Println("关闭 游戏机")
}

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("打开 投影仪")
}

func (p *Projector) Off() {
	fmt.Println("关闭 投影仪")
}

// facade

type HomePlayerFacade struct {
	tv        TV
	light     Light
	xbox      Xbox
	projector Projector
}

func (h *HomePlayerFacade) KTVMode() {
	fmt.Println("KTV Mode")
	h.tv.On()
	h.light.Off()
	h.projector.On()
}

func (h *HomePlayerFacade) GameMode() {
	fmt.Println("Game Mode")
	h.tv.On()
	h.light.On()
	h.xbox.On()
}

func main() {
	homePlayer := new(HomePlayerFacade)
	homePlayer.KTVMode()
	fmt.Println("--------")
	homePlayer.GameMode()
}
