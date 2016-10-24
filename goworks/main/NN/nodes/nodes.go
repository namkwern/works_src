package nodes
import(
)

type NodeInterface interface{
	 Method()
	 Result(bool)
}
type Next struct{
	Node Node
	Weight int
}
type Node struct{
	Name string
	Next map[string]Next
	Active bool
	Method func()
	Result func()
}


func MakeNodeList() []Node{
	nodeList := []Node{MakeNode1(), MakeNode2()}
	for _, v := range nodeList{
		for _, v1 := range nodeList{
			if v.Name != v1.Name{
				v.Next[v1.Name] = Next{
					Node:	v1,
					Weight:	100,
				}
			}
		}
	}
	return nodeList
}


func MakeNode1() Node{
	node := Node{
		Name:	"node1",
		Next:	map[string]Next{},
		Active:	true,
		Method:	func(){
			
		},
		Result:	func(){
			
		},
	}
	return node
}


func MakeNode2() Node{
	node := Node{
		Name:	"node2",
		Next:	map[string]Next{},
		Active:	true,
		Method:	func(){
			
		},
		Result:	func(){
			
		},
	}
	return node
}