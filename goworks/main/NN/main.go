package main
import(
	"./nodes"
	"fmt"
)

func main(){
	list := nodes.MakeNodeList()
	for n, v := range list{
		fmt.Println(n, v.Name)
		v.Method()
	}
}