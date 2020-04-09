package main

import "fmt"

// Author: Mr.Liang

type One struct {
	a int
}

func (one One) Get()  int{
	return one.a
}

func (one *One) Set(a int)  {
	one.a = a
}

func (one One) SetTwo(a int)  {
	one.a = a
}

func main() {

	t := &One{a:1}
	// 方法表达式，相当于func Get(one One) int
	fmt.Println(One.Get(*t))
	// 相当于一个闭包
	get := One.Get
	fmt.Println(get(*t))

	// 方法表达式传参
	// 这里特别需要注意*One，Set第一个参数值是指针的形式传值的，因为Set方法定义的是以指针方式取值
	(*One).Set(t, 2)
	fmt.Println(get(*t))

	// SetTwo方法指定接收者不是指针类型，这里无法修改到实例t的值
	// get(*t)依然是2
	// 因为go语言所有方法传递值都是值传递的方式，使用方法表达式很容易看出
	// 在定义类(struct类型数据)的方法时, 需要修改实例的属性，必须使用指针的形式指定接收者(即：(typeName *TypeName))
	// go语言值，方法(函数)参数是值类型，则传递的是值的副本；参数是指针类型，传递的也是指针的副本
	One.SetTwo(*t, 10)

	fmt.Println(get(*t))

	// 编译器做了响应处理（注意这里*t取值了，开头定义的t是地址类型(指针)）
	// (*t).Set 会编译成t.Set
	(*t).Set(10)
	fmt.Println(get(*t))
	t.Set(11)
	fmt.Println(get(*t))
	// t.Get() 会编译成(*t).Get()
	t.Get()
}
