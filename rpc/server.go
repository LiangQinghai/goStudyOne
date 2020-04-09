package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

func main() {
	myServer := new(MyRPCServer)
	err := rpc.Register(myServer)
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

type MyRPCServer struct {
	A,B string
}
type Args struct {
	A,B string
}

func (my *MyRPCServer) One(args *Args, reply *string)  error{

	fmt.Println(args.A, args.B)
	*reply = args.A + args.B
	return nil
}