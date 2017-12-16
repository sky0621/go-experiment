package main

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
)

const (
	// 1億
	arraySize = 100000000
)

// サイズ指定なしでスライス定義した場合と
// サイズ指定ありでスライス定義した場合の速度差
// およそ１０倍
//
// [サイズ指定なし]
// try count 1 : 1.3670205s
// try count 2 : 1.0777746s
// try count 3 : 1.0487405s
// try count 4 : 1.0887689s
// try count 5 : 1.0467384s
//
// [サイズ指定なし]
// try count 1 : 160.1065ms
// try count 2 : 148.1115ms
// try count 3 : 145.1081ms
// try count 4 : 149.0877ms
// try count 5 : 146.1105ms

func main() {
	defer profile.Start().Stop()
	fmt.Printf("Experiment No.01 - ArraySize:%d\n", arraySize)
	for i := 0; i < 5; i++ {
		fmt.Printf("try count %d : ", i+1)
		du := exp01()
		fmt.Println(du)
	}

	fmt.Printf("Experiment No.02 - ArraySize:%d\n", arraySize)
	for i := 0; i < 5; i++ {
		fmt.Printf("try count %d : ", i+1)
		du := exp02()
		fmt.Println(du)
	}
}

func exp01() time.Duration {
	var ary01 []int
	start := time.Now()
	for i := 0; i < arraySize; i++ {
		ary01 = append(ary01, i)
	}
	end := time.Now()
	return end.Sub(start)
}

func exp02() time.Duration {
	ary02 := make([]int, 0, arraySize)
	start := time.Now()
	for i := 0; i < arraySize; i++ {
		ary02 = append(ary02, i)
	}
	end := time.Now()
	return end.Sub(start)
}
