package utils

import "fmt"

func init() {

	fmt.Println("This is init function in package utils, it will be executed first(need to import package utils)...")

}

func Calc(x,y int) int{

	return x + y

}

func ThisIsTest() string {

	fmt.Println("This is a test。。。")

	return "Test"

}