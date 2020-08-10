package main

import (
	"fmt"
	"linkname/b"
)

func main() {
	s := b.Greet("world")
	fmt.Println(s)
	s = b.Hi("world")
	fmt.Println(s)
}
