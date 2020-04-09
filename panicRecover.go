package main

import (
	"errors"
	"fmt"
)

func TestPanic(num int)  (err error){

	//定义匿名函数处理错误
	defer func() {
		errString := recover()
		if errString != nil {
			fmt.Println(errString)
			err = errors.New(errString.(string))
		}
	}()

	if num < 0 || num > 100{
		panic("The number is invalidated...")
	}

	return
}

func severalPanic() (error error) {
	// 第一个defer最后执行，只会捕获上一步抛出的panic
	// 所以这里只会捕获到First One异常
	defer func() {
		if err := recover();err != nil {
			fmt.Println(err)
			error = errors.New(err.(string))
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("First defer recover: ", err)
		}
		panic("First One")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Second defer recover: ", err)
		}
		panic("Second One")
	}()

	panic("Main")
	return
}

func main() {

	err := TestPanic(101)
	fmt.Println(err)

	err = TestPanic(10)
	fmt.Println(err)

	err = severalPanic()
	fmt.Println(err)

}
