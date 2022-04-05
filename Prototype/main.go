// Builder 设计模式·创建型模式·原型
// 所有原型类都必须有一个通用的接口， 使得类在未知的情况下也能复制。
package main

import "fmt"

// Prototype 原型接口
type Prototype interface {
	Print()           // 打印
	Clone() Prototype // 复制
}

// Leaf 树叶
type Leaf struct{ Name string }

func (o *Leaf) Print()           { fmt.Println(o.Name) }
func (o *Leaf) Clone() Prototype { return &Leaf{Name: "[clone]" + o.Name} }

// LiuShu 柳树
type LiuShu struct {
	Name string
	Leaf []Prototype
}

func (o *LiuShu) Print() {
	fmt.Println(o.Name)
	for _, v := range o.Leaf {
		v.Print()
	}
}
func (o *LiuShu) Clone() Prototype {
	leaf := []Prototype{}
	for _, v := range o.Leaf {
		leaf = append(leaf, v.Clone())
	}
	return &LiuShu{Name: "[copy]" + o.Name, Leaf: leaf}
}

func main() {
	a := &LiuShu{
		Name: "tree",
		Leaf: []Prototype{&Leaf{Name: "a"}, &Leaf{Name: "b"}, &Leaf{Name: "c"}}}
	a.Print()

	b := a.Clone()
	b.Print()
}
