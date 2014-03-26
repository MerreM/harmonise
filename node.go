// harmonise project harmonise.go
package harmonise

// TODO For hashing look at https://github.com/stathat/consistent
import (
	"fmt"
	"math"
)

const (
	k int = 6
)

var (
	MAX = int(math.Pow(2, float64(k)))
)

func decrease(value, size int) int {
	if size <= value {
		return value - size
	} else {
		return MAX - (size - value)
	}
}

func between(value, init, end int) bool {
	if init == end {
		return true
	} else if init > end {
		shift := MAX - init
		init = 0
		end = (end + shift) % MAX
		value = (value + shift) % MAX
	}
	result := init < value
	result = result && value < end
	return result
}

func Ebetween(value, init, end int) bool {
	if value == init {
		return true
	} else {
		return between(value, init, end)
	}
}

func betweenE(value, init, end int) bool {
	if value == end {
		return true
	} else {
		return between(value, init, end)
	}
}

type Node struct {
	id          int
	finger      map[int]*Node
	start       map[int]int
	predecessor *Node
}

func CreateStartId(id, i int) int {
	result := id + int(math.Pow(2, float64(i)))%int(math.Pow(2, float64(k)))
	return result
}
func CreateId(id int) int {
	return id
}

func Create(id int) Node {
	startMap := make(map[int]int)
	for i := 0; i < 10; i += 1 {
		startMap[i] = CreateStartId(id, i)
	}
	node := Node{
		CreateId(id),
		make(map[int]*Node),
		startMap,
		nil,
	}
	node.predecessor = *node
	return node
}

func (node *Node) GetSuccessor() *Node {
	return node.finger[0]
}

func (node *Node) GetPredecessor() *Node {

}

func (node *Node) FindSuccessor(id int) *Node {
	if betweenE(id, node.Predecessor.Id, node.id) {
		return node
	}
	predecessor := node.FindPredecessor(id)
	return predecessor.GetSuccessor()
}

func (node *Node) FindPredecessor(id int) {
	if node.id == id {
		return node.GetPredecessor()
	}
	n1 := node

}
func (node *Node) ClosestPrecedingFinger(id int) {

}

func main() {
	fmt.Println("Hello World!")
}
