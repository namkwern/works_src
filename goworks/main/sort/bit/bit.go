package main

import(
	"fmt"
	"rand"
)

func main(){
	
	bitsort()
}

//log_2(max)を渡す
func bitsort(data []int){
	max := data[0]
	for n, v := range data{
		if v > max{
			max = n
		}
	}
	
}