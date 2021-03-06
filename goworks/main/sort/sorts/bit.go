//命名ルール 先頭=b|B
package sorts

import(
	//"fmt"
)

//log_2(max)を渡す
func Bitsort(data []int){
	max, min := bmax(data[:])
	minb := blog2(min)
	maxb := blog2(max)
	minus, plus := bdivpm(data[:])
	bitpm(plus, maxb)
	bitpm(minus, minb)
	
}

//+-部分のソート
func bitpm(data []int, bnum int){
	L, R := bset(data, bnum)
	//fmt.Println(s, data[:])
	if bnum > 0{
		bitpm(L, bnum - 1)
		bitpm(R, bnum - 1)
	}
}

//ビットソートの最小単位の交換処理
func bset(data []int, bnum int) ([]int, []int){
	if len(data) <= 1{
		return []int{}, []int{}
	}
	start := 0
	end := len(data) - 1
	outer:for{
		for refbit(data[start], uint(bnum)) == 0{
			if start == end{
				start++
				break outer
			}
			start++
		}
		for refbit(data[end], uint(bnum)) == 1{
			if start == end{
				break outer
			}
			end--
		}
		data[start],data[end] = data[end],data[start]
	}
	return data[:start], data[start:]
}

//任意のビット位置の値を参照する
func refbit(i int, b uint) int {
    return (i >> b) & 1
}

//+-振り分け
func bdivpm(data []int) ([]int, []int){
	start := 0
	end := len(data) - 1
	outer:for{
		for data[start] < 0{
			if start == end{
				break outer
			}
			start++
		}
		
		for data[end] >= 0{
			if start == end{
				break outer
			}
			end--
		}
		
		data[start],data[end] = data[end],data[start]
	}
	return data[:start], data[start:]
}

//log_2(n)を返す
func blog2(v int) int{
	var count int
	for n := v; n > 1 || n < -1; n = n >> 1{
		count++
	}
	return count
}

//最大値と最小値を返す
func bmax(data []int) (int, int){
	max := 0
	min := 0
	for n, v := range data{
		if v > data[max]{
			max = n
		}
		if v < data[min]{
			min = n
		}
	}
	return data[max], data[min]
}