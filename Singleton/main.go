// Builder 设计模式·创建型模式·单例
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	instance     Singleton
	instanceOnce sync.Once
)

// Singleton 单例接口
type Singleton interface {
	Foo()
}

// MyStruct 自定义结构体
type MyStruct struct {
	NanoTime string
}

func (o MyStruct) Foo() { fmt.Println("NanoTime:", o.NanoTime) }

// 获取单例
func getInstance() Singleton {
	if instance != nil {
		return instance
	}
	instanceOnce.Do(func() {
		instance = &MyStruct{NanoTime: time.Now().Format(".999999999")}
	})
	return instance
	// return &MyStruct{NanoTime: time.Now().Format(".999999999")}
}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			ins := getInstance()
			ins.Foo()
		}()
	}

	fmt.Scanln()
}
