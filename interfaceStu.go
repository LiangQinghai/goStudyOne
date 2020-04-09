package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write(data interface{})
}

type IO struct {
	Data interface{}
}

func (io IO) Read()  {
	fmt.Println(io.Data)
}

func (io IO) Write(data interface{})  {
	io.Data = data
}

type Inter interface {
	Ping()
	Pong()
}

type St struct {

}

func (St) Ping()  {
	fmt.Println("Ping")
}

func (*St) Pong()  {
	fmt.Println("Pong")
}


func main() {

	io := IO{Data: "This is virtual data"}
	var reader Reader = io

	if r, ok := reader.(Writer); ok {
		r.Write("Change By Writer")
	}

	reader.Read()

	// 使用.(type)获取类型
	switch i := reader.(type) {
	case Writer:
		fmt.Println("writer")
	case Reader:
		fmt.Println("reader")
	case IO:
		fmt.Println("io")
	default:
		fmt.Println(i)
		return
	}

	//var one St
	var tmp *St  = nil
	//*tmp = one
	// 依然能够调用
	tmp.Pong()
	var writer Inter = tmp

	fmt.Printf("%p\n", writer)
	fmt.Printf("%p\n", tmp)

	// 此处执行为true
	if writer != nil {
		writer.Pong()
		// 产生panic异常
		// panic: value method main.St.Ping called using nil *St pointer
		writer.Ping()
	}
}