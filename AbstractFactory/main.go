// Abstract Factory 对象创新模式·抽象工厂
package main

import "fmt"

// Factory 一个标准工厂
type Factory interface {
	Workshop() Workshop
	Eatery() Eatery
}

// Workshop 车间标准
type Workshop interface {
	Work()
}

// Eatery 食堂标准
type Eatery interface {
	Eating()
}

// FactoryA 工厂A
type FactoryAInfo struct{}

func (o *FactoryAInfo) Workshop() Workshop { return &FactoryA{} }
func (o *FactoryAInfo) Eatery() Eatery     { return &FactoryA{} }

// FactoryA 工厂A的车间和食堂（可以把车间和食堂才开）
type FactoryA struct{}

func (o *FactoryA) Work()   { fmt.Println("FactoryA::Work") }
func (o *FactoryA) Eating() { fmt.Println("FactoryA::Eating") }

// FactoryB 工厂A
type FactoryBInfo struct{}

func (o *FactoryBInfo) Workshop() Workshop { return &FactoryB{} }
func (o *FactoryBInfo) Eatery() Eatery     { return &FactoryB{} }

// FactoryB 工厂A的车间和食堂（可以把车间和食堂才开）
type FactoryB struct{}

func (o *FactoryB) Work()   { fmt.Println("FactoryB::Work") }
func (o *FactoryB) Eating() { fmt.Println("FactoryB::Eating") }

func FindWork(Money int) Factory {
	if Money > 10000 {
		return &FactoryAInfo{}
	} else {
		return &FactoryBInfo{}
	}
}

// golang 设计模式·对象创新模式·抽象工厂
func main() {
	w20k := FindWork(20000)
	w20k.Workshop().Work()
	w20k.Eatery().Eating()

	w2k := FindWork(2000)
	w2k.Workshop().Work()
	w2k.Eatery().Eating()
}
