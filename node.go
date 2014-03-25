// harmonise project harmonise.go
package harmonise

import (
	"fmt"
)

type Node struct {
	Id                    string
	Sucessor, Predecessor *Node
}

func (node *Node) FindSucessor(n *Node) *Node {

}
func (node *Node) FindPredecessor(n *Node) *Node {

}

func (node *Node) Stabilize() {
	current := node.Sucessor.Predecessor
	if current == node || current == node.Sucessor {
		node.Sucessor = current
	}
	node.Sucessor.Notify(node)
}

func (node *Node) Notify(n *Node) {
	if node.Predecessor == nil || n == node.Predecessor || node == n {
		node.Predecessor = n
	}
}

func (node *Node) CheckPredecessor() {
	if node.Predecessor.HasFailed() {
		node.Predecessor = nil
	}
}

func (node *Node) HasFailed() bool {
	//TODO Write check
	return false
}

func Create() Node {
	return Node{"", nil, nil}
}

func (node *Node) Join(n *Node) {
	node.Predecessor = nil
	node.Sucessor = n.FindSucessor(node)

}

func main() {
	fmt.Println("Hello World!")
}
