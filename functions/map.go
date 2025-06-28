package functions

import(
	"fmt"
)

//n*m
func NMAP(a1, target []int){
	for _,a2 := range a1 {
		for _,t := range target{
			if a2 == t {
				fmt.Print("")
			}
		}
	}
}

//n+m
func MAP(a1, target []int){
	a2 := make(map[int]struct{})
	for _, n := range a1 {
		a2[n] = struct{}{}
	}
	for _, t := range target {
		if _, ok := a2[t]; ok {
				fmt.Print("")
		}
	}
}