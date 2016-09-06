//命名ルール 先頭=q|Q
package sorts

import(
)

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
	start := 0
	end := len(data) - 1
	center := (data[start] + data[end]) >> 1
	outer:for{
		for data[start] < center{
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
	return data[:start], data[start:]
}