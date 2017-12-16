package main

import (
	"fmt"
	"time"
)

const (
	arraySize = 100000000
)

func main() {
	var ary01 []int
	fmt.Printf("Experiment No.01 - ArraySize:%d\n", arraySize)
	fmt.Print("[Start] ")
	start := time.Now()
	fmt.Println(start.Format(time.StampMilli))
	for i := 0; i < arraySize; i++ {
		ary01 = append(ary01, i)
	}
	end := time.Now()
	fmt.Print("[End]   ")
	fmt.Println(end.Format(time.StampMilli))
	fmt.Print("[Diff]  ")
	fmt.Println(end.Sub(start))
}
