package node

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