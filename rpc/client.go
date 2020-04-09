package main

import (
	"fmt"
	"net/rpc"
)

type MyType struct {
	A,B string
}

func main() {

	client, e := rpc.DialHTTP("tcp", "127.0.0.1:3000")
	if e != nil {
		panic(e)
	}

	myType := MyType{"Hello", "World"}

	var reply string
	e = client.Call("MyRPCServer.One", myType, &reply)
	if e != nil {
		panic(e)
	}
	fmt.Println(reply)

}
