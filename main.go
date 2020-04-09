package main

import (
	"fmt"
	"github.com/LiangQinghai/goStudyOne/utils"
	"runtime"
)

func init() {
	fmt.Println("This is init function in package main, it will be executed after all imported packages init function...")
}

func main() {
	res := utils.Calc(100, 500)
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(55)
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(res)
}

