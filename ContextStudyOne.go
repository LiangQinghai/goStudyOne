package main

import (
	"context"
	"fmt"
	"time"
)

type myContext struct {
	context.Context
}

func work(context context.Context, name string)  {
	for {
		select {
		case <- context.Done():
			fmt.Printf("Recieve a message from %s\n", name)
			return
		default:
			fmt.Printf("%s is running...\n", name)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func workWithValue(context context.Context, name string)  {
	for {
		select {
		case <- context.Done():
			fmt.Printf("Recieve a stop message from %s\n", name)
			return
		default:
			value := context.Value("KEY").(string)
			fmt.Printf("The value of key named KEY is %s, the routine keep running...\n", value)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	// 创建一个cancelContext，指定父上下文为空的Background
	withCancelContext, cancel := context.WithCancel(context.Background())
	go work(withCancelContext, "withCancelContext")

	// 创建一个带有超时的上下文对象，父上下文指定为cancelContext
	withDeadlineContext, _ := context.WithDeadline(withCancelContext, time.Now().Add(time.Second*5))
	go work(withDeadlineContext, "withDeadlineContext")

	// 创建一个带有键值对的上线文valueContext
	withValueContext := context.WithValue(withDeadlineContext, "KEY", "Tis is a value....")
	go workWithValue(withValueContext, "withValueContext")

	time.Sleep(time.Second * 10)
	cancel()

	time.Sleep(time.Second * 4)
}
