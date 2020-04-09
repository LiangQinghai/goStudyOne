package goroutineStu

import (
	"fmt"
	"sync"
	"time"
)

var (
	storage map[int]int
	lock sync.Mutex
)

func count(n int){

	res := 1
	for i := n; i <= n+10; i++ {
		res *= i
	}

	lock.Lock()
	storage[n] = res
	lock.Unlock()
}


func Start(){

	storage = make(map[int]int)
	for i := 1;i <= 20; i++{
		go count(i)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()

	for k, v := range storage{
		fmt.Println(k, v)
	}

	lock.Unlock()

}