//命名ルール 先頭=h|H
package sorts

import(
	"fmt"
)

func Hsort(data []int){
	
}

func Makeheap(data []int){
	i := (len(data) - 2) >> 1
	for n := i; n >= 0; n--{
		hset(data, n)
	}
}

//最小単位の交換処理 親要素を最大にする
func hset(data []int, index int) bool{
	L := index << 1 + 1
	
	if len(data) <= L + 1{//右要素が存在しない
		return false
	}
	if data[L] < data[L + 1]{
		L++
	}
	fmt.Print("")
	if data[index] < data[L]{
		data[index],data[L] = data[L],data[index]
	}
	return true
}