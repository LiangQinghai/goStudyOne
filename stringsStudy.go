package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//str1长度为17，unicode编码中一个汉字占用3个字节，len()函数计算的是字符串字节长度
	str1 := "HelloWorld,您好"
	println(len(str1))

	//遍历字符串会输出unicode编码，无法输出汉字
	for i := 0; i < len(str1) ; i ++  {
		fmt.Print(str1[i], " ", string(str1[i]), " ")
	}
	fmt.Println()
	//处理汉字
	r := []rune(str1)
	for i:=0; i<len(r);i++{
		fmt.Printf("%c", r[i])
		fmt.Print(string(r[i]))
	}
	fmt.Println()
	//range会自动处理string中的汉字
	for _,i := range str1{
		fmt.Print(i, " ", string(i), " ")
	}
	fmt.Println()
	for _,i := range str1{
		fmt.Print(i, " ", string(i), " ")
	}

	//使用strconv API
	//数字转字符串
	fmt.Println(strconv.Itoa(99))
	//字符串转数字,转换出错会返回错误信息
	num1, e := strconv.Atoi("1234")
	if e == nil{
		fmt.Println(num1)
	}else {
		panic(e)
	}
	//十进制转换其他进制，返回类型是string
	fmt.Println(strconv.FormatInt(10000, 16))

	//字符串、字节数组相互转换
	stringByte := []byte("您好，世界")
	fmt.Println(stringByte)
	fmt.Println(string(stringByte))

	//strings包API使用
	//判断字符串相等,忽略大小写， 判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同。
	fmt.Println(strings.EqualFold("ABC", "abC")) //true
	fmt.Println("ABC" == "abC") //false

	//前缀判断
	fmt.Println(strings.HasPrefix("HelloWorld", "Hello")) //true
	//后缀判断
	fmt.Println(strings.HasSuffix("HelloWorld", "World")) //true
	//是否包含子串
	fmt.Println(strings.Contains("HelloWorld", "World")) //true
	//判断字符串是否包含子串中的任意字符
	fmt.Println(strings.ContainsAny("HelloWorld", "ThisIsWorld"))//true
	//计算字符串中包含子串的数量
	fmt.Println(strings.Count("HelloWorldHelloWorldHelloWorld", "World")) //3
	//计算子串在字符串中第一次出现的位置，没有返回-1
	fmt.Println(strings.Index("HelloWorld", "World"))//5
	//子串sep在字符串s中最后一次出现的位置，不存在则返回-1。
	fmt.Println(strings.LastIndex("HelloWorld", "World"))//5
	//首字母大写
	fmt.Println(strings.Title("hello world"))
	//转换成小写
	fmt.Println(strings.ToLower("HELLO WORLD"))
	//转换成大写
	fmt.Println(strings.ToUpper("hello world"))
	//
	fmt.Println(strings.ToTitle("Hello world"))
	//重复输出字符串
	fmt.Println(strings.Repeat("HelloWorld", 5))
	//替换,返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	fmt.Println(strings.Replace("HelloWorld","World", "WORLD", -1))
	//将s前后端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Println(strings.Trim("/restful/", "/"))
	//前后去掉空格
	fmt.Println(strings.TrimSpace(" Hello World "))
	//返回将s前端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Println(strings.TrimLeft("/HelloWorld/", "/"))
	//去掉右边的
	fmt.Println(strings.TrimRight("/HelloWorld/", "/"))
	//去掉前缀
	fmt.Println(strings.TrimPrefix("HelloWorld", "Hello"))
	//去掉后缀
	fmt.Println(strings.TrimSuffix("HelloWorld", "World"))
	//返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片
	fields := strings.Fields("Hello World")
	fmt.Println(len(fields))
	//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串
	fmt.Println(strings.Split("HelloWorld", "l"))
	//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目：
	//
	//n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
	//n == 0: 返回nil
	//n < 0 : 返回所有的子字符串组成的切片
	fmt.Println(strings.SplitN("HelloWorld", "l", 2))
	//用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	fmt.Println(strings.SplitAfter("HelloWorld", "l"))
	//将一系列字符串连接为一个字符串，之间用sep来分隔。
	fmt.Println(strings.Join(fields, "--"))
}
