package functions

//先に容量を確保しない場合
func NCAP(){
	var slice []int
	for i:=0;i<1000000;i++{
		slice = append(slice,i)
	}
}

//先に容量を確保する場合
func CAP(){
	slice := make([]int,0,1000000)
	for i:=0;i<1000000;i++{
		slice = append(slice,i)
	}
}