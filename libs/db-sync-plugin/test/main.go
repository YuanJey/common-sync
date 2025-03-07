package main

import (
	"fmt"
	"runtime"
)

func main() {
	start := 1734992727
	end := 1735031008
	fmt.Println(end - start)
	runtime.GOMAXPROCS(runtime.NumCPU())
}
