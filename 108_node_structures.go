package main

import "fmt"

type Node struct {
	le *Node
	data interface{} // 空接口可以存储任何类型的数据
	ri *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data=data
}

func main() {
	root:=NewNode(nil, nil)
	root.SetData("root node")

	// make child(leaf) nodes:
	a:=NewNode(nil,nil)
	a.SetData("left node")
	b:=NewNode(nil,nil)
	b.SetData("right node")

	// build the tree:
	root.le=a
	root.ri=b
	fmt.Printf("%v\n", root)
}
