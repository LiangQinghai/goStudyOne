package main

import "fmt"

func main() {


	//定义数组,指定下标
	strArr := [...]string{1:"Hello",0:"World",2:"!!!"}
	fmt.Println(strArr)

	//数组指针，指向第一个元素的内存地址
	changeArr(&strArr)
	fmt.Println(strArr)


}

func changeArr(strArr *[3]string) {

	strArr[len(strArr)/2] = "Change"

}
