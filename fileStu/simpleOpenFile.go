package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	file, e := os.Open("E:/QLDownload/test.txt")
	defer file.Close()
	if e == nil{
		fmt.Printf("The file is %v\n", file)
		//默认缓冲区大小为4096，
		reader := bufio.NewReader(file)
		for{
			s, e := reader.ReadString('\n')
			if e == io.EOF{
				break
			}
			fmt.Print(s)
		}


		//使用ioutil一次性将文件内容读取到内存中
		readFile, e := ioutil.ReadFile("E:/QLDownload/test1.txt")
		if e == nil {
			fmt.Println(string(readFile))
		}else {
			fmt.Println(e)
		}

	}else {
		fmt.Println(e)
	}
}
