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

type NodeId struct {
	id int
}

type Node struct {
	id     NodeId
	finger map[int]Node
	start  map[int]NodeId
}

func CreateStartId(id, i int) NodeId {
	result := id + int(math.Pow(2, float64(i)))%int(math.Pow(2, float64(k)))
	return NodeId{result}
}
func CreateId(id int) NodeId {
	return NodeId{id}
}

func Create(id int) Node {
	startMap := make(map[int]NodeId)
	for i := 0; i < 10; i += 1 {
		startMap[i] = CreateStartId(id, i)
	}
	return Node{
		CreateId(id),
		make(map[int]Node),
		startMap,
	}
}

func (node *Node) getSuccessor() Node {
	return node.finger[0]
}

func main() {
	fmt.Println("Hello World!")
}
