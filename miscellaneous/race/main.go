package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int

func main() {
	runtime.GOMAXPROCS(2)

	go func () {
		for {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i++
	}
}
