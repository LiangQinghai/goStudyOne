package main

import (
	"fmt"
	"reflect"
)

type PersonOne struct {
	Name string "This is a Tag"
	Age int `a:"TagA";b:"TagB"`
}

func main() {
	one := PersonOne{}
	// 反射获取实例信息
	typeOf := reflect.TypeOf(one)
	// FieldByName：根据名称获取属性，StructField
	if name, b := typeOf.FieldByName("Name"); b {
		// .Tag获取属性标签
		fmt.Println(name.Tag)
	}
	if age, b := typeOf.FieldByName("Age"); b {
		tag := age.Tag
		// 获取标签中的内容
		fmt.Println(tag.Get("a"), tag.Get("b"))
	}
	// 当前实例类名
	fmt.Println("Name: ", typeOf.Name())
	// 属性个数
	fmt.Println("NumField: ", typeOf.NumField())
	// 实例类所属包名
	fmt.Println("PkgPath: ", typeOf.PkgPath())
	// 报名.类名
	fmt.Println("String: ", typeOf.String())
	// 获取底层数据类型
	fmt.Println("Kind: ", typeOf.Kind().String())

	// 通过valueOf()获取值信息
	one.Name = "One"
	one.Age = 23
	valueOf := reflect.ValueOf(one)
	// 获取实例类型信息，相当于typeOf
	typeOf = valueOf.Type()
	for i := 0; i < typeOf.NumField(); i++ {
		fmt.Println(typeOf.Field(i).Name, valueOf.Field(i).Interface())
	}

	name := "HelloWorld"

	fieldByName := valueOf.FieldByName("Name")
	// 值类型是不可修改的
	if fieldByName.CanSet() {
		fieldByName.Set(reflect.ValueOf(name))
	}
	fmt.Println(fieldByName.CanSet(), one)

	// 指针类型的是可以修改的
	pointerValue := reflect.ValueOf(&one)
	// elem返回指针所指向的值
	if byName := pointerValue.Elem().FieldByName("Name"); byName.CanSet() {
		byName.Set(reflect.ValueOf(name))
	}
	fmt.Println(one)
}
