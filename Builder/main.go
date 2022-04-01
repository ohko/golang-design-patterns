// Builder 设计模式·创建型模式·生成器
package main

import "fmt"

// Builder 生成器接口定义
type Builder interface {
	Produce()   // 生产
	Check()     // 检测
	DeepCheck() // 深度检测
}

// FactoryA 工厂A
type FactoryA struct{}

func (o *FactoryA) Produce()   { fmt.Println("FactoryA::Produce") }
func (o *FactoryA) Check()     { fmt.Println("FactoryA::Check") }
func (o *FactoryA) DeepCheck() { fmt.Println("nothing") }

// FactoryB 工厂B
type FactoryB struct{}

func (o *FactoryB) Produce()   { fmt.Println("FactoryB::Produce") }
func (o *FactoryB) Check()     { fmt.Println("FactoryB::Check") }
func (o *FactoryB) DeepCheck() { fmt.Println("FactoryB::DeepCheck") }

// 导向器 (主管/导演)
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

// Build 普通生产
func (o *Director) Build() {
	o.builder.Produce()
	o.builder.Check()
}

// BuildForVip 高端生产
func (o *Director) BuildForVip() {
	o.builder.Produce()
	o.builder.Check()
	o.builder.DeepCheck()
}

// FindBuilder 创建生成器
func FindBuilder(money int) Builder {
	if money > 10000 {
		return &FactoryB{}
	} else {
		return &FactoryA{}
	}

}

func main() {
	normal := FindBuilder(2000)
	normalDirector := NewDirector(normal)
	normalDirector.Build()
	normalDirector.BuildForVip()

	vip := FindBuilder(20000)
	vipDirector := NewDirector(vip)
	vipDirector.Build()
	vipDirector.BuildForVip()
}
