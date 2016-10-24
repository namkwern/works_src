package node2
import(
	"../node"
	"fmt"
)


func Make() node.Node{
	node := node.Node{
		Name:	"node2",
		Next:	map[string]node.Next{},
		Active:	true,
		Method:	func(){
			fmt.Println("3")
		},
		Result:	func(){
			fmt.Println(2)
		},
	}
	return node
}