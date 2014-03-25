// harmonise project harmonise.go
package harmonise

// TODO For hashing look at https://github.com/stathat/consistent
import (
	"fmt"
	"net"
)

type NodeId struct {
	ip net.IP
}

type Node struct {
	Id                    NodeId
	Sucessor, Predecessor *Node
}

func FindClosestSucessor(nodeId NodeId) *Node {
	//TODO Replace placeholder
	return nil
}

func FindClosestPrecedingNode(nodeId NodeId) *Node {
	//TODO Replace placeholder
	return nil
}
func (node *Node) FindSucessor(nodeId NodeId) *Node {
	if nodeId == node.Sucessor.Id || nodeId == node.Id {
		return node.Sucessor
	} else {
		precedingNode := FindClosestPrecedingNode(nodeId)
		return precedingNode.FindSucessor(nodeId)
	}
}
func (node *Node) FindPredecessor(nodeId NodeId) *Node {

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
	return Node{NodeId{}, nil, nil}
}

func (node *Node) Join(n *Node) {
	node.Predecessor = nil
	node.Sucessor = n.FindSucessor(node.Id)

}

func main() {
	fmt.Println("Hello World!")
}
