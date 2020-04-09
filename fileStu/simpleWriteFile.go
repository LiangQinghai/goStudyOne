package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//WriteTwo()
	//WriteAppend()
	//WriteUseIoUtil()

	info, e := os.Stat("E:/QLDownload/test4/")
	exist := os.IsNotExist(e)
	fmt.Println(info, e, exist)

}

func WriteOne()  {
	file, e := os.OpenFile("E:/QLDownload/test2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if e != nil{
		fmt.Println(e)
	}

	str := "HelloGoLang\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++{
		writer.Write([]byte(str))
	}
	writer.Flush()
}

func WriteTwo()  {
	//写入之前清空
	file, e := os.OpenFile("E:/QLDownload/test2.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()

	if e == nil{

		str := "你好世界！\t你好世界！\r\n"
		writer := bufio.NewWriter(file)
		for i := 0; i < 10 ;i++  {
			writeString, e := writer.WriteString(str)
			fmt.Println(writeString, e)
		}
		writer.Flush()

	}else {
		fmt.Println(e)
	}

}

func WriteAppend()  {
	//追加
	file, e := os.OpenFile("E:/QLDownload/test2.txt", os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if e == nil {
		str := "这是一行追加内容\r\n"
		writer := bufio.NewWriter(file)
		for i := 0; i < 10; i++{
			writer.WriteString(str)
		}
		writer.Flush()
	}else {
		fmt.Println(e)
	}
}

func WriteUseIoUtil()  {

	bytes, e := ioutil.ReadFile("E:/QLDownload/test2.txt")

	e = ioutil.WriteFile("E:/QLDownload/test3.txt", bytes, 0666)

	fmt.Println(e)

}
