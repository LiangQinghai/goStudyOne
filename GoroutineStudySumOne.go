package main

import (
	"fmt"
	"sync"
)

type Task struct {
	begin int
	end int
	result chan<- int
}

func (task *Task) job()  {
	sum := 0
	for i := task.begin; i <= task.end ; i ++ {
		sum += i
	}
	task.result <- sum
}

// 初始化任务
// 分配任务到taskList管道中
func initTask( taskList chan<- Task, result chan int, num int)  {
	count := num / 10
	mod := num % 10
	for i := 0; i < count; i++ {
		begin := i * 10 + 1
		end := (i + 1) * 10
		taskList <- Task{begin, end, result}
	}
	if mod > 0{
		taskList <- Task{begin: num - mod, end: num, result: result}
	}
	close(taskList)
}

// 分组开始执行takList管道中任务的job
func handleTask(taskList <-chan Task, wg *sync.WaitGroup, result chan int)  {
	for v := range taskList {
		wg.Add(1)
		go func() {
			v.job()
			wg.Done()
		}()
		wg.Wait()
	}
	close(result)
}

func main() {
	taskList := make(chan Task, 10)
	result := make(chan int, 10)

	initTask(taskList, result, 100)
	wg := &sync.WaitGroup{}
	handleTask(taskList, wg, result)
	for v := range result {
		fmt.Println(v)
	}
}