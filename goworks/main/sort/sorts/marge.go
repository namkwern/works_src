package sorts
import(
	"fmt"
)

func Msort(data []int) []int{
	st := len(data) >> 1
	fmt.Print()
	var L []int
	var R []int
	if st != 0 && len(data[:st]) >= 1{
		L = Msort(data[:st])
	}
	if st != 0 && len(data[st:]) >= 1{
		R = Msort(data[st:])
	}
	if len(L) >= 100000{
		fmt.Println(count)
	}
	if len(L) != 0 && len(R) != 0{
		return merge(L, R)
	}
	return data
}

var count int
//マージする、最小単位の処理
func merge(L []int, R []int) []int{
	work := make([]int, len(L) + len(R))
	indexL := 0
	indexR := 0
	for n, _ := range work{
		if indexR >= len(R) || indexL < len(L) && L[indexL] < R[indexR]{
			work[n] = L[indexL]
			indexL++
		}else{
			work[n] = R[indexR]
			indexR++
		}
		count++
	}
	return work
}