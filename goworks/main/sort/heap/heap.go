package main

import(
	"fmt"
	"math/rand"
	//"time"
)

func main(){
	/*
	var data [200000]int
	for n, _ := range data{
		data[n] = rand.Int() % 32768
	}
	
	fmt.Println(data[:30])
	s := time.Now()
	qsort(data[:])//ソートの実行
	e := time.Now()
	
	fmt.Println(e.Sub(s))//時間の表示
	fmt.Println(data[:30])
	*/
	//data := [3]int
	var n int64
	for ; n < 10; n++{
		rand.Seed(n)
		data := rand.Perm(10)
		fmt.Print(data)
		heapset(data, 0)
		fmt.Println(data)
	}
	
}

func hsort(data []int){
	
}

func heapset(data []int, index int){
	outer:for n := 1; ; n += n{
		if !set(data[:], index + m){
			break outer
		}
		index += index
	}
}

//最小単位の交換処理 親要素を最大にする
func set(data []int, index int) bool{
	RExist := true
	L := index << 1 + 1
	R := L + 1
	var max int
	if len(data) <= R{
		max = L
		RExist = false
	}else{
		if data[L] >= data[R]{
			max = L
		}else{
			max = R
		}
	}
	if data[index] < data[max]{
		data[index],data[max] = data[max],data[index]
	}
	return RExist
}