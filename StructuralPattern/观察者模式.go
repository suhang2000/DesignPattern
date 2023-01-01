package main

import "fmt"

// >>>>抽象层<<<<

type Listener interface {
	OnTeacherComing()
}

type Notifier interface {
	AddListener(listener Listener)    // attach
	RemoveListener(listener Listener) // detach
	Notify()
}

// >>>>实现层<<<<
// 观察者 学生

type Stu1 struct {
	BadThing string
}

func (s *Stu1) DoBadThing() {
	fmt.Println("Stu1 正在 ", s.BadThing)
}

func (s *Stu1) OnTeacherComing() {
	fmt.Println("Stu1 停止 ", s.BadThing)
}

type Stu2 struct {
	BadThing string
}

func (s *Stu2) DoBadThing() {
	fmt.Println("Stu2 正在 ", s.BadThing)
}

func (s *Stu2) OnTeacherComing() {
	fmt.Println("Stu2 停止 ", s.BadThing)
}

type Stu3 struct {
	BadThing string
}

func (s *Stu3) DoBadThing() {
	fmt.Println("Stu3 正在 ", s.BadThing)
}

func (s *Stu3) OnTeacherComing() {
	fmt.Println("Stu3 停止 ", s.BadThing)
}

// 通知者 班长
// 实现Notifier接口

type Monitor struct {
	listenerList []Listener
}

func (m *Monitor) AddListener(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}

func (m *Monitor) RemoveListener(listener Listener) {
	for i, l := range m.listenerList {
		if listener == l {
			m.listenerList = append(m.listenerList[:i], m.listenerList[i+1:]...)
			break
		}
	}
}

func (m *Monitor) Notify() {
	for _, listener := range m.listenerList {
		listener.OnTeacherComing()
	}
}

func main() {
	s1 := &Stu1{BadThing: "抄作业"}
	s2 := &Stu2{BadThing: "打游戏"}
	s3 := &Stu3{BadThing: "看视频"}
	monitor := new(Monitor)
	monitor.AddListener(s1)
	monitor.AddListener(s2)
	monitor.AddListener(s3)

	fmt.Println("上课了，老师还没来")
	s1.DoBadThing()
	s2.DoBadThing()
	s3.DoBadThing()

	fmt.Println("老师来了，班长通知学生们")
	monitor.Notify()
}
