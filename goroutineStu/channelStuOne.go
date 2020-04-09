package goroutineStu

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func StudyChannelOne(){

	//使用chan int定义一个int类型的管道
	var intCha chan int
	//使用make对int管道进行内存分配，分配cap容量为9，管道一旦分配好容量不能更改
	intCha = make(chan int, 9)

	//向管道中存放数据
	num := 99
	intCha <- num

	//从管道中获取数据
	tmp := <- intCha
	fmt.Println(tmp)

	//管道如同队列，从管道中存放一个数据管道长度会增加一个单位，从管道中取出一个数据管道长度会减少一个单位

}


type MVP struct {
	Name string
	Team string
}

func StudyChannelTwo(){

	mc := make(chan interface{},10)

	mc <- MVP{Name:"KobeBryant", Team:"Lakers"}
	mc <- MVP{Name:"MichaelJordan", Team:"Bulls"}
	mc <- "Hello"

	mvp := <- mc
	mvp,ok := mvp.(MVP)
	if ok{
		fmt.Println(mvp)
	}

}


func ForRangeChannel(){


	ch := make(chan int, 100)

	for i := 0; i < 100 ; i++  {
		ch <- i
	}

	//遍历管道之前需要关闭管道，否则会出现deadlock现象，关闭管道之后，不能对管道进行写操作，但是能够读取管道里面的数据
	close(ch)
	for v := range ch{
		fmt.Println(v)
	}

}

//参数ch生命为只写管道，<-符号在chan后
func writeData( ch chan<- int)  {

	for i := 0; i < 100 ; i++ {
		ch<-i
		fmt.Println("Write...", i)
	}
	close(ch)

}

//参数readChannel声明为只读管道，以防进行误操作，符号<-在chan前
func readData(readChannel <-chan int, exitChannel chan bool){
	for  {
		v, ok := <-readChannel
		if !ok{
			break
		}
		fmt.Println("ReadData...", v)
		time.Sleep(time.Second * 5)
	}
	exitChannel<-true
	close(exitChannel)
}

func WriteReadData()  {

	writeReadChannel := make(chan int, 10)
	exitChannel := make(chan bool,1)

	go writeData(writeReadChannel)
	go readData(writeReadChannel, exitChannel)

	for {
		_, ok := <- exitChannel
		if !ok {
			break
		}
	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//使用select处理管道阻塞发生的deadlock
func SelectTest(){

	intChan := make(chan int, 10)
	for i :=0 ; i < 10 ; i++  {
		intChan <- i
	}

	stringChan := make(chan string, 10)
	for i := 1; i < 10 ; i++ {
		stringChan <- fmt.Sprintf("This is %d", i)
	}

	ChannelLoop:
		for {
			select {
			case v := <-intChan:
				fmt.Println("Read From IntChan...", v)
			case v := <-stringChan:
				fmt.Println("Read From StringChan...", v)
			default:
				fmt.Println("There is nothing to read...")
				//跳出for循环
				break ChannelLoop
			}
		}

}

/*****************************************************************************************************************************************************************************************************/
//使用defer+recover处理goroutine中出现的错误

func SyntaxStu()  {

	c := make(chan struct{})

	go func(i chan struct{}) {
		fmt.Printf("%T\n", i)
		for i:= 0; i <= 10 ; i++  {
			fmt.Printf("%T\n", i)
		}
		i <- struct{}{}
	}(c)

	<-c

}

func SyntaxOne()  {
	c := make(chan struct{})
	data := make(chan int, 100)

	go func(synChan chan struct{}, dataChan chan int) {
		for i:=0;i<100 ;i++  {
			data  <- i
		}
		// 关闭通道防止产生读取数据时死锁
		close(data)
		c <- struct{}{}
	}(c, data)

	<- c
	for key := range data {
		println(key)
	}
}

// waitGroup同步机制
func WaitGroupStu() {
	wg := sync.WaitGroup{}
	urls := []string{
	"https://www.baidu.com",
	"https://www.taobao.com",
	}

	for _, url := range urls {
		//  添加一个任务
		wg.Add(1)
		go func(url string) {
			// 释放任务，等同于wg.Add(-1)
			defer wg.Done()
			resp, err := http.Get(url)
			if err!=nil {
				fmt.Println(err)
			} else {
				fmt.Println(resp.Status)
			}
		}(url)
	}
	// 等待任务执行完毕
	wg.Wait()
}

// select可以用于向管道读取数据或写人数据
func SelectStudyTwo()  {
	one := make(chan int, 10)
	go func(chan int) {
		for {
			select {
			case one <- 0:
				fmt.Println("写入0操作...")
			case one <- 1:
				fmt.Println("写入1操作...")
			}
		}
	}(one)

	for i := 0; i <= 10; i++ {
		fmt.Print(<-one, "---")
	}

	two := make(chan string, 10)

	go func(chan string) {
		for i := 0; i <= 10; i++ {
			two <- fmt.Sprintf("No.%d", i+1)
		}
	}(two)
	doneSign := make(chan int)
	go func() {
	Label:
		for {
			select {
			case v := <-two:
				fmt.Println("Two: ", v)
			case v := <-one:
				fmt.Println("One: ", v)
			case <- doneSign: // 接收到主线程穿过来的停止信号，退出当前循环监听通道
				close(one)
				close(two)
				break Label
			}
		}
	}()
	time.Sleep(3 * time.Second)
	// 发送关闭通知
	close(doneSign)
}

// 带缓冲的随机数生成器
// for循环向channel中存数据，数据满时发生阻塞等待取出数据
func RandomGenerator() chan int{

	randomStore := make(chan int, 10)

	go func() {
		for {
			randomStore <- rand.Int()
		}
	}()

	return randomStore
}

func RandomGenerator2()  {
	randomStore := make(chan int, 10)

	go func() {
		for {
			randomStore <- rand.Int()
		}
	}()

	fmt.Println(<-randomStore, <-randomStore)
}

func MultiRoutineGenerator()  chan int{
	randomStore := make(chan int, 20)
	go func() {
		for {
			select {
			case randomStore <- <-RandomGenerator():
			case randomStore <- <-RandomGenerator():
			}
		}
	}()
	return randomStore
}
