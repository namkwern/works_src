package main

import(
	"fmt"
	"math/rand"
	"time"
	"./sorts"
)

func main(){
	data := arr1(200000)
	
	fmt.Println(data[:10], data[len(data) - 10:])
	
	s := time.Now()
	sorts.Qsort(data)//ソートの実行
	e := time.Now()
	
	for n, v := range data{
		if n != 0 && data[n - 1] > v{
			fmt.Println("Error!! n =", n, data[n:n + 2])
			break
		}
	}
	
	fmt.Println(e.Sub(s))//時間の表示
	fmt.Println(data[:100])
}

func arr1(size int) []int{
	fmt.Print("乱数のシード値を入力してください:")
	var i int64
	fmt.Scanf("%d", &i)//乱数のシード値入力
	rand.Seed(i)
	return rand.Perm(size)//ランダムスライス生成
}

func arr2(size int) []int{
	data := make([]int, size)
	for n, _ := range data{
		data[n] = n
	}
	return data[:]
}

func arr3(size int) []int{
	data := make([]int, size)
	for n, _ := range data{
		data[n] = n
		if n > len(data) / 2{
			data[n] = len(data) - n
		}
	}
	return data[:]
}