package main

import (
	"errors"
	"fmt"
)

// defer是先进后出的，最先定义的defer函数最后执行
// defer后面必须是立即执行函数(即方法调用)，不能为语句
func filo() (err error){
	defer func() {
		fmt.Println("The first one...")
		if er := recover(); er != nil {
			err = errors.New(er.(string))
		}
	}()
	defer func() {
		fmt.Println("The second one...")
	}()
	defer fmt.Println("The third one...")
	fmt.Println("Main func filo...")
	// 主动使得程序退出，defer不会执行
	//os.Exit(1)
	panic("出错")
	return
}

func main() {
	err := filo()
	if err != nil {
		fmt.Println(err)
	}
}