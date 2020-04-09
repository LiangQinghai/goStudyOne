package main

import "fmt"

func main() {

	//字符串不可变
	str := "HelloWorld"
	strSlice := str[:]
	fmt.Println(strSlice)

	for i,v := range strSlice{
		fmt.Println(i, string(v))
	}

	//将字符串转换为字节数组
	byteStr := []byte(str)
	byteStr[0] = 'T'
	fmt.Println(string(byteStr))

	//使用[]rune处理中文
	str2 := "你好世界"
	strRune := []rune(str2)
	strRune[2] = '好'
	fmt.Println(string(strRune))


	sliceOne := make([]string, 3)

	sliceOne[2] = "Have"

	fmt.Println(sliceOne)
	//slice改变，对应的数组也会相应的给改，slice中存储的是数组对应的地址
	arr := [3]int{1 ,2, 3}
	sliceTwo := arr[:]
	sliceTwo[2] = 10
	fmt.Println(arr)
	fmt.Println(sliceTwo)

	//使用append方法对slice进行扩容
	// arr[:]...表示首先将数组转换为slice，然后取slice中的所用元素追加到对应的slice上
	//使用append进行slice扩容，其实是重新分配一个新的数组
	sliceTwo = append(sliceTwo, arr[:]...)
	fmt.Println(sliceTwo)

	//使用copy进行slice复制
}
