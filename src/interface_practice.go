package main

import "fmt"

/*
C++或者Java是需要主动声明基础类，而Golang，
只需要实现某个interface的全部方法，那么就是实现了该类型。
Golang的继承关系是非侵入式的，这也是Golang的特色与优点。
*/

// 必须实现接口中定义的所有函数
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	height float64
	width  float64
}

// 实现Rect结构体的接口
func (r *Rect) Area() float64 {
	return r.width * r.height
}

// 实现Rect结构体的接口
func (r *Rect) Perimeter() float64 {
	return r.width*2 + r.height*2
}

func interfaceTest() {
	//声明Shape接口类型的s其实现子类函数是Rect中的
	//Shape接口类型的s，指向一个Rect对象的，并调用其方法。
	var s Shape = &Rect{height: 10, width: 5}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}
