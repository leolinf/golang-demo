package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	resultCh := make(chan chan string, 5000)
	//开一个gotoutine 接受所有返回值并打印
	go replay(resultCh)
	//使用waitgroup 等待一下所有gorountie执行完毕，记录时间
	wg := sync.WaitGroup{}

	startTime := time.Now()

	//operation1内部sleep 1秒
	//operation2内部sleep 2秒
	//如果是同步执行下列调用需要 8秒左右
	//目前用异步调用 理论上只需要2秒
	//但于丹的问题是 不能实现先进先出的需求
	operation2(resultCh, "aaa", &wg)
	operation2(resultCh, "bbb", &wg)
	operation1(resultCh, "ccc", &wg)
	operation1(resultCh, "ddd", &wg)
	operation2(resultCh, "eee", &wg)
	wg.Wait()
	endTime := time.Now()
	fmt.Printf("Process time %s", endTime.Sub(startTime))
}

func replay(resultCh chan chan string) {
	for {
		c := <-resultCh
		r := <-c
		fmt.Println(r)
	}
}

func operation1(ch chan chan string, str string, wg *sync.WaitGroup) {
	c := make(chan string)
	ch <- c
	wg.Add(1)
	go func(str string) {
		time.Sleep(time.Second * 1)
		c <- "operation1:" + str
		wg.Done()
	}(str)
}

func operation2(ch chan chan string, str string, wg *sync.WaitGroup) {
	c := make(chan string)
	ch <- c
	wg.Add(1)
	go func(str string, wg *sync.WaitGroup) {
		time.Sleep(time.Second * 2)
		c <- "operation2:" + str
		wg.Done()
	}(str, wg)
}
