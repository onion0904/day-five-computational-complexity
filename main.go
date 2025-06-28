package main

import (
	"day-five/functions"
	"fmt"
	"time"
)

func main() {
	measureExecutionTime("ダミー",functions.NCAP)//最初だけなぜか時間がかかる
	a1, target := generateData(10000)
	measureExecutionTime("map不使用",func(){functions.NMAP(a1,target)})
	measureExecutionTime("map使用", func(){functions.MAP(a1,target)})
	measureExecutionTime("capacityなし",functions.NCAP)
	measureExecutionTime("capacityあり", functions.CAP)
	measureExecutionTime("string_builderを不使用",functions.NSB)
	measureExecutionTime("string_builderを使用", functions.SB)
	s1 := functions.BigStruct{}
	measureExecutionTime("pointerを不使用",func(){functions.NPOINTER(s1)})
	measureExecutionTime("pointerを使用", func(){functions.POINTER(&s1)})
	s2 := []functions.BigStruct{}
	measureExecutionTime("コピーでrangeを操作",func(){functions.NINDEX(s2)})
	measureExecutionTime("INDindexでrangeを操作EX", func(){functions.INDEX(s2)})
	measureExecutionTime("普通にfile読み込み",func(){functions.NBUFIO()})
	measureExecutionTime("bufioでfile読み込み", func(){functions.BUFIO()})
}



func measureExecutionTime(name string, fn func()) {
	if name == "ダミー"{
		fn()
		return
	}
	start := time.Now()//計測開始
	defer func() {
		fmt.Printf("%s にかかった時間: %s\n", name, time.Since(start))
	}()
	fn()
}

//決められたサイズのデータを二つ用意
func generateData(size int) ([]int, []int) {
	a1 := make([]int, size)
	for i := 0; i < size; i++ {
		a1[i] = 1 // もしくは rand.Intn(10) など
	}
	target := make([]int, size)
	for i := 0; i < size; i++ {
		target[i] = i + size // ヒットしない値
	}
	return a1, target
}