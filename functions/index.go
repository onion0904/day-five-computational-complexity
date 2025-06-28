package functions

func NINDEX(s []BigStruct){
	for _,v := range s{
		_ = v
	}
}

func INDEX(s []BigStruct){
	for i := range s{
		_ = s[i].data1[0]
	}
}