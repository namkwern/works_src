package node1
import(
	"../node"
)


func Make() node.Node{
	node := node.Node{
		Name:	"node1",
		Next:	map[string]node.Next{},
		Active:	true,
		Method:	meth,
		Result:	resu,
	}
	return node
}

func meth(){

}
func resu(){

}