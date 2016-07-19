package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	/*
	fmt.Print("乱数のシード値を入力してください:")
	var i int64
	fmt.Scanf("%d", &i)//乱数のシード値入力
	rand.Seed(i)
	data := rand.Perm(200000)//ランダムスライス生成
	
	*/
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
}
//int(32or64)型スライスに対してソート
func qsort(data []int){
	start := set(data[:])
	if len(data[:start]) > 1{
		sort(data[:start])
	}
	
	if len(data[start:]) > 1{
		sort(data[start:])
	}
}

//基準(最初と最後の要素の平均)をもとにスライス左右に振り分ける、クイックソートの最小単位の処理
func set(data []int) int{
	start := 0
	end := len(data) - 1
	center := (data[start] + data[end]) >> 1
	outer:for{
		for data[start] <= center{
			if start == end{
				break outer
			}
			start++
		}
		
		for data[end] > center{
			if start == end{
				break outer
			}
			end--
		}
		
		data[start],data[end] = data[end],data[start]
	}
	return start
}