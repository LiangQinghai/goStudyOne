package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	strings := os.Args

	for i,v := range strings{
		fmt.Println(i, v)
	}

	var user, password string

	flag.StringVar(&user, "u", "root","用户名，默认root")
	flag.StringVar(&password, "p", "1234.abcd", "密码，默认1234.abcd")

	flag.Parse()
	fmt.Println(user, password)

}