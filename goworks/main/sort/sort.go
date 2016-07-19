package main

import(
	"fmt"
	"math/rand"
	"time"
	"./sorts"
)

func main(){
	data := ranArr1()
	
	fmt.Println(data[:10])
	
	s := time.Now()
	sorts.Makeheap(data[:])//ソートの実行
	e := time.Now()
	
	for n := 1; n < len(data); n++{
		if data[n - 1] > data[n]{
			fmt.Println("Error")
			break
		}
	}
	
	fmt.Println(e.Sub(s))//時間の表示
	fmt.Println(data[:10])
}

func ranArr1() []int{
	fmt.Print("乱数のシード値を入力してください:")
	var i int64
	fmt.Scanf("%d", &i)//乱数のシード値入力
	rand.Seed(i)
	return rand.Perm(200000)//ランダムスライス生成
}

func ranArr2() []int{
	var data [10000]int
	for n, _ := range data{
		data[n] = 1 << 62 - 1
	}
	return data[:]
}