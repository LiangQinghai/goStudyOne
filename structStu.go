package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student struct {
	// 匿名字段，类型必须为命名类型或其指针
	// 默认字段名是其类型/指针指向的类型
	*Person
	//  不能使用Person字段，*Person默认为Person *Person
	//Person Person
}

type USB interface {
	ReadDiskData()
}

type TypeC struct {
}

func (c *TypeC) ReadDiskData() {
	fmt.Println("Read...")
}

func (c *TypeC) CFunc()  {

}

func MyRead(usb USB) {
	usb.ReadDiskData()
}

type A struct {
}

func (a *A) SayA() {
	fmt.Println("A")
}

type B struct {
}

func (b *B) SayB() {
	fmt.Println("B")
}


//继承
type C struct {
	A
	B
}

func (c *C) SayA() {
	fmt.Println("CA")
}
func (c *C) SayB() {
	fmt.Println("CB")
}


type One struct {

}

func (one One) MF()  {
	fmt.Println("One")
}

func (one One) TF()  {
	fmt.Println("One Tf")
}

type Two struct {
	One
}

func (two Two) MF()  {
	fmt.Println("Two")
}

type Three struct {
	Two
}

func (three Three) MF()  {
	fmt.Println("Three")
}

func main() {

	var p *Person = new(Person) //使用new的实例返回的是指针
	(*p).Name = "Hello"
	p.Name = "World" //两者等价，第一种是正规写法，第二种在go编译的时候也会转换为第一种写法

	p.Age = 23

	bytes, e := json.Marshal(*p)
	if e == nil {
		fmt.Println(string(bytes))
	}

	fmt.Println()

	//one := One{}

	two := Two{}

	three := Three{}
	three.MF()

	two.MF()
	// 嵌套其他结构类型，相当于集成其所有方法，可以重新编写其方法
	two.TF()

	two.One.MF()

}
