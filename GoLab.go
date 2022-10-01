package main

import (
	"fmt"
	"time"
)

func f(num int) {
	fmt.Println("Hello from goroutine", num)
}

func main() {
	for i := 1; i < 6; i++ {
		go f(i)
	}
	time.Sleep(3 * time.Second)
}
