package main
import (
	"fmt"
	"flag"
)

func main(){
	
	str := flag.String("f", "", "flag テスト")
	flag.Parse()
	
	for n := 0; n < len(flag.Args()); n++{
		fmt.Println(flag.Arg(n))
	}
	fmt.Println(*str)
	
	var arr [30]int
	for n := range arr{
		arr[n] = n
	}
	slice := arr[0:10]
	slice = append(slice, arr[0:len(arr)]...)
	arr[3] = 100
	fmt.Println(arr)
	fmt.Print(slice)
	
}