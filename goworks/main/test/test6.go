package main
import(
	"fmt"
	"math/rand"
	"hash/fnv"
)

func main(){
	var s string
	fmt.Scanln(&s)
	rand.Seed(makeHash(s))
	eva := []byte{'S', 'A', 'B', 'C', 'D'}//天賦の才・得意・普通・苦手・致命的
	prov := []int{5, 30, 40, 20, 5}
	
	str := []string{"体力", "魔力", "筋力", "耐久", "効率","節約", "回復", "速度", "道具"}
	status := make([]byte, len(str))
	for n, _ := range str{
		r := rand.Int63n(100) + 1
		sum := 0
		var index int
		for n, v := range prov{
			sum += v
			if sum >= int(r){
				index = n
				break
			}
		}
		status[n] = eva[index]
	}
	index := -1
	for n, v := range status{
		if v == 'S'{
			index = n
			break
		}
	}
	if index != -1{
		for n, v := range status{
			if v != 'S' && v != 'D'{
				status[n] = status[n] + 1
			}
		}
	}
	for n, v := range str{
		fmt.Println(v, ":", string(status[n]))
	}
}

func makeHash(str string) int64{
	hasher := fnv.New64()
	hasher.Write([]byte(str))
	return int64(hasher.Sum64())
}

