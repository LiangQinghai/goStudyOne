package main

import (
	//_代表引用不使用，在go中引用的包必须使用到，如果未使用到回导致编译报错，使用_可以避免这种问题
	_ "fmt"
	_ "os"
)

func main() {

	//_用于代表占用
	var a int = 40

	a = a + 1

}
