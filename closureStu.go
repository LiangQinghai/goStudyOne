package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

//函数闭包的使用
func completeImage(prefix, suffix string) func(string) string {

	return func(image string) string {

		if strings.HasPrefix(image, prefix) {
			if strings.HasSuffix(image, suffix) {
				return image
			} else {
				return image + suffix
			}
		} else {
			if strings.HasSuffix(image, suffix) {
				return prefix + image
			} else {
				return prefix + image + suffix
			}
		}
	}
}

// 返回多个闭包，共享同一个局部变量（其实用的是同一个地址的值）
func multiReturn(localVar int) (func(one int) int, func(two int) int)  {

	one := func(a int) int{
		localVar += a
		return localVar
	}
	two := func(b int) int{
		localVar -= b
		return localVar
	}
	return one, two
}

func main() {

	helpImage := completeImage("LQH_", ".png")

	fmt.Println(helpImage("Hello.png"))
	fmt.Println(helpImage("Hello"))
	fmt.Println(helpImage("LQH_Hello"))
	fmt.Println(helpImage("LQH_Hello.png"))

	sum := md5.Sum([]byte("HelloWorld"))
	fmt.Println(sum)
	// 初始化闭包局部变量值为10
	one, two := multiReturn(10)

	fmt.Println(one(2)) // 12
	fmt.Println(two(2)) // 10

}
