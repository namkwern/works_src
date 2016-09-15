package main

import(
	"fmt"
	"math/rand"
)

func main(){
	data := arr1(200000)	//シャッフル済み配列生成　引数は要素数
	Qsort(data)				//ソートの実行
	
	//ソート済みかチェック
	for n, v := range data{
		if n != 0 && data[n - 1] > v{
			fmt.Println("Error!! n =", n, data[n:n + 2])
			break
		}
	}
	
	fmt.Println(data[:100])
}

func arr1(size int) []int{
	fmt.Print("乱数のシード値を入力してください:")
	var i int64
	fmt.Scanf("%d", &i)//乱数のシード値入力
	rand.Seed(i)
	return rand.Perm(size)//ランダムスライス生成
}

//int(32or64)型スライスに対してソート
func Qsort(data []int){
	L, R := qset(data[:])
	if len(L) > 1{
		Qsort(L)
	}
	if len(R) > 1{
		Qsort(R)
	}
}

//基準(最初と最後の要素の平均)をもとにスライス左右に振り分ける、クイックソートの最小単位の処理
func qset(data []int) ([]int, []int){
	start, end := 0, len(data) - 1
	center := (data[start] + data[end]) >> 1
	f := true
	outer:for{
		for data[start] < center{
			if start >= end{break outer}
			start++
		}
		for data[end] > center{
			if start >= end{break outer}
			end--
		}
		
		if data[start] == data[end]{//同値要素回避処理
			if f{
				start++
			}else{
				end--
			}
			f = !f
			if start >= end{break outer}
		}else{//交換処理
			data[start],data[end] = data[end],data[start]
			f = true
		}
	}
	return data[:start], data[start:]
}