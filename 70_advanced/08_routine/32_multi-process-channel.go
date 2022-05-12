package main

import (
	"fmt"
	"runtime"
)

// 当使用单线程执行时，会按部就班，按照顺序1，2，3，4执行下去
// 当使用多个CPU核数时，任务分配是不定的，

func main() {

	// 使用多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go3(c, i)
	}

	// 设置一个缓存长度为10 的channel
	for i := 0; i < 10; i++ {
		<-c
	}

}

func Go3(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}

	fmt.Println(index, a)
	c <- true

}
