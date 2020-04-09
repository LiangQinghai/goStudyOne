package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dist string) (written int64, err error){

	file, e := os.Open(src)
	defer file.Close()
	if e != nil{
		panic(e)
		return -1,e
	}

	reader := bufio.NewReader(file)

	//
	openFile, e := os.OpenFile(dist, os.O_WRONLY|os.O_CREATE, 0666)
	defer openFile.Close()
	if e != nil{
		panic(e)
		return -1,e
	}
	writer := bufio.NewWriter(openFile)

	return io.Copy(writer, reader)

}

func main() {
	written, err := CopyFile("E:/QLDownload/test2.txt", "E:/QLDownload/test4.txt")
	fmt.Println(written, err)
}