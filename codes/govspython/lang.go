package main

import (
	"fmt"
	"time"
)

func foo(c chan int) {
	time.Sleep(1 * time.Second)
	c <- 1024
}

func main() {
	c := make(chan int)
	go foo(c)
	fmt.Println("wait chan 'c' for 1 second")
	fmt.Println(<-c)
}
