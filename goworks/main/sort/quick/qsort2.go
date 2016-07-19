package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Print("乱数のシード値を入力してください:")
	var i int64
	fmt.Scanf("%d", &i)//乱数のシード値入力
	rand.Seed(i)
	data := rand.Perm(200000)//ランダムスライス生成
	fmt.Println(data[:30])
	s := time.Now()
	qsort(data)//ソートの実行
	e := time.Now()
	fmt.Println(e.Sub(s))//時間の表示
	fmt.Println(data[:30])
}
func qsort(a []int) []int {
  if len(a) < 2 { return a }

  left, right := 0, len(a) - 1

  // Pick a pivot
  pivotIndex := rand.Int() % len(a)

  // Move the pivot to the right
  a[pivotIndex], a[right] = a[right], a[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range a {
    if a[i] < a[right] {
      a[i], a[left] = a[left], a[i]
      left++
    }
  }

  // Place the pivot after the last smaller element
  a[left], a[right] = a[right], a[left]

  // Go down the rabbit hole
  qsort(a[:left])
  qsort(a[left + 1:])


  return a
}