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
	pm := bsetpm(data[:])
	bitp(data[pm:], maxb)
	bitm(data[:pm], minb)
	
}

//+部分のソート
func bitp(data []int, bnum int){
	s := bset(data, bnum)
	//fmt.Println(s, data[:])
	if bnum > 0{
		bitp(data[:s], bnum - 1)
		bitp(data[s:], bnum - 1)
	}
}

//-部分のソート
func bitm(data []int, bnum int){
	s := bset(data, bnum)
	//fmt.Println(s, data[:])
	if bnum > 0{
		bitp(data[:s], bnum - 1)
		bitp(data[s:], bnum - 1)
	}
}

//ビットソートの最小単位の交換処理
func bset(data []int, bnum int) int{
	if len(data) <= 1{
		return 0
	}
	start := 0
	end := len(data) - 1
	outer:for{
		for refbit(data[start], uint(bnum)) == 0{
			if start == end{
				return start + 1
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
	return start
}

//任意のビット位置の値を参照する
func refbit(i int, b uint) int {
    return (i >> b) & 1
}

//+-振り分け
func bsetpm(data []int) int{
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
	return start
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