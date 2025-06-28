package functions

func NPOINTER(b BigStruct){
	b.data1[0] = 42 // コピー側の操作
}

func POINTER(b *BigStruct){
	b.data1[0] = 42 // 実体の操作
}

//フィールドが多い構造体
//フィールドが少ないと実行速度はあまり変わらない
type BigStruct struct {
	data1 [10000]int
	data2 [10000]int
	data3 [10000]int
	data4 [10000]int
	data5 [10000]int
	data6 [10000]int
	data7 [10000]int
	data8 [10000]int
	data9 [10000]int
	data10 [10000]int
	data11 [10000]int
	data12 [10000]int
	data13 [10000]int
	data14 [10000]int
}