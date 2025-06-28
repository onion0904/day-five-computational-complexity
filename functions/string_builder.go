package functions

import "strings"

// 悪い例
func NSB(){
	var str1 string
	words := []string{"hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go"}
	for _, w := range words {
		str1 += w + " "
	}
}

// 良い例
func SB(){
	words := []string{"hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go","hello", "world", "go"}
	var builder strings.Builder
	// あらかじめ必要なサイズを指定するとさらに効率的
	builder.Grow(100)
	for _, w := range words {
		builder.WriteString(w)
		builder.WriteString(" ")
	}
}