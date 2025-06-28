package functions

import (
	"bufio"
	"os"
)

func NBUFIO() {
	// 直接書き込むより...
	f, _ := os.Create("file.txt")
	f.WriteString("line 1\n")
	f.WriteString("line 2\n")
}

func BUFIO() {
	// bufioを使ってバッファリングする方が効率的
	f, _ := os.Create("file.txt")
	defer f.Close()
	
	writer := bufio.NewWriter(f)
	writer.WriteString("line 1\n")
	writer.WriteString("line 2\n")
	
	// バッファに残っている内容をディスクに書き込む
	writer.Flush() 
}