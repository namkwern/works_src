//命名ルール 先頭=q|Q
package sorts
import "fmt"
import "time"

type Sortable interface{
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Qsort2(data Sortable){
	qsort(data, 0, data.Len() - 1)
}

//int(32or64)型スライスに対してソート
func qsort(data Sortable, start, end int){
	part := qset2(data, start, end)
	if part - start > 2{
		qsort(data, start, part)
	}
	if end - part > 2{
		qsort(data, part, end)
	}
}

//基準(最初と最後の要素の平均)をもとにスライス左右に振り分ける、クイックソートの最小単位の処理
func qset2(data Sortable, start, end int) int{
	center := (start + end) >> 1
	f := true
	outer:for{
		for data.Less(start, center){
			if start >= end{break outer}
			start++
		}
		for data.Less(center, end){
			if start >= end{break outer}
			end--
		}
		
		if !data.Less(start, end) && !data.Less(end, start){//同値要素回避処理
			if f{
				start++
			}else{
				end--
			}
			f = !f
			if start >= end{break outer}
		}else{//交換処理
			if center == start{
				center = end
			}else if center == end{
				center = start
			}
			data.Swap(start, end)
			f = true
		}
	}
	return start
}

func se(){fmt.Print();time.Now()}