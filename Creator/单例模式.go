package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 一个类只能有一个实例
// 它可以自行创建这个实例
// 它会向外提供这个实例

// ====饿汉式====

// 1. private constructor

type singleton1 struct{}

// 2. private pointer
// 指针不能被外部模块访问，以免被修改指向对象，只向外提供读权限
var instance1 *singleton1 = new(singleton1)

// 3. provide instance

func GetInstance1() *singleton1 {
	return instance1
}

// ====懒汉式====

var lock sync.Mutex
var initialized uint32

type singleton2 struct{}

var instance2 *singleton2

func GetInstance2() *singleton2 {
	// 如果已初始化，直接返回
	// 使用原子类，比加锁消耗更小
	if atomic.LoadUint32(&initialized) == 1 {
		return instance2
	}
	// 加锁
	lock.Lock()
	defer lock.Unlock()
	if initialized == 0 {
		instance2 = new(singleton2)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance2
}

// ====懒汉式once====

var once sync.Once

type singleton3 struct{}

var instance3 *singleton3

func GetInstance3() *singleton3 {
	once.Do(func() {
		instance3 = new(singleton3)
	})
	return instance3
}

func main() {
	// 饿汉式
	i1 := GetInstance1()
	i2 := GetInstance1()
	fmt.Println(i1 == i2)
	// 懒汉式
	i3 := GetInstance2()
	i4 := GetInstance2()
	fmt.Println(i3 == i4)
	// once
	i5 := GetInstance3()
	i6 := GetInstance3()
	fmt.Println(i5 == i6)
	// fmt.Println(new(singleton3) == new(singleton3))
}
