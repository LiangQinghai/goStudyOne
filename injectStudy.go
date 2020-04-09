package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type S1 interface {

}

func say(one string, s1 S1, i int)  {
	fmt.Println(one, s1, i)
}

func main() {

	injector := inject.New()

	injector.Map("One")
	injector.MapTo("hh", (*S1)(nil))
	injector.Map(12)

	values, e := injector.Invoke(say)
	fmt.Println(values, e)

	_ = 2
	// no new variables on left side of :=，_ 作为预设变量，不能被重新定义
	//_ := 2

}
