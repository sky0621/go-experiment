package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("[pre]NumCPU:%d\n", runtime.NumCPU())
	fmt.Printf("[pre]NumGoroutine:%d\n", runtime.NumGoroutine())
	fmt.Printf("[pre]Version:%s\n", runtime.Version())
	fmt.Printf("[pre]CgoCall:%v\n", runtime.NumCgoCall())
	fmt.Printf("[pre]Compiler:%v\n", runtime.Compiler)
	fmt.Printf("[pre]GOOS:%v\n", runtime.GOOS)
	fmt.Printf("[pre]GOARCH:%v\n", runtime.GOARCH)
	fmt.Println("==========================================================")
	//fmt.Printf("[pre]MemStats: %#v\n", runtime.MemStats{})

}
